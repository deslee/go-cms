package data

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Group struct {
	Id   string     `db:"Id"`
	Name string     `db:"Name"`
	Data JSONObject `db:"Data"`
	AuditFields
}

func (Group) TableName() string {
	return "Groups"
}

func (group Group) Items(ctx context.Context, db *sqlx.DB) ([]Item, error) {
	panic("not implemented")
}
