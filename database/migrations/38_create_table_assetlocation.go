package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

// //No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
const createTableAssetLocationSQL = `
	CREATE TABLE IF NOT EXISTS public."assetlocation"
	(
		"assetlocation_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"asset_id" UUID NOT NULL,

		"latitude" float NOT NULL,
		"longitude" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableAssetLocationSQL = `DROP TABLE IF EXISTS public."assetlocation";`

const addFKAssetLocationAssetSQL = `
	ALTER TABLE public."assetlocation"
	ADD CONSTRAINT fk_assetlocation_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKAssetLocationAssetSQL = `DROP CONSTRAINT IF EXISTS fk_assetlocation_asset;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table Asset location...")
		var scripts = [2]string{
			createTableAssetLocationSQL,
			addFKAssetLocationAssetSQL,
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

		fmt.Println("[Migration] Seeding table Asset location...")
		Locations, err := GetAssetLocationData()
		if err != nil {
			fmt.Println("Cannot get even Asset location data")
			return err
		}

		for _, location := range Locations {
			insertlocationSQL := fmt.Sprintf(`
			INSERT INTO public."assetlocation"("asset_id","latitude","longitude") VALUES(
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			'%s','%s');`,
				location.Asset,
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
		fmt.Println("[Migration] Droping table Asset location...")
		var scripts = [2]string{
			dropFKAssetLocationAssetSQL,
			dropTableAssetLocationSQL,
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
