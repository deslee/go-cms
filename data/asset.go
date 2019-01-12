package data

import (
	"context"
	"database/sql"
	"time"
)

type Asset struct {
	ID            string
	State         string
	Type          string
	Data          string
	CreatedAt     time.Time
	CreatedBy     string
	LastUpdatedAt time.Time
	LastUpdatedBy string
}

func (asset Asset) FileName() string {
	panic("not implemented")
}

func (asset Asset) Extension() string {
	panic("not implemented")
}

func (asset Asset) Items(context context.Context, db sql.DB) ([]Item, error) {
	panic("not implemented")
}

func GetAsset(ctx context.Context, db sql.DB, s string) (*Asset, error) {
	panic("not implemented")
}

func DeleteAsset(ctx context.Context, db sql.DB, assetId string) (GenericResult, error) {
	panic("not implemented")
}