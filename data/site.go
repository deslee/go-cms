package data

import (
	"context"
	"database/sql"
)

type Site struct {
	ID            string
	Name          string
	Data          JSONObject
	AuditFields
}

type SiteInput struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
	Data string  `json:"data"`
}

type SiteResult struct {
	GenericResult
	Data         *Site   `json:"data"`
}

func (site Site) Items(ctx context.Context, db *sql.DB) ([]Item, error) {
	panic("not implemented")
}

func (site Site) Groups(ctx context.Context, db *sql.DB) ([]Group, error) {
	panic("not implemented")
}

func (site Site) Assets(ctx context.Context, db *sql.DB) ([]Asset, error) {
	panic("not implemented")
}

func GetSites(ctx context.Context, db *sql.DB) ([]Site, error) {
	panic("not implemented")
}

func GetSite(ctx context.Context, db *sql.DB, siteId string) (*Site, error) {
	panic("not implemented")
}

func DeleteSite(ctx context.Context, db *sql.DB, siteId string) (GenericResult, error) {
	panic("not implemented")
}


func UpsertSite(ctx context.Context, db *sql.DB, site SiteInput) (SiteResult, error) {
	panic("not implemented")
}