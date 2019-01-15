package data

import (
	"context"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
)

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

func QueryGetItems(ctx context.Context, db *sqlx.DB, s string) ([]Item, error) {
	panic("not implemented")
}

func QueryGetItem(ctx context.Context, db *sqlx.DB, s string) (*Item, error) {
	panic("not implemented")
}

func MutationUpsertItem(ctx context.Context, db *sqlx.DB, item ItemInput, siteId string) (ItemResult, error) {
	panic("not implemented")
}

func MutationDeleteItem(ctx context.Context, db *sqlx.DB, itemId string) (GenericResult, error) {
	panic("not implemented")
}

func GroupsFromItem(ctx context.Context, db *sqlx.DB, item Item) ([]Group, error) {
	return repository.ScanGroupList(ctx, db, "SELECT G.* FROM ItemGroups IG INNER JOIN Groups G on IG.GroupId = G.Id WHERE IG.ItemId=?", item.Id)
}
