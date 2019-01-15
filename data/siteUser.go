package data

import (
	"context"
	"errors"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/jmoiron/sqlx"
)

func MutationAddUserToSite(ctx context.Context, db *sqlx.DB, userId string, siteId string) (GenericResult, error) {
	siteUser := SiteUser{
		UserId:      userId,
		SiteId:      siteId,
		Order:       0,
		AuditFields: CreateAuditFields(ctx),
	}

	err := repository.UpsertSiteUser(ctx, db, siteUser)
	if err != nil {
		return UnexpectedErrorGenericResult(err), nil
	}

	return GenericSuccess(), nil
}

func getAllSitesForUserInContext(ctx context.Context, db *sqlx.DB) ([]Site, error) {
	user, err := QueryGetCurrentUser(ctx, db)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("unauthorized")
	}

	return SitesFromUser(ctx, db, *user)
}
