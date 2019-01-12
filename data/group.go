package data

import (
	"context"
	"database/sql"
)

type Group struct {
	ID            string
	Name          string
	AuditFields
}

func (group Group) Items(context context.Context, db *sql.DB) ([]Item, error) {
	panic("not implemented")
}
