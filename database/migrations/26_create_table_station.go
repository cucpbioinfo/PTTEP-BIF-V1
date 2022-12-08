package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableStationSQL = `
	CREATE TABLE IF NOT EXISTS public."station"
	(
		"station_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"station_name" varchar(500) UNIQUE NOT NULL,
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"latitude" float DEFAULT NULL,
		"longitude" float DEFAULT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableStationSQL = `DROP TABLE IF EXISTS public."station";`

const addFKStationAssetSQL = `
	ALTER TABLE public."station"
	ADD CONSTRAINT fk_station_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKStationAssetSQL = `DROP CONSTRAINT IF EXISTS fk_station_asset;`

const addFKStationPlatformSQL = `
	ALTER TABLE public."station"
	ADD CONSTRAINT fk_station_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKStationPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_station_platform;`

const addStationUniqueConstraint = `
ALTER TABLE public."station" ADD CONSTRAINT "station_unique" UNIQUE ("station_id")
`
const dropStationUniqueConstraint = `DROP CONSTRAINT IF EXISTS station_unique`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table station...")
		var scripts = [4]string{
			createTableStationSQL,
			addFKStationAssetSQL,
			addFKStationPlatformSQL,
			addStationUniqueConstraint,
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

		fmt.Println("[Migration] Seeding table station...")
		stations, err := GetStationData()
		if err != nil {
			fmt.Println("Cannot get station data")
			return err
		}

		for _, station := range stations {
			insertStationSQL := fmt.Sprintf(`
			INSERT INTO public."station"("station_name","asset_id","platform_id") VALUES('%s',
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1))`,
				station.Station,
				station.Asset,
				station.Platform)
			_, err := db.Exec(insertStationSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table station...")
		var scripts = [4]string{
			dropStationUniqueConstraint,
			dropFKStationPlatformSQL,
			dropFKStationAssetSQL,
			dropTableStationSQL,
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
