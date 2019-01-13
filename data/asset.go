package data

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Asset struct {
	Id    string `db:"Id"`
	State string `db:"State"`
	Type  string `db:"Type"`
	Data  string `db:"Data"`
	AuditFields
}

func (Asset) TableName() string {
	return "Assets"
}

func (asset Asset) FileName() string {
	panic("not implemented")
}

func (asset Asset) Extension() string {
	panic("not implemented")
}

func (asset Asset) Items(ctx context.Context, db *sqlx.DB) ([]Item, error) {
	panic("not implemented")
}

func GetAsset(ctx context.Context, db *sqlx.DB, s string) (*Asset, error) {
	panic("not implemented")
}

func DeleteAsset(ctx context.Context, db *sqlx.DB, assetId string) (GenericResult, error) {
	panic("not implemented")
}
