package data

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Group struct {
	ID          string     `gorm:"type:text;primary_key;column:Id"`
	Name        string     `gorm:"type:text;column:Name"`
	Data        JSONObject `gorm:"type:text;column:Data"`
	AuditFields `gorm:"type:text"`
}

func (group Group) Items(ctx context.Context, db *gorm.DB) ([]Item, error) {
	panic("not implemented")
}
