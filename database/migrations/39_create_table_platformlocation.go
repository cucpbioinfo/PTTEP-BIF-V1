package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

// //No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
const createTablePlatformLocationSQL = `
	CREATE TABLE IF NOT EXISTS public."platformlocation"
	(
		"platformlocation_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"latitude" float NOT NULL,
		"longitude" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTablePlatformLocationSQL = `DROP TABLE IF EXISTS public."platformlocation";`

const addFKPlatformLocationAssetSQL = `
	ALTER TABLE public."platformlocation"
	ADD CONSTRAINT fk_platformlocation_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKPlatformLocationAssetSQL = `DROP CONSTRAINT IF EXISTS fk_platformlocation_asset;`

const addFKPlatformLocationPlatformSQL = `
	ALTER TABLE public."platformlocation"
	ADD CONSTRAINT fk_platformlocation_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKPlatformLocationPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_platformlocation_platform;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table Platform location...")
		var scripts = [3]string{
			createTablePlatformLocationSQL,
			addFKPlatformLocationAssetSQL,
			addFKPlatformLocationPlatformSQL,
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

		fmt.Println("[Migration] Seeding table Platform location...")
		Locations, err := GetPlatformLocationData()
		if err != nil {
			fmt.Println("Cannot get even Platform location data")
			return err
		}

		for _, location := range Locations {
			insertlocationSQL := fmt.Sprintf(`
			INSERT INTO public."platformlocation"("asset_id", "platform_id","latitude","longitude") VALUES(
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			'%s','%s');`,
				location.Asset,
				location.Platform,
				location.Latitude,
				location.Longitude)
			_, err := db.Exec(insertlocationSQL)
			if err != nil {
				fmt.Println(insertlocationSQL)
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table Platform location...")
		var scripts = [3]string{
			dropFKPlatformLocationPlatformSQL,
			dropFKPlatformLocationAssetSQL,
			dropTablePlatformLocationSQL,
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
