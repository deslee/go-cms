//generated

package data

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func upsertSite(ctx context.Context, db *sqlx.DB, site Site) error {
	stmt, err := db.PrepareNamedContext(ctx, `
		INSERT INTO Sites VALUES (:Id,:Name,:Data,:CreatedBy,:LastUpdatedBy,:CreatedAt,:LastUpdatedAt)
		ON CONFLICT(Id) DO UPDATE SET 
		Name=excluded.Name,
		Data=excluded.Data,
		CreatedBy=excluded.CreatedBy,
		LastUpdatedBy=excluded.LastUpdatedBy,
		CreatedAt=excluded.CreatedAt,
		LastUpdatedAt=excluded.LastUpdatedAt
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(site)
	if err != nil {
		return err
	}

	return err
}

func upsertSiteUsers(ctx context.Context, db *sqlx.DB, siteUser SiteUser) error {
	stmt, err := db.PrepareNamedContext(ctx, `
		INSERT INTO SiteUsers VALUES (:UserId,:SiteId,:Order,:CreatedBy,:LastUpdatedBy,:CreatedAt,:LastUpdatedAt)
		ON CONFLICT(Id) DO UPDATE SET
		Order=excluded.Order,
		CreatedBy=excluded.CreatedBy,
		LastUpdatedBy=excluded.LastUpdatedBy,
		CreatedAt=excluded.CreatedAt,
		LastUpdatedAt=excluded.LastUpdatedAt
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(siteUser)
	if err != nil {
		return err
	}

	return err
}


/**
	Returns a user by email. User is nil if it doesn't exist.
 */
func getUserByEmail(ctx context.Context, db *sqlx.DB, email string) (*User, error) {
	user := User{}

	err := db.QueryRowx("SELECT * from Users WHERE Email = ?", email).StructScan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

/**
	Returns a user by id. User is nil if it doesn't exist.
 */
func getUserById(ctx context.Context, db *sqlx.DB, id string) (*User, error) {
	// initialize a zero value user
	user := User{}

	// select user
	err := db.QueryRowx("SELECT * from Users WHERE Id = :Id", id).StructScan(&user)
	if err != nil {
		// if the error is that no rows were returned, just return nil
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// otherwise, return an error
		return nil, err
	}

	// return a pointer to the user
	return &user, nil
}