package main

import (
	"context"
	"github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

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
}
