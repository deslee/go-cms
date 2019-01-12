package data

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID            string
	Email         string
	Password      string
	Salt          string
	Data          JSONObject
	CreatedAt     time.Time
	CreatedBy     string
	LastUpdatedAt time.Time
	LastUpdatedBy string
}

func (user User) Sites(ctx context.Context, db *sql.DB) ([]Site, error) {
	panic("not implemented")
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	Data         *User   `json:"data"`
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
	Token        *string `json:"token"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Data     string `json:"data"`
	Password string `json:"password"`
}

type UserInput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Data  string `json:"data"`
}

type UserResult struct {
	Data         *User   `json:"data"`
	ErrorMessage *string `json:"errorMessage"`
	Success      bool    `json:"success"`
}

func UserFromContext(ctx context.Context, db *sql.DB) (*User, error) {
	panic("not implemented")
}

func UpdateUser(ctx context.Context, db *sql.DB, user UserInput) (UserResult, error) {
	panic("not implemented")
}

func Register(ctx context.Context, db *sql.DB, registration RegisterInput) (UserResult, error) {
	panic("not implemented")
}

func Login(ctx context.Context, db *sql.DB, login LoginInput) (LoginResult, error) {
	panic("not implemented")
}

func AddUserToSite(ctx context.Context, db *sql.DB, userId string, siteId string) (GenericResult, error) {
	panic("not implemented")
}
