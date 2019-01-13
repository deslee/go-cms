package data

import (
	"context"
	"time"
)

type AuditFields struct {
	CreatedAt     string `db:"CreatedAt"`
	CreatedBy     string `db:"CreatedBy"`
	LastUpdatedAt string `db:"LastUpdatedAt"`
	LastUpdatedBy string `db:"LastUpdatedBy"`
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
	if len(userId) == 0 {
		modifier = "Unknown"
	} else {
		modifier = userId
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
