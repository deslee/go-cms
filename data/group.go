package data

import (
	"context"
	"database/sql"
	"time"
)

type Group struct {
	ID            string
	Name          string
	CreatedAt     time.Time
	CreatedBy     string
	LastUpdatedAt time.Time
	LastUpdatedBy string
}

func (group Group) Items(context context.Context, db *sql.DB) ([]Item, error) {
	panic("not implemented")
}
