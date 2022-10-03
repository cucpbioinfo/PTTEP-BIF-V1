package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableDiversitySQL = `
	CREATE TABLE IF NOT EXISTS public."diversity"
	(
		"diversity_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"station_id" UUID NOT NULL,
		"surface" float NOT NULL,
		"euphotic_zone" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableDiversitySQL = `DROP TABLE IF EXISTS public."diversity";`

const addFKDiversityAssetSQL = `
	ALTER TABLE public."diversity"
	ADD CONSTRAINT fk_diversity_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKDiversityAssetSQL = `DROP CONSTRAINT IF EXISTS fk_diversity_asset;`

const addFKDiversityPlatformSQL = `
	ALTER TABLE public."diversity"
	ADD CONSTRAINT fk_diversity_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKDiversityPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_diversity_platform;`

const addFKDiversityStationSQL = `
	ALTER TABLE public."diversity"
	ADD CONSTRAINT fk_diversity_station
	FOREIGN KEY ("station_id") REFERENCES public."station" ("station_id")
	;`
const dropFKDiversityStationSQL = `DROP CONSTRAINT IF EXISTS fk_diversity_station;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table diversity...")
		var scripts = [4]string{
			createTableDiversitySQL,
			addFKDiversityAssetSQL,
			addFKDiversityPlatformSQL,
			addFKDiversityStationSQL,
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

		fmt.Println("[Migration] Seeding table diversity...")
		Diversities, err := GetDiversitiesData()
		if err != nil {
			fmt.Println("Cannot get even diversity data")
			return err
		}

		for _, diversity := range Diversities {
			insertdiversitySQL := fmt.Sprintf(`
			INSERT INTO public."diversity"("year",
			"asset_id", "platform_id","station_id","surface","euphotic_zone") VALUES('%s',
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			(SELECT station_id from public."station" WHERE station_name = '%s' LIMIT 1),
			'%s','%s');`,
				diversity.Year,
				diversity.Asset,
				diversity.Platform,
				diversity.Station,
				diversity.Surface,
				diversity.Euphotic_zone)
			_, err := db.Exec(insertdiversitySQL)
			if err != nil {
				fmt.Println(insertdiversitySQL)
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table diversity...")
		var scripts = [4]string{
			dropFKDiversityStationSQL,
			dropFKDiversityPlatformSQL,
			dropFKDiversityAssetSQL,
			dropTableDiversitySQL,
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
