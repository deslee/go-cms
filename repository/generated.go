// Code generated by genRepo.go, DO NOT EDIT.

package repository

import (
	"context"
	"database/sql"
	model "github.com/deslee/cms/model"
	sqlx "github.com/jmoiron/sqlx"
)

func UpsertUser(ctx context.Context, db *sqlx.DB, obj model.User) error {
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

func UpsertUserTx(ctx context.Context, tx *sqlx.Tx, obj model.User) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO Users VALUES (:Id,:Email,:Password,:Salt,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Email`=excluded.`Email`,`Password`=excluded.`Password`,`Salt`=excluded.`Salt`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindUserById(ctx context.Context, db *sqlx.DB, keyId string) (*model.User, error) {
	obj := model.User{}

	err := db.QueryRowx("SELECT * FROM Users WHERE Id=?", keyId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func FindUserByEmail(ctx context.Context, db *sqlx.DB, keyEmail string) (*model.User, error) {
	obj := model.User{}

	err := db.QueryRowx("SELECT * FROM Users WHERE Email=?", keyEmail).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertSite(ctx context.Context, db *sqlx.DB, obj model.Site) error {
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

func UpsertSiteTx(ctx context.Context, tx *sqlx.Tx, obj model.Site) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO Sites VALUES (:Id,:Name,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Name`=excluded.`Name`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindSiteById(ctx context.Context, db *sqlx.DB, keyId string) (*model.Site, error) {
	obj := model.Site{}

	err := db.QueryRowx("SELECT * FROM Sites WHERE Id=?", keyId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertItem(ctx context.Context, db *sqlx.DB, obj model.Item) error {
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

func UpsertItemTx(ctx context.Context, tx *sqlx.Tx, obj model.Item) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO Items VALUES (:Id,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindItemById(ctx context.Context, db *sqlx.DB, keyId string) (*model.Item, error) {
	obj := model.Item{}

	err := db.QueryRowx("SELECT * FROM Items WHERE Id=?", keyId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertGroup(ctx context.Context, db *sqlx.DB, obj model.Group) error {
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

func UpsertGroupTx(ctx context.Context, tx *sqlx.Tx, obj model.Group) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO Groups VALUES (:Id,:Name,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`Name`=excluded.`Name`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindGroupById(ctx context.Context, db *sqlx.DB, keyId string) (*model.Group, error) {
	obj := model.Group{}

	err := db.QueryRowx("SELECT * FROM Groups WHERE Id=?", keyId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertAsset(ctx context.Context, db *sqlx.DB, obj model.Asset) error {
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

func UpsertAssetTx(ctx context.Context, tx *sqlx.Tx, obj model.Asset) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO Assets VALUES (:Id,:State,:Type,:Data,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(Id) DO UPDATE SET `Id`=excluded.`Id`,`State`=excluded.`State`,`Type`=excluded.`Type`,`Data`=excluded.`Data`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindAssetById(ctx context.Context, db *sqlx.DB, keyId string) (*model.Asset, error) {
	obj := model.Asset{}

	err := db.QueryRowx("SELECT * FROM Assets WHERE Id=?", keyId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertSiteUser(ctx context.Context, db *sqlx.DB, obj model.SiteUser) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO SiteUsers VALUES (:UserId,:SiteId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(UserId,SiteId) DO UPDATE SET `UserId`=excluded.`UserId`,`SiteId`=excluded.`SiteId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func UpsertSiteUserTx(ctx context.Context, tx *sqlx.Tx, obj model.SiteUser) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO SiteUsers VALUES (:UserId,:SiteId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(UserId,SiteId) DO UPDATE SET `UserId`=excluded.`UserId`,`SiteId`=excluded.`SiteId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindSiteUserByUserIdAndSiteId(ctx context.Context, db *sqlx.DB, keyUserId string, keySiteId string) (*model.SiteUser, error) {
	obj := model.SiteUser{}

	err := db.QueryRowx("SELECT * FROM SiteUsers WHERE UserId=?AND SiteId=?", keyUserId, keySiteId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertItemGroup(ctx context.Context, db *sqlx.DB, obj model.ItemGroup) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO ItemGroups VALUES (:ItemId,:GroupId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(ItemId,GroupId) DO UPDATE SET `ItemId`=excluded.`ItemId`,`GroupId`=excluded.`GroupId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func UpsertItemGroupTx(ctx context.Context, tx *sqlx.Tx, obj model.ItemGroup) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO ItemGroups VALUES (:ItemId,:GroupId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(ItemId,GroupId) DO UPDATE SET `ItemId`=excluded.`ItemId`,`GroupId`=excluded.`GroupId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindItemGroupByItemIdAndGroupId(ctx context.Context, db *sqlx.DB, keyItemId string, keyGroupId string) (*model.ItemGroup, error) {
	obj := model.ItemGroup{}

	err := db.QueryRowx("SELECT * FROM ItemGroups WHERE ItemId=?AND GroupId=?", keyItemId, keyGroupId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}

func UpsertItemAsset(ctx context.Context, db *sqlx.DB, obj model.ItemAsset) error {
	stmt, err := db.PrepareNamedContext(ctx, "INSERT INTO ItemAssets VALUES (:ItemId,:AssetId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(ItemId,AssetId) DO UPDATE SET `ItemId`=excluded.`ItemId`,`AssetId`=excluded.`AssetId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func UpsertItemAssetTx(ctx context.Context, tx *sqlx.Tx, obj model.ItemAsset) error {
	stmt, err := tx.PrepareNamedContext(ctx, "INSERT INTO ItemAssets VALUES (:ItemId,:AssetId,:Order,:CreatedAt,:CreatedBy,:LastUpdatedAt,:LastUpdatedBy) ON CONFLICT(ItemId,AssetId) DO UPDATE SET `ItemId`=excluded.`ItemId`,`AssetId`=excluded.`AssetId`,`Order`=excluded.`Order`,`CreatedAt`=excluded.`CreatedAt`,`CreatedBy`=excluded.`CreatedBy`,`LastUpdatedAt`=excluded.`LastUpdatedAt`,`LastUpdatedBy`=excluded.`LastUpdatedBy`")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(obj)
	if err != nil {
		return err
	}

	return err
}

func FindItemAssetByItemIdAndAssetId(ctx context.Context, db *sqlx.DB, keyItemId string, keyAssetId string) (*model.ItemAsset, error) {
	obj := model.ItemAsset{}

	err := db.QueryRowx("SELECT * FROM ItemAssets WHERE ItemId=?AND AssetId=?", keyItemId, keyAssetId).StructScan(&obj)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &obj, nil
}
