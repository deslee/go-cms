package models

type User struct {
	ID            string
	Email         string
	Password      string
	Salt          string
	Data          string
	CreatedAt     string
	CreatedBy     string
	LastUpdatedAt string
	LastUpdatedBy string
}