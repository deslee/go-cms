package data

import (
	"context"
	"time"
)

type AuditFields struct {
	CreatedAt     string
	CreatedBy     string
	LastUpdatedAt string
	LastUpdatedBy string
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

	modifier = "System" // TODO

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
