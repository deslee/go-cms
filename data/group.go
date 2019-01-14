package data

import (
	"context"
	. "github.com/deslee/cms/model"
	"github.com/jmoiron/sqlx"
)

func ItemsFromGroup(ctx context.Context, db *sqlx.DB, group Group) ([]Item, error) {
	panic("not implemented")
}
