package data

import (
	"context"
	"database/sql"
	"time"
)

type Site struct {
	ID            string
	Name          string
	Data          JSONObject
	CreatedAt     time.Time
	CreatedBy     string
	LastUpdatedAt time.Time
	LastUpdatedBy string
}

type SiteInput struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
	Data string  `json:"data"`
}

type SiteResult struct {
	Data         *Site   `json:"data"`
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
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

func UpsertSite(ctx context.Context, db *sql.DB, site SiteInput) (SiteResult, error) {
	panic("not implemented")
}

func DeleteSite(ctx context.Context, db *sql.DB, siteId string) (GenericResult, error) {
	panic("not implemented")
}
