package data

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Item struct {
	Id   string     `db:"Id"`
	Data JSONObject `db:"Data"`
	AuditFields
}

func (Item) TableName() string {
	return "Items"
}

type ItemInput struct {
	Id     *string  `json:"id"`
	Type   string   `json:"type"`
	Data   string   `json:"data"`
	Groups []string `json:"groups"`
}

type ItemResult struct {
	GenericResult
	Data *Item `json:"data"`
}

func GetItems(ctx context.Context, db *sqlx.DB, s string) ([]Item, error) {
	panic("not implemented")
}

func GetItem(ctx context.Context, db *sqlx.DB, s string) (*Item, error) {
	panic("not implemented")
}

func (item Item) Groups(ctx context.Context, db *sqlx.DB) ([]Group, error) {
	panic("not implemented")
}

func UpsertItem(ctx context.Context, db *sqlx.DB, item ItemInput, siteId string) (ItemResult, error) {
	panic("not implemented")
}

func DeleteItem(ctx context.Context, db *sqlx.DB, itemId string) (GenericResult, error) {
	panic("not implemented")
}
