package data

import (
	"context"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
)

func FileNameFromAsset(asset Asset) string {
	panic("not implemented")
}

func ExtensionFromAsset(asset Asset) string {
	panic("not implemented")
}

func ItemsFromAsset(ctx context.Context, db *sqlx.DB, asset Asset) ([]Item, error) {
	return repository.ScanItemList(ctx, db, "SELECT I.* FROM ItemAssets IA INNER JOIN Items I on IA.ItemId = I.Id WHERE IA.AssetId=?", asset.Id)
}

func GetAsset(ctx context.Context, db *sqlx.DB, s string) (*Asset, error) {
	panic("not implemented")
}

func DeleteAsset(ctx context.Context, db *sqlx.DB, assetId string) (GenericResult, error) {
	panic("not implemented")
}
