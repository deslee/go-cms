package repository

import (
	"context"
	"database/sql"
	models "github.com/deslee/cms/models"
	sqlx "github.com/jmoiron/sqlx"
)

func UpsertUser(ctx context.Context, db *sqlx.DB, obj models.User) error {
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

func FindUserById(ctx context.Context, db *sqlx.DB, val string) (*models.User, error) {
	obj := models.User{}

	err := db.QueryRowx("SELECT * FROM Users WHERE Id=?", val).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertSite(ctx context.Context, db *sqlx.DB, obj models.Site) error {
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

func FindSiteById(ctx context.Context, db *sqlx.DB, val string) (*models.Site, error) {
	obj := models.Site{}

	err := db.QueryRowx("SELECT * FROM Sites WHERE Id=?", val).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertItem(ctx context.Context, db *sqlx.DB, obj models.Item) error {
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

func FindItemById(ctx context.Context, db *sqlx.DB, val string) (*models.Item, error) {
	obj := models.Item{}

	err := db.QueryRowx("SELECT * FROM Items WHERE Id=?", val).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertGroup(ctx context.Context, db *sqlx.DB, obj models.Group) error {
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

func FindGroupById(ctx context.Context, db *sqlx.DB, val string) (*models.Group, error) {
	obj := models.Group{}

	err := db.QueryRowx("SELECT * FROM Groups WHERE Id=?", val).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertAsset(ctx context.Context, db *sqlx.DB, obj models.Asset) error {
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

func FindAssetById(ctx context.Context, db *sqlx.DB, val string) (*models.Asset, error) {
	obj := models.Asset{}

	err := db.QueryRowx("SELECT * FROM Assets WHERE Id=?", val).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertSiteUser(ctx context.Context, db *sqlx.DB, obj models.SiteUser) error {
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

func UpsertItemGroup(ctx context.Context, db *sqlx.DB, obj models.ItemGroup) error {
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

func UpsertItemAsset(ctx context.Context, db *sqlx.DB, obj models.ItemAsset) error {
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
