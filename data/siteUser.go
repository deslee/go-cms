package data

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type SiteUser struct {
	UserId string `db:"UserId"`
	SiteId string `db:"SiteId"`
	Order  int    `db:"Order"`
	AuditFields
}

func (SiteUser) TableName() string {
	return "SiteUsers"
}

func AddUserToSite(ctx context.Context, db *sqlx.DB, userId string, siteId string) (GenericResult, error) {
	siteUser := SiteUser{
		UserId:      userId,
		SiteId:      siteId,
		Order:       0,
		AuditFields: CreateAuditFields(ctx, nil),
	}

	err := upsertSiteUsers(ctx, db, siteUser)
	if err != nil {
		return GenericUnexpectedError(err), nil
	}

	return GenericSuccess(), nil
}

