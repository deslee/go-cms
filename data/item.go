package data

import (
	"context"
	"fmt"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"strings"
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

func QueryGetItems(ctx context.Context, db *sqlx.DB, siteId string) ([]Item, error) {
	// validate
	hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, siteId)
	if err != nil {
		return nil, err
	}
	userId := UserIdFromContext(ctx)
	if hasAccess == false {
		return nil, errors.Errorf("User %s does not have access to site %s", userId, siteId)
	}

	// get all items
	items, err := repository.ScanItemList(ctx, db, "SELECT I.* FROM Items I WHERE I.SiteId=?", siteId)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func QueryGetItem(ctx context.Context, db *sqlx.DB, itemId string) (*Item, error) {
	// get the item
	item, err := repository.FindItemById(ctx, db, itemId)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, nil
	}

	// validate
	hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, item.SiteId)
	if err != nil {
		return nil, err
	}
	userId := UserIdFromContext(ctx)
	if hasAccess == false {
		return nil, errors.Errorf("User %s does not have access to item %s", userId, itemId)
	}

	return item, nil
}

func MutationDeleteItem(ctx context.Context, db *sqlx.DB, itemId string) (GenericResult, error) {
	// find the item
	item, err := repository.FindItemById(ctx, db, itemId)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}
	if item == nil {
		return ErrorGenericResult(fmt.Sprintf("Item %s does not exist", itemId)), nil
	}

	// validate
	hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, item.SiteId)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}
	userId := UserIdFromContext(ctx)
	if hasAccess == false {
		return ErrorGenericResult(fmt.Sprintf("User %s does not have access to item %s", userId, itemId)), nil
	}

	// delete
	err = repository.DeleteItemById(ctx, db, itemId)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}

	return GenericSuccess(), nil
}

func GroupsFromItem(ctx context.Context, db *sqlx.DB, item Item) ([]Group, error) {
	return repository.ScanGroupList(ctx, db, "SELECT G.* FROM ItemGroups IG INNER JOIN Groups G on IG.GroupId = G.Id WHERE IG.ItemId=?", item.Id)
}

