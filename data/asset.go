package data

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Asset struct {
	ID    string `gorm:"type:text;primary_key;column:Id"`
	State string `gorm:"type:text;column:State"`
	Type  string `gorm:"type:text;column:Type"`
	Data  string `gorm:"type:text;column:Data"`
	AuditFields
}

func (asset Asset) FileName() string {
	panic("not implemented")
}

func (asset Asset) Extension() string {
	panic("not implemented")
}

func (asset Asset) Items(ctx context.Context, db *gorm.DB) ([]Item, error) {
	panic("not implemented")
}

func GetAsset(ctx context.Context, db *gorm.DB, s string) (*Asset, error) {
	panic("not implemented")
}

func DeleteAsset(ctx context.Context, db *gorm.DB, assetId string) (GenericResult, error) {
	panic("not implemented")
}
