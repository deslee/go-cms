package data

import (
	"context"
	"database/sql"
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

func UnexpectedErrorSiteResult(err error) SiteResult {
	return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Unexpected error %s", err))}
}

func ItemsFromSite(ctx context.Context, db *sqlx.DB, site Site) ([]Item, error) {
	panic("not implemented")
}

func GroupsFromSite(ctx context.Context, db *sqlx.DB, site Site) ([]Group, error) {
	panic("not implemented")
}

func AssetsFromSite(FromSitectx context.Context, db *sqlx.DB, site Site) ([]Asset, error) {
	panic("not implemented")
}

func GetSites(ctx context.Context, db *sqlx.DB) ([]Site, error) {
	return getAllSitesForUserInContext(ctx, db)
}

func GetSite(ctx context.Context, db *sqlx.DB, siteId string) (*Site, error) {
	if validated, err := assertUserHasAccessToSite(ctx, db, siteId); validated == false || err != nil {
		if err != nil {
			log.Printf("%s", err)
		}
		return nil, nil
	}
	return getSiteById(ctx, db, siteId)
}

func DeleteSite(ctx context.Context, db *sqlx.DB, siteId string) (GenericResult, error) {
	if validated, err := assertUserHasAccessToSite(ctx, db, siteId); validated == false || err != nil {
		if err != nil {
			log.Printf("%s", err)
		}
		return GenericErrorMessage(UnauthenticatedMsg), nil
	}

	_, err := db.Exec("DELETE FROM Sites WHERE Id=?", siteId)
	if err != nil {
		return GenericErrorMessage(fmt.Sprintf("Error: %s", err)), nil
	}
	return GenericSuccess(), nil
}

func UpsertSite(ctx context.Context, db *sqlx.DB, input SiteInput) (SiteResult, error) {
	var (
		site Site
	)

	// start transaction
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	// get the current logged in user
	user, err := UserFromContext(ctx, db)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}
	if user == nil {
		return SiteResult{GenericResult: GenericErrorMessage(UnauthenticatedMsg)}, nil
	}

	if input.ID == nil {
		// if we are creating, just generate an id
		site = Site{
			Id:          generateId(),
			Name:        input.Name,
			Data:        input.Data,
			AuditFields: CreateAuditFields(ctx, nil),
		}
	} else {
		// otherwise, we need to do some validations...

		// validate that the user has access to the site
		existingSite, err := getSiteById(ctx, db, *input.ID)
		if err != nil {
			return UnexpectedErrorSiteResult(err), nil
		}
		if existingSite == nil {
			return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Site %s does not exist", *input.ID))}, nil
		}
		validated, err := assertUserHasAccessToSite(ctx, db, existingSite.Id)
		if err != nil {
			return UnexpectedErrorSiteResult(err), nil
		}
		if validated == false {
			return SiteResult{GenericResult: GenericErrorMessage(UnauthenticatedMsg)}, nil
		}
		site = Site{
			Id:          *input.ID,
			Name:        input.Name,
			Data:        input.Data,
			AuditFields: CreateAuditFields(ctx, &existingSite.AuditFields),
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
		AuditFields: CreateAuditFields(ctx, nil),
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
	existingSite, err := getSiteById(ctx, db, site.Id)
	if err != nil {
		return UnexpectedErrorSiteResult(err), nil
	}

	return SiteResult{
		GenericResult: GenericSuccess(),
		Data:          existingSite,
	}, nil
}

func getAllSitesForUserInContext(ctx context.Context, db *sqlx.DB) ([]Site, error) {
	var (
		sites []Site
	)

	user, err := UserFromContext(ctx, db)
	if user == nil {
		return sites, nil
	}

	rows, err := db.Queryx(`SELECT S.* FROM SiteUsers SU INNER JOIN Sites S ON S.Id=SU.SiteId WHERE SU.UserId=?`, user.Id)
	if err != nil {
		fmt.Printf("Failed to query SiteUsers: %s", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var site Site
		err = rows.StructScan(&site)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	return sites, nil
}

func getSiteById(ctx context.Context, db *sqlx.DB, id string) (*Site, error) {
	site := Site{}

	row := db.QueryRowx("SELECT  * FROM Sites WHERE Id=?", id)
	err := row.StructScan(&site)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &site, nil
}
