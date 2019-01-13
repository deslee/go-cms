package data

import (
	"context"
	"database/sql"
	sqlx "github.com/jmoiron/sqlx"
)

func RepoUpsertUser(ctx context.Context, db *sqlx.DB, obj User) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO Users VALUES (:Id,:Email,:Password,:Salt,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Email`=excluded.`Email`,`Password`=excluded.`Password`,`Salt`=excluded.`Salt`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindUserById(ctx context.Context, db *sqlx.DB, id string) (*User, error) {
	obj := User{}

	err := db.QueryRowx("SELECT * FROM Users WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertSite(ctx context.Context, db *sqlx.DB, obj Site) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO Sites VALUES (:Id,:Name,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Name`=excluded.`Name`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindSiteById(ctx context.Context, db *sqlx.DB, id string) (*Site, error) {
	obj := Site{}

	err := db.QueryRowx("SELECT * FROM Sites WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertItem(ctx context.Context, db *sqlx.DB, obj Item) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO Items VALUES (:Id,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindItemById(ctx context.Context, db *sqlx.DB, id string) (*Item, error) {
	obj := Item{}

	err := db.QueryRowx("SELECT * FROM Items WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertGroup(ctx context.Context, db *sqlx.DB, obj Group) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO Groups VALUES (:Id,:Name,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Name`=excluded.`Name`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindGroupById(ctx context.Context, db *sqlx.DB, id string) (*Group, error) {
	obj := Group{}

	err := db.QueryRowx("SELECT * FROM Groups WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertAsset(ctx context.Context, db *sqlx.DB, obj Asset) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO Assets VALUES (:Id,:State,:Type,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`State`=excluded.`State`,`Type`=excluded.`Type`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindAssetById(ctx context.Context, db *sqlx.DB, id string) (*Asset, error) {
	obj := Asset{}

	err := db.QueryRowx("SELECT * FROM Assets WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertSiteUser(ctx context.Context, db *sqlx.DB, obj SiteUser) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO SiteUsers VALUES (:UserId,:SiteId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `UserId`=excluded.`UserId`,`SiteId`=excluded.`SiteId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindSiteUserById(ctx context.Context, db *sqlx.DB, id string) (*SiteUser, error) {
	obj := SiteUser{}

	err := db.QueryRowx("SELECT * FROM SiteUsers WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertItemGroup(ctx context.Context, db *sqlx.DB, obj ItemGroup) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO ItemGroups VALUES (:ItemId,:GroupId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `ItemId`=excluded.`ItemId`,`GroupId`=excluded.`GroupId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindItemGroupById(ctx context.Context, db *sqlx.DB, id string) (*ItemGroup, error) {
	obj := ItemGroup{}

	err := db.QueryRowx("SELECT * FROM ItemGroups WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func RepoUpsertItemAsset(ctx context.Context, db *sqlx.DB, obj ItemAsset) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO ItemAssets VALUES (:ItemId,:AssetId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `ItemId`=excluded.`ItemId`,`AssetId`=excluded.`AssetId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func RepoFindItemAssetById(ctx context.Context, db *sqlx.DB, id string) (*ItemAsset, error) {
	obj := ItemAsset{}

	err := db.QueryRowx("SELECT * FROM ItemAssets WHERE Id=?", id).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}
