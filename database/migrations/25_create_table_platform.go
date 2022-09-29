package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTablePlatformSQL = `
	CREATE TABLE IF NOT EXISTS public."platform"
	(
		"platform_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"platform_name" varchar(500) UNIQUE NOT NULL,
		"asset_id" UUID NOT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTablePlatformSQL = `DROP TABLE IF EXISTS public."platform";`

const addFKPlatformAssetSQL = `
	ALTER TABLE public."platform"
	ADD CONSTRAINT fk_platform_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKPlatformAssetSQL = `DROP CONSTRAINT IF EXISTS fk_platform_asset;`

const addPlatformUniqueConstraint = `
ALTER TABLE public."platform" ADD CONSTRAINT "platform_unique" UNIQUE ("platform_name", "asset_id")
`
const dropPlatformUniqueConstraint = `DROP CONSTRAINT IF EXISTS platform_unique`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table platform...")
		var scripts = [3]string{
			createTablePlatformSQL,
			addFKPlatformAssetSQL,
			addPlatformUniqueConstraint,
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

		fmt.Println("[Migration] Seeding table platform...")
		Platforms, err := GetPlatformData()
		if err != nil {
			fmt.Println("Cannot get platform data")
			return err
		}

		for _, platform := range Platforms {
			insertPlatformSQL := fmt.Sprintf(`
			INSERT INTO public."platform"("platform_name","asset_id") 
			VALUES('%s',(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1))`,
				platform.Platform,
				platform.Asset)
			_, err := db.Exec(insertPlatformSQL)
			if err != nil {
				return err
			}
		}
		return nil

	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table platform...")
		var scripts = [3]string{
			dropPlatformUniqueConstraint,
			dropFKPlatformAssetSQL,
			dropTablePlatformSQL,
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
		return nil
	})
}
