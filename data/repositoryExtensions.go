// TODO: make this generated as well
// TODO: tags need to be richer
// Tag settings: createGetBy,Pk,columnName

package data

import (
	"context"
	"database/sql"
	. "github.com/deslee/cms/models"
	"github.com/jmoiron/sqlx"
)

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
