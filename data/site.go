package data

import (
	"context"
	"github.com/jinzhu/gorm"
)

type Site struct {
	ID   string     `gorm:"type:text;primary_key;column:Id"`
	Name string     `gorm:"type:text;column:Name"`
	Data JSONObject `gorm:"type:text;column:Data"`
	AuditFields
}

type SiteInput struct {
	ID   *string `json:"id"`
	Name string  `json:"name"`
	Data string  `json:"data"`
}

type SiteResult struct {
	GenericResult
	Data *Site `json:"data"`
}

func (site Site) Items(ctx context.Context, db *gorm.DB) ([]Item, error) {
	panic("not implemented")
}

func (site Site) Groups(ctx context.Context, db *gorm.DB) ([]Group, error) {
	panic("not implemented")
}

func (site Site) Assets(ctx context.Context, db *gorm.DB) ([]Asset, error) {
	panic("not implemented")
}

func GetSites(ctx context.Context, db *gorm.DB) ([]Site, error) {
	panic("not implemented")
}

func GetSite(ctx context.Context, db *gorm.DB, siteId string) (*Site, error) {
	panic("not implemented")
}

func DeleteSite(ctx context.Context, db *gorm.DB, siteId string) (GenericResult, error) {
	panic("not implemented")
}

func UpsertSite(ctx context.Context, db *gorm.DB, site SiteInput) (SiteResult, error) {
	user, err := UserFromContext(ctx, db)
	die(err)
	if user == nil {
		return SiteResult{GenericResult: GenericErrorMessage(UnauthenticatedMsg)}, nil
	}

	panic("not implemented")
}
