package data

import (
	"context"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
)

func ItemsFromGroup(ctx context.Context, db *sqlx.DB, group Group) ([]Item, error) {
	return repository.ScanItemList(ctx, db,"SELECT I.* FROM ItemGroups IG INNER JOIN Items I on IG.ItemId = I.Id WHERE IG.GroupId=?", group.Id)
}
