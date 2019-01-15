package data

import (
	"context"
	"fmt"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func ItemsFromAsset(ctx context.Context, db *sqlx.DB, asset Asset) ([]Item, error) {
	return repository.ScanItemList(ctx, db, "SELECT I.* FROM ItemAssets IA INNER JOIN Items I on IA.ItemId = I.Id WHERE IA.AssetId=?", asset.Id)
}

func QueryGetAsset(ctx context.Context, db *sqlx.DB, assetId string) (*Asset, error) {
	// find the asset
	asset, err := repository.FindAssetById(ctx, db, assetId)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, nil
	}

	// validate
	hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, asset.SiteId)
	if err != nil {
		return nil, err
	}
	userId := UserIdFromContext(ctx)
	if hasAccess == false {
		return nil, errors.Errorf("User %s does not have access to asset %s", userId, assetId)
	}

	return asset, nil
}

func MutationDeleteAsset(ctx context.Context, db *sqlx.DB, assetId string) (GenericResult, error) {
	// find the asset
	asset, err := repository.FindAssetById(ctx, db, assetId)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}
	if asset == nil {
		return ErrorGenericResult(fmt.Sprintf("Asset %s does not exist", assetId)), nil
	}

	// validate
	hasAccess, err := AssertContextUserHasAccessToSite(ctx, db, asset.SiteId)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}
	userId := UserIdFromContext(ctx)
	if hasAccess == false {
		return ErrorGenericResult(fmt.Sprintf("User %s does not have access to asset %s", userId, assetId)), nil
	}

	// delete
	err = repository.DeleteAssetById(ctx, db, assetId)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}

	return GenericSuccess(), nil
}
