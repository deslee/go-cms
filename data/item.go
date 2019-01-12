package data

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Item struct {
	ID   string     `gorm:"type:text;primary_key;column:Id"`
	Data JSONObject `gorm:"type:text;column:Data"`
	AuditFields
}

type ItemInput struct {
	ID     *string  `json:"id"`
	Type   string   `json:"type"`
	Data   string   `json:"data"`
	Groups []string `json:"groups"`
}

type ItemResult struct {
	GenericResult
	Data *Item `json:"data"`
}

func GetItems(ctx context.Context, db *gorm.DB, s string) ([]Item, error) {
	panic("not implemented")
}

func GetItem(ctx context.Context, db *gorm.DB, s string) (*Item, error) {
	panic("not implemented")
}

func (item Item) Groups(ctx context.Context, db *gorm.DB) ([]Group, error) {
	panic("not implemented")
}

func UpsertItem(ctx context.Context, db *gorm.DB, item ItemInput, siteId string) (ItemResult, error) {
	panic("not implemented")
}

func DeleteItem(ctx context.Context, db *gorm.DB, itemId string) (GenericResult, error) {
	panic("not implemented")
}
