package data

import (
	"context"
	. "github.com/deslee/cms/models"
	"github.com/jmoiron/sqlx"
)

func AddUserToSite(ctx context.Context, db *sqlx.DB, userId string, siteId string) (GenericResult, error) {
	siteUser := SiteUser{
		UserId:      userId,
		SiteId:      siteId,
		Order:       0,
		AuditFields: CreateAuditFields(ctx, nil),
	}

	err := RepoUpsertSiteUser(ctx, db, siteUser)
	if err != nil {
		return GenericUnexpectedError(err), nil
	}

	return GenericSuccess(), nil
}
