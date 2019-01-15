package data

import (
	"context"
	. "github.com/deslee/cms/model"
	"time"
)

func CreateAuditFields(ctx context.Context) AuditFields {
	var (
		createdAt     string
		lastUpdatedAt string
		createdBy     string
		lastUpdatedBy string
		modifier      string
		now           = time.Now().UTC().Format(time.RFC3339)
	)

	userId := UserIdFromContext(ctx)
	if len(userId) == 0 {
		modifier = "Unknown"
	} else {
		modifier = userId
	}

	createdAt = now
	createdBy = modifier
	lastUpdatedBy = modifier
	lastUpdatedAt = now

	return AuditFields{
		CreatedBy:     createdBy,
		CreatedAt:     createdAt,
		LastUpdatedBy: lastUpdatedBy,
		LastUpdatedAt: lastUpdatedAt,
	}
}

func CreateAuditFieldsFromExisting(ctx context.Context, previous AuditFields) AuditFields {
	var (
		createdAt     string
		lastUpdatedAt string
		createdBy     string
		lastUpdatedBy string
		modifier      string
		now           = time.Now().UTC().Format(time.RFC3339)
	)

	userId := UserIdFromContext(ctx)
	if len(userId) == 0 {
		modifier = "Unknown"
	} else {
		modifier = userId
	}

	createdAt = previous.CreatedAt
	createdBy = previous.CreatedBy

	lastUpdatedBy = modifier
	lastUpdatedAt = now

	return AuditFields{
		CreatedBy:     createdBy,
		CreatedAt:     createdAt,
		LastUpdatedBy: lastUpdatedBy,
		LastUpdatedAt: lastUpdatedAt,
	}
}
