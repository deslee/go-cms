package data

import (
	"context"
	"fmt"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
	"log"
)

type SiteInput struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
	Data string  `json:"data"`
}

type SiteResult struct {
	GenericResult
	Data *Site `json:"data"`
}

func ItemsFromSite(ctx context.Context, db *sqlx.DB, site Site) ([]Item, error) {
	return repository.ScanItemList(ctx, db, "SELECT I.* FROM Items I WHERE I.SiteId=?", site.Id)
}

func GroupsFromSite(ctx context.Context, db *sqlx.DB, site Site) ([]Group, error) {
	return repository.ScanGroupList(ctx, db, "SELECT G.* FROM Groups G Where G.SiteId=?", site.Id)
}

func AssetsFromSite(ctx context.Context, db *sqlx.DB, site Site) ([]Asset, error) {
	return repository.ScanAssetList(ctx, db, `SELECT A.* from Assets A WHERE A.SiteId=?`, site.Id)
}

func QueryGetSites(ctx context.Context, db *sqlx.DB) ([]Site, error) {
	return getAllSitesForUserInContext(ctx, db)
}

func QueryGetSite(ctx context.Context, db *sqlx.DB, siteId string) (*Site, error) {
	validated, err := AssertContextUserHasAccessToSite(ctx, db, siteId)
	if err != nil {
		return nil, err
	}
	if validated == false {
		return nil, nil
	}

	site, err := repository.FindSiteById(ctx, db, siteId)
	if err != nil {
		return nil, err
	}
	return site, nil
}

func MutationDeleteSite(ctx context.Context, db *sqlx.DB, siteId string) (GenericResult, error) {
	if validated, err := AssertContextUserHasAccessToSite(ctx, db, siteId); validated == false || err != nil {
		if err != nil {
			log.Printf("%s", err)
		}
		return ErrorGenericResult(UnauthenticatedMsg), nil
	}

	err := repository.DeleteSiteById(ctx, db, siteId)
	if err != nil {
		return ErrorGenericResult(fmt.Sprintf("Error: %s", err)), nil
	}
	return GenericSuccess(), nil
}

func MutationUpsertSite(ctx context.Context, db *sqlx.DB, input SiteInput) (SiteResult, error) {
	var (
		site Site
	)

	// start transaction
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	// get the current logged in user
	user, err := QueryGetCurrentUser(ctx, db)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}
	if user == nil {
		return SiteResult{GenericResult: ErrorGenericResult(UnauthenticatedMsg)}, nil
	}

	if input.ID == nil {
		// if we are creating, just generate an id
		site = Site{
			Id:          GenerateId(),
			Name:        input.Name,
			Data:        input.Data,
			AuditFields: CreateAuditFields(ctx),
		}
	} else {
		// otherwise, we need to do some validations...

		// validate that the user has access to the site
		existingSite, err := repository.FindSiteById(ctx, db, *input.ID)
		if err != nil {
			return UnexpectedErrorSiteResult(err), nil
		}
		if existingSite == nil {
			return SiteResult{GenericResult: ErrorGenericResult(fmt.Sprintf("Site %s does not exist", *input.ID))}, nil
		}
		validated, err := AssertContextUserHasAccessToSite(ctx, db, existingSite.Id)
		if err != nil {
			return UnexpectedErrorSiteResult(err), nil
		}
		if validated == false {
			return SiteResult{GenericResult: ErrorGenericResult(UnauthenticatedMsg)}, nil
		}
		site = Site{
			Id:          *input.ID,
			Name:        input.Name,
			Data:        input.Data,
			AuditFields: CreateAuditFieldsFromExisting(ctx, existingSite.AuditFields),
		}
	}

	// upsert the site
	err = repository.UpsertSiteTx(ctx, tx, site)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	// associate the user with the site
	siteUser := SiteUser{
		UserId:      user.Id,
		SiteId:      site.Id,
		Order:       0,
		AuditFields: CreateAuditFields(ctx),
	}
	err = repository.UpsertSiteUserTx(ctx, tx, siteUser)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	// get the site back out to return to the user
	existingSite, err := repository.FindSiteById(ctx, db, site.Id)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	return SiteResult{
		GenericResult: GenericSuccess(),
		Data:          existingSite,
	}, nil
}
