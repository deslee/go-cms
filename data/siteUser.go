package data

import (
	"context"
	"github.com/jinzhu/gorm"
)

type SiteUser struct {
	UserID string `gorm:"type:text;primary_key;column=UserId"`
	SiteID string `gorm:"type:text;primary_key;column=SiteId"`
	Order  int    `gorm:"type:integer;column=Order"`
	AuditFields
}

func AddUserToSite(ctx context.Context, db *gorm.DB, userId string, siteId string) (GenericResult, error) {
	panic("not implemented")
}
