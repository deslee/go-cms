package data

import (
	"context"
	. "github.com/deslee/cms/models"
	"github.com/jmoiron/sqlx"
)

func FileNameFromAsset(asset Asset) string {
	panic("not implemented")
}

func ExtensionFromAsset(asset Asset) string {
	panic("not implemented")
}

func ItemsFromAsset(ctx context.Context, db *sqlx.DB, asset Asset) ([]Item, error) {
	panic("not implemented")
}

func GetAsset(ctx context.Context, db *sqlx.DB, s string) (*Asset, error) {
	panic("not implemented")
}

func DeleteAsset(ctx context.Context, db *sqlx.DB, assetId string) (GenericResult, error) {
	panic("not implemented")
}
