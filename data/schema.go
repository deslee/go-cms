package data

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

func CreateTablesAndIndicesIfNotExist(db *sqlx.DB) {
	var (
		err error
		ctx = context.Background()
	)

	// begin the transaction
	tx, err := db.BeginTx(ctx, nil)
	die(err)
	defer tx.Rollback()
	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "Users" (
			"Id" TEXT NOT NULL CONSTRAINT "PK_Users" PRIMARY KEY,
			"Email" TEXT NOT NULL,
			"Password" TEXT NOT NULL,
			"Salt" TEXT NOT NULL,
			"Data" JSON NOT NULL,
			"CreatedBy" TEXT NOT NULL,
			"LastUpdatedBy" TEXT NOT NULL,
			"CreatedAt" TEXT NOT NULL,
			"LastUpdatedAt" TEXT NOT NULL
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "Sites" (
			"Id" TEXT NOT NULL CONSTRAINT "PK_Sites" PRIMARY KEY,
			"Name" TEXT NOT NULL,
			"Data" JSON NOT NULL,
			"CreatedBy" TEXT NOT NULL,
			"LastUpdatedBy" TEXT NOT NULL,
			"CreatedAt" TEXT NOT NULL,
			"LastUpdatedAt" TEXT NOT NULL
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "Items" (
			"Id" TEXT NOT NULL CONSTRAINT "PK_Items" PRIMARY KEY,
			"SiteId" TEXT NOT NULL,
			"Data" JSON NOT NULL,
			"Type" TEXT NOT NULL,
			"CreatedBy" TEXT NOT NULL,
			"LastUpdatedBy" TEXT NOT NULL,
			"CreatedAt" TEXT NOT NULL,
			"LastUpdatedAt" TEXT NOT NULL,
			CONSTRAINT "FK_Items_Sites_SiteId" FOREIGN KEY ("SiteId") REFERENCES "Sites" ("Id") ON DELETE CASCADE
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "Groups" (
			"Id" TEXT NOT NULL CONSTRAINT "PK_Groups" PRIMARY KEY,
			"SiteId" TEXT NOT NULL,
			"Data" JSON NOT NULL,
			"Name" TEXT NOT NULL,
			"CreatedBy" TEXT NOT NULL,
			"LastUpdatedBy" TEXT NOT NULL,
			"CreatedAt" TEXT NOT NULL,
			"LastUpdatedAt" TEXT NOT NULL,
			CONSTRAINT "FK_Groups_Sites_SiteId" FOREIGN KEY ("SiteId") REFERENCES "Sites" ("Id") ON DELETE CASCADE
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "Assets" (
			"Id" TEXT NOT NULL CONSTRAINT "PK_Assets" PRIMARY KEY,
			"SiteId" TEXT NOT NULL,
			"State" TEXT NOT NULL,
			"Type" TEXT NOT NULL,
			"Data" JSON NOT NULL,
			"CreatedBy" TEXT NOT NULL,
			"LastUpdatedBy" TEXT NOT NULL,
			"CreatedAt" TEXT NOT NULL,
			"LastUpdatedAt" TEXT NOT NULL,
			CONSTRAINT "FK_Assets_Sites_SiteId" FOREIGN KEY ("SiteId") REFERENCES "Sites" ("Id") ON DELETE CASCADE
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "SiteUsers" (
			"UserId" TEXT NOT NULL,
			"SiteId" TEXT NOT NULL,
			"Order" INTEGER NOT NULL,
			"CreatedBy" TEXT NULL,
			"LastUpdatedBy" TEXT NULL,
			"CreatedAt" TEXT NOT NULL,
			"LastUpdatedAt" TEXT NOT NULL,
			CONSTRAINT "PK_SiteUsers" PRIMARY KEY ("UserId", "SiteId"),
			CONSTRAINT "FK_SiteUsers_Sites_SiteId" FOREIGN KEY ("SiteId") REFERENCES "Sites" ("Id") ON DELETE CASCADE,
			CONSTRAINT "FK_SiteUsers_Users_UserId" FOREIGN KEY ("UserId") REFERENCES "Users" ("Id") ON DELETE CASCADE
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "ItemGroups" ( 
			"ItemId" TEXT NOT NULL, 
			"GroupId" TEXT NOT NULL, 
			"Order" INTEGER NOT NULL, 
			"CreatedBy" TEXT NOT NULL, 
			"LastUpdatedBy" TEXT NOT NULL, 
			"CreatedAt" TEXT NOT NULL, 
			"LastUpdatedAt" TEXT NOT NULL, 
			CONSTRAINT "PK_ItemGroups" PRIMARY KEY ("ItemId", "GroupId"), 
			CONSTRAINT "FK_ItemGroups_Groups_GroupId" FOREIGN KEY ("GroupId") REFERENCES "Groups" ("Id") ON DELETE CASCADE, 
			CONSTRAINT "FK_ItemGroups_Items_ItemId" FOREIGN KEY ("ItemId") REFERENCES "Items" ("Id") ON DELETE CASCADE 
		)
	`)

	executeSql(tx, `
		CREATE TABLE IF NOT EXISTS "ItemAssets" ( 
			"ItemId" TEXT NOT NULL, 
			"AssetId" TEXT NOT NULL, 
			"Order" INTEGER NOT NULL, 
			"CreatedBy" TEXT NOT NULL, 
			"LastUpdatedBy" TEXT NOT NULL, 
			"CreatedAt" TEXT NOT NULL, 
			"LastUpdatedAt" TEXT NOT NULL, 
			CONSTRAINT "PK_ItemAssets" PRIMARY KEY ("ItemId", "AssetId"), 
			CONSTRAINT "FK_ItemAssets_Assets_AssetId" FOREIGN KEY ("AssetId") REFERENCES "Assets" ("Id") ON DELETE CASCADE, 
			CONSTRAINT "FK_ItemAssets_Items_ItemId" FOREIGN KEY ("ItemId") REFERENCES "Items" ("Id") ON DELETE CASCADE 
		)
	`)

	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_Assets_SiteId" ON "Assets" ("SiteId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_Groups_SiteId" ON "Groups" ("SiteId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_Items_SiteId" ON "Items" ("SiteId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_ItemAssets_AssetId" ON "ItemAssets" ("AssetId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_ItemAssets_ItemId" ON "ItemAssets" ("ItemId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_ItemGroups_GroupId" ON "ItemGroups" ("GroupId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_ItemGroups_ItemId" ON "ItemGroups" ("ItemId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_SiteUsers_SiteId" ON "SiteUsers" ("SiteId")`)
	executeSql(tx, `CREATE INDEX IF NOT EXISTS "IX_SiteUsers_UserId" ON "SiteUsers" ("UserId")`)
	executeSql(tx, `CREATE UNIQUE INDEX IF NOT EXISTS "IX_Users_Email" ON "Users" ("Email")`)

	// commit the transaction
	err = tx.Commit()
	die(err)
}

func executeSql(tx *sql.Tx, sql string, args ...interface{}) {
	// prepare a statement
	stmt, err := tx.Prepare(sql)
	die(err)
	defer stmt.Close()

	// execute statement
	_, err = stmt.Exec(args...)
	die(err)
}