func MutationUpsertItem(ctx context.Context, db *sqlx.DB, input ItemInput, siteId string) (ItemResult, error) {
	var (
		item Item
	)

	// validate
	hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, siteId)
	if err != nil {
		return UnexpectedErrorItemResult(err), nil
	}
	userId := UserIdFromContext(ctx)
	if hasAccess == false {
		return ErrorItemResult(fmt.Sprintf("User %s does not have access to site %s", userId, siteId)), nil
	}

	// start transaction
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return UnexpectedErrorItemResult(err), nil
	}

	if input.Id == nil {
		// generate an id
		item = Item{
			Id:          GenerateId(),
			Type:        input.Type,
			Data:        input.Data,
			SiteId:      siteId,
			AuditFields: CreateAuditFields(ctx),
		}
	} else {
		// we need to validate
		// get existing item
		existingItem, err := repository.FindItemById(ctx, db, *input.Id)
		if err != nil {
			return UnexpectedErrorItemResult(err), nil
		}
		if existingItem == nil {
			return ErrorItemResult(fmt.Sprintf("Item %s does not exist", *input.Id)), nil
		}
		// validate the item belongs to the right site
		if existingItem.SiteId != siteId {
			return ErrorItemResult(fmt.Sprintf("Item %s does not belong to site %s", existingItem.Id, siteId)), nil
		}

		// validate user has access
		hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, existingItem.SiteId)
		if err != nil {
			return UnexpectedErrorItemResult(err), nil
		}
		userId := UserIdFromContext(ctx)
		if hasAccess == false {
			return ErrorItemResult(fmt.Sprintf("User %s does not have access to item %s", userId, existingItem.Id)), nil
		}

		item = Item{
			Id:          existingItem.Id,
			Type:        input.Type,
			Data:        input.Data,
			SiteId:      siteId,
			AuditFields: CreateAuditFieldsFromExisting(ctx, existingItem.AuditFields),
		}

		// remove groups in item that are not in payload
		// algorithm:
		// 1. get all the groups in the existing item
		// 2. loop through each group of 1).
		// 3. if the group does not exist by name in the input, delete it

		// construct a map of all the groups by name
		mapInputGroups := make(map[string]bool)
		for _, g := range input.Groups {
			mapInputGroups[strings.ToLower(g)] = true
		}

		existingGroups, err := repository.ScanGroupList(ctx, db, "SELECT G.Id FROM ItemGroups IG INNER JOIN Groups G on IG.GroupId = G.Id WHERE IG.ItemId = ?", item.Id)
		if err != nil {
			return UnexpectedErrorItemResult(err), nil
		}

		// do the algorithm
		for _, g := range existingGroups {
			if mapInputGroups[strings.ToLower(g.Name)] == false {
				err = repository.DeleteItemGroupByItemIdAndGroupIdTx(ctx, tx, item.Id, g.Id)
				if err != nil {
					return UnexpectedErrorItemResult(err), nil
				}
			}
		}
	}

	// upsert the item
	err = repository.UpsertItemTx(ctx, tx, item)
	if err != nil {
		return UnexpectedErrorItemResult(err), nil
	}

	// upsert all the groups in the input
	for order, g := range input.Groups {
		var itemGroup ItemGroup
		// check if the group exists
		igs, err := repository.ScanItemGroupList(ctx, db, "SELECT IG.* FROM ItemGroups IG INNER JOIN Groups G ON IG.GroupId = G.Id WHERE IG.ItemId=? AND G.SiteId=? AND G.Name=?", item.Id, siteId, strings.ToLower(g))
		if err != nil {
			return UnexpectedErrorItemResult(err), nil
		}

		if len(igs) > 0 {
			// if the group exists already, just update the order
			itemGroup = ItemGroup{
				ItemId:      igs[0].ItemId,
				GroupId:     igs[0].GroupId,
				Order:       order,
				AuditFields: CreateAuditFieldsFromExisting(ctx, igs[0].AuditFields),
			}
		} else {
			// otherwise, make a new one
			group := Group{
				Id:          GenerateId(),
				SiteId:      siteId,
				Name:        g,
				Data:        "{}",
				AuditFields: CreateAuditFields(ctx),
			}

			err := repository.UpsertGroupTx(ctx, tx, group)
			if err != nil {
				return UnexpectedErrorItemResult(err), nil
			}

			itemGroup = ItemGroup{
				ItemId:      item.Id,
				GroupId:     group.Id,
				Order:       order,
				AuditFields: CreateAuditFields(ctx),
			}
		}

		err = repository.UpsertItemGroupTx(ctx, tx, itemGroup)
		if err != nil {
			return UnexpectedErrorItemResult(err), nil
		}
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return UnexpectedErrorItemResult(err), nil
	}

	// get the item back out to return to the user
	existingItem, err := repository.FindItemById(ctx, db, item.Id)
	if err != nil {
		return UnexpectedErrorItemResult(err), nil
	}

	// remove groups from site that no longer have items after transaction
	err = deleteGroupsFromSiteWithNoItems(ctx, db)
	if err != nil {
		return UnexpectedErrorItemResult(err), nil
	}

	return ItemResult{
		GenericResult: GenericSuccess(),
		Data:          existingItem,
	}, nil
}

func deleteGroupsFromSiteWithNoItems(ctx context.Context, db *sqlx.DB) error {
	// start transaction
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	groupsToDelete, err := repository.ScanGroupList(ctx, db, "SELECT G.* FROM Groups G LEFT JOIN ItemGroups IG ON G.Id = IG.GroupId WHERE IG.GroupId IS NULL")
	if err != nil {
		return err
	}
	for _, groupToDelete := range groupsToDelete {
		err = repository.DeleteGroupByIdTx(ctx, tx, groupToDelete.Id)
		if err != nil {
			return err
		}
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
