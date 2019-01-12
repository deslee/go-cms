package data

import (
	"context"
	"fmt"
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
	return getAllSitesForUserInContext(ctx, db)
}

func GetSite(ctx context.Context, db *gorm.DB, siteId string) (*Site, error) {
	if assertUserHasAccessToSite(ctx, db, siteId) == false {
		panic("Not authenticated")
	}
	return getSiteById(ctx, db, siteId)
}

func DeleteSite(ctx context.Context, db *gorm.DB, siteId string) (GenericResult, error) {
	if assertUserHasAccessToSite(ctx, db, siteId) == false {
		return GenericErrorMessage(UnauthenticatedMsg), nil
	}

	err := db.Delete(Site{}, "SiteId=?", siteId).Error
	if err != nil {
		return GenericErrorMessage(fmt.Sprintf("Error: %s", err)), nil
	}
	return GenericSuccess(), nil
}

func UpsertSite(ctx context.Context, db *gorm.DB, input SiteInput) (SiteResult, error) {
	var (
		site Site
	)
	user, err := UserFromContext(ctx, db)
	die(err)
	if user == nil {
		return SiteResult{GenericResult: GenericErrorMessage(UnauthenticatedMsg)}, nil
	}
	if input.ID == nil {
		site = Site{
			ID:          generateId(),
			Name:        input.Name,
			Data:        input.Data,
			AuditFields: CreateAuditFields(ctx, nil),
		}
		err = db.Create(site).Error
		if err != nil {
			return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Error: %s", err))}, nil
		}
	} else {
		existingSite, err := getSiteById(ctx, db, *input.ID)
		die(err)
		if existingSite == nil {
			return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Site %s does not exist", *input.ID))}, nil
		}
		site = Site{
			ID:          *input.ID,
			Name:        input.Name,
			Data:        input.Data,
			AuditFields: CreateAuditFields(ctx, &existingSite.AuditFields),
		}
		err = db.Update(site).Error
		if err != nil {
			return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Error: %s", err))}, nil
		}
	}

	siteUser := SiteUser{
		UserID:      user.ID,
		SiteID:      site.ID,
		Order:       0,
		AuditFields: CreateAuditFields(ctx, nil),
	}
	if err := db.Create(siteUser).Error; err != nil {
		if err != nil {
			return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Error: %s", err))}, nil
		}
	}

	existingSite, err := getSiteById(ctx, db, site.ID)
	if err != nil {
		return SiteResult{GenericResult: GenericErrorMessage(fmt.Sprintf("Error: %s", err))}, nil
	}

	return SiteResult{
		GenericResult: GenericSuccess(),
		Data:          existingSite,
	}, nil
}

func getAllSitesForUserInContext(ctx context.Context, db *gorm.DB) ([]Site, error) {
	var (
		sites []Site
	)

	user, err := UserFromContext(ctx, db)
	die(err)
	if user == nil {
		panic("Not authenticated")
	}

	rows, err := db.Raw(`select S.* from SiteUsers SU 
		INNER JOIN Sites S
		WHERE SU.UserId=?`, user.ID).Rows()
	die(err)
	defer rows.Close()
	for rows.Next() {
		var site Site
		err = db.ScanRows(rows, &site)
		die(err)
		sites = append(sites, site)
	}

	return sites, nil
}

func getSiteById(ctx context.Context, db *gorm.DB, id string) (*Site, error) {
	site := Site{}

	if err := db.Where("Id = ?", id).First(&site).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &site, nil
}
