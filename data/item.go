package data

import (
	"context"
	"database/sql"
	"time"
)

type Item struct {
	ID            string
	Data          JSONObject
	CreatedAt     time.Time
	CreatedBy     string
	LastUpdatedAt time.Time
	LastUpdatedBy string
}

type ItemInput struct {
	ID     *string  `json:"id"`
	Type   string   `json:"type"`
	Data   string   `json:"data"`
	Groups []string `json:"groups"`
}

type ItemResult struct {
	Data         *Item   `json:"data"`
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
}

func GetItems(ctx context.Context, db *sql.DB, s string) ([]Item, error) {
	panic("not implemented")
}

func GetItem(ctx context.Context, db *sql.DB, s string) (*Item, error) {
	panic("not implemented")
}

func (item Item) Groups(ctx context.Context, db *sql.DB) ([]Group, error) {
	panic("not implemented")
}

func UpsertItem(ctx context.Context, db *sql.DB, item ItemInput, siteId string) (ItemResult, error) {
	panic("not implemented")
}

func DeleteItem(ctx context.Context, db *sql.DB, itemId string) (GenericResult, error) {
	panic("not implemented")
}
