package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableAssetSQL = `
	CREATE TABLE IF NOT EXISTS public."asset"
	(
		"asset_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"asset_name" varchar(500) UNIQUE NOT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableAssetSQL = `DROP TABLE IF EXISTS public."asset";`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table asset...")
		var scripts = [1]string{
			createTableAssetSQL,
		}
		var firstError error
		for _, script := range scripts {
			_, err := db.Exec(script)
			if err != nil {
				firstError = err
				break
			}
		}
		if firstError != nil {
			return firstError
		}

		fmt.Println("[Migration] Seeding table asset...")
		assets, err := GetAssetData()
		if err != nil {
			return err
		}

		for _, asset := range assets {
			insertAssetSQL := fmt.Sprintf(`
				INSERT INTO public."asset" ("asset_name") VALUES ('%s')`, asset.Asset)
			_, err := db.Exec(insertAssetSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table asset...")
		_, err := db.Exec(dropTableAssetSQL)
		return err
	})
}
