package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/disintegration/imaging"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"time"
)


var widths = []int{1200, 500, 320}

func main() {
	db, err := sqlx.Open("sqlite3", "database.sqlite?_loc=auto")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	for {
		assets, err := repository.ScanAssetList(ctx, db, "SELECT A.* FROM Assets A WHERE A.State=?", "NONE")
		if err != nil {
			log.Printf("%s", err)
		}

		// TODO: maybe parallelize this?
		for _, asset := range assets {
			process(ctx, db, asset)
		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func process(ctx context.Context, db *sqlx.DB, asset model.Asset) {
	log.Printf("Processing Asset %s", asset.Id)

	fileName := asset.Key()
	if len(fileName) == 0 {
		log.Print("Filename not found!")
		return
	}
	filePath := fmt.Sprintf("./assets/%s", fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Print("File does not exist on disk!")
		return
	}

	src, err := imaging.Open(filePath)
	if err != nil {
		log.Print(err)
		return
	}

	sizes := make(map[string]string)

	log.Printf("Loaded asset %s", asset.Id)

	for _, w := range widths {
		// compute the width
		var width int
		if w < src.Bounds().Size().X {
			width = w
		} else {
			width = src.Bounds().Size().X
		}

		// compute the output file path
		key := fmt.Sprintf("%s-%d%s", asset.Id, width, asset.Extension())
		outputFilePath := fmt.Sprintf("./assets/%s", key)

		// widthHeightRatio := float64(src.Bounds().Size().Y) / float64(src.Bounds().Size().X)

		// resize the image
		resized := imaging.Resize(src, width, 0, imaging.Lanczos)

		// save the image
		err = imaging.Save(resized, outputFilePath)
		if err != nil {
			log.Print(err)
			return
		}

		// store the filepath in the map
		sizes[strconv.Itoa(width)] = key
	}

	// save the image
	var assetDataIf interface{}
	err = json.Unmarshal([]byte(asset.Data), &assetDataIf)
	if err != nil {
		log.Print(err)
		return
	}

	assetData, ok := assetDataIf.(map[string]interface{})
	if !ok {
		log.Printf("Could not deserialize asset data for asset %s", asset.Id)
		return
	}

	assetData["sizes"] = sizes
	jsonBytes, err := json.Marshal(assetData)
	if err != nil {
		log.Print(err)
		return
	}
	asset.Data = string(jsonBytes)
	asset.State = "RESIZED"
	err = repository.UpsertAsset(ctx, db, asset)
	if err != nil {
		log.Print(err)
		return
	}
}
