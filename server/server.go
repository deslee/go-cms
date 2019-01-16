package main

import (
	"encoding/json"
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/deslee/cms"
	"github.com/deslee/cms/data"
	"github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"
)

const defaultPort = "3000"

var db *sqlx.DB

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	var err error
	db, err = sqlx.Open("sqlite3", "database.sqlite?_loc=auto")
	if err != nil {
		panic(err)
	}

	data.CreateTablesAndIndicesIfNotExist(db)
	withCors := cors.AllowAll().Handler

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle(
		"/graphql",
		withCors(
			withAuth(
				handler.GraphQL(cms.NewExecutableSchema(
					cms.Config{
						Resolvers: &cms.Resolver{DB: db},
					}),
				),
			),
		),
	)

	http.Handle(
		"/asset/",
		withCors(
			http.HandlerFunc(viewAssetHandler),
		),
	)

	http.Handle(
		"/uploadAsset",
		withCors(
			withAuth(
				http.HandlerFunc(uploadAssetHandler),
			),
		),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func viewAssetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get the id from the route
	routeSegments := strings.SplitN(r.URL.Path, "/asset/", 2)
	if len(routeSegments) != 2 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	id := routeSegments[1]

	// find the asset by id
	asset, err := repository.FindAssetById(r.Context(), db, id)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if asset == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// deserialize the data part of the asset
	var d interface{}
	err = json.Unmarshal([]byte(asset.Data), &d)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	assetData, ok := d.(map[string]interface{})
	if !ok {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	// construct the filename
	fileName := asset.Key()
	if len(fileName) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// get the size
	size := r.URL.Query().Get("w")
	if len(size) != 0 {
		// if a size is requested, get the sizeKeyMap from the asset data
		dataSizes, ok := assetData["sizes"]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		sizeKeyMap, ok := dataSizes.(map[string]interface{})
		if !ok {
			fmt.Printf("%s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// get the filename out of the sizeKeyMap
		fileNameData := sizeKeyMap[size]
		fileName, ok = fileNameData.(string)
		if !ok {
			fmt.Printf("%s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// get the file
	file, err := os.Open(fmt.Sprintf("./assets/%s", fileName))
	defer file.Close()
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", asset.Type)
	}
}

func uploadAssetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	siteId := r.Form.Get("siteId")
	if len(siteId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// validate
	hasAccess, err := data.AssertContextUserHasAccessToSite(r.Context(), db, siteId)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if hasAccess == false {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	name := strings.Split(header.Filename, ".")
	filename := name[0]
	extension := fmt.Sprintf(".%s", name[len(name) - 1])
	id := data.GenerateId()
	savedFilename := fmt.Sprintf("%s%s", id, extension)

	// ensure asset directory exists
	if _, err := os.Stat("./assets"); os.IsNotExist(err) {
		if err == os.Mkdir("./assets", os.ModePerm) {
			fmt.Printf("%s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// save asset to directory
	f, err := os.Create(fmt.Sprintf("./assets/%s", savedFilename))
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// save asset to database
	asset := model.Asset{
		Id: id,
		State: "NONE",
		SiteId: siteId,
		Type: mime.TypeByExtension(extension),
		Data: fmt.Sprintf(`{"extension": "%s", "originalFilename": "%s", "key": "%s"}`, extension, filename, savedFilename),
	}

	err = repository.UpsertAsset(r.Context(), db, asset)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	serialized, err := json.Marshal(asset)
	if err != nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// really make sure we have the data on disk
	err = f.Sync()
	if err == nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write(serialized)
	if err == nil {
		fmt.Printf("%s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}

func withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(auth) == 2 {
			if auth[0] == "Bearer" {
				token := auth[1]
				ctx, err := data.ParseTokenToContext(r.Context(), token)
				if err != nil {
					http.Error(w, "authorization failed", http.StatusUnauthorized)
					return
				}
				r = r.WithContext(ctx)
			} else {
				http.Error(w, "invalid scheme", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
