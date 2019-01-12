package data

import (
	"context"
	"time"
)

type AuditFields struct {
	CreatedAt     string `gorm:"type:text;column:CreatedAt"`
	CreatedBy     string `gorm:"type:text;column:CreatedBy"`
	LastUpdatedAt string `gorm:"type:text;column:LastUpdatedAt"`
	LastUpdatedBy string `gorm:"type:text;column:LastUpdatedBy"`
}

func CreateAuditFields(ctx context.Context, previous *AuditFields) AuditFields {
	var (
		createdAt     string
		lastUpdatedAt string
		createdBy     string
		lastUpdatedBy string
		modifier      string
		now           = time.Now().UTC().Format(time.RFC3339)
	)

	userId := UserIdFromContext(ctx)
	if userId == nil {
		modifier = "Unknown"
	} else {
		modifier = *userId
	}

	if previous != nil {
		createdAt = previous.CreatedAt
		createdBy = previous.CreatedBy
	} else {
		createdAt = now
		createdBy = modifier
	}

	lastUpdatedBy = modifier
	lastUpdatedAt = now

	return AuditFields{
		CreatedBy:     createdBy,
		CreatedAt:     createdAt,
		LastUpdatedBy: lastUpdatedBy,
		LastUpdatedAt: lastUpdatedAt,
	}
}
