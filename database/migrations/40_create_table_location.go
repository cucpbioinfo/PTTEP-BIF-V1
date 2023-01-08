package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

// //No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
const createTableLocationSQL = `
	CREATE TABLE IF NOT EXISTS public."location"
	(
		"location_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"station_id" UUID NOT NULL,
		"latitude" float NOT NULL,
		"longitude" float NOT NULL,
		"type" varchar(5) DEFAULT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableLocationSQL = `DROP TABLE IF EXISTS public."location";`

const addFKLocationAssetSQL = `
	ALTER TABLE public."location"
	ADD CONSTRAINT fk_location_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKLocationAssetSQL = `DROP CONSTRAINT IF EXISTS fk_location_asset;`

const addFKLocationPlatformSQL = `
	ALTER TABLE public."location"
	ADD CONSTRAINT fk_location_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKLocationPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_location_platform;`

const addFKLocationStationSQL = `
	ALTER TABLE public."location"
	ADD CONSTRAINT fk_location_station
	FOREIGN KEY ("station_id") REFERENCES public."station" ("station_id")
	;`
const dropFKLocationStationSQL = `DROP CONSTRAINT IF EXISTS fk_location_station;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table location...")
		var scripts = [4]string{
			createTableLocationSQL,
			addFKLocationAssetSQL,
			addFKLocationPlatformSQL,
			addFKLocationStationSQL,
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

		fmt.Println("[Migration] Seeding table location...")
		Locations, err := GetLocationData()
		if err != nil {
			fmt.Println("Cannot get even location data")
			return err
		}

		for _, location := range Locations {
			insertlocationSQL := fmt.Sprintf(`
			INSERT INTO public."location"("asset_id", "platform_id","station_id","latitude","longitude","type") VALUES(
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			(SELECT station_id from public."station" WHERE station_name = '%s' LIMIT 1),
			'%s','%s','%s');`,
				location.Asset,
				location.Platform,
				location.Station,
				location.Latitude,
				location.Longitude,
				location.Type)
			_, err := db.Exec(insertlocationSQL)
			if err != nil {
				fmt.Println(insertlocationSQL)
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table location...")
		var scripts = [4]string{
			dropFKLocationStationSQL,
			dropFKLocationPlatformSQL,
			dropFKLocationAssetSQL,
			dropTableLocationSQL,
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
