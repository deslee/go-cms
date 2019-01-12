package data

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
)

type SiteUser struct {
	UserID string `gorm:"type:text;primary_key;column:UserId"`
	SiteID string `gorm:"type:text;primary_key;column:SiteId"`
	Order  int    `gorm:"type:integer;column:Order"`
	AuditFields
}

func (SiteUser) TableName() string {
	return "SiteUsers"
}

func AddUserToSite(ctx context.Context, db *gorm.DB, userId string, siteId string) (GenericResult, error) {
	siteUser := SiteUser{
		UserID:      userId,
		SiteID:      siteId,
		Order:       0,
		AuditFields: CreateAuditFields(ctx, nil),
	}
	if err := db.Create(siteUser).Error; err != nil {
		return GenericErrorMessage(fmt.Sprintf("Failed: %s", err)), nil
	}
	return GenericSuccess(), nil
}
