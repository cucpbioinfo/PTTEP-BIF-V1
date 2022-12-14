package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

////No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
const createTableDensitySQL = `
	CREATE TABLE IF NOT EXISTS public."density"
	(
		"density_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"species_id" UUID,
		"species_name" varchar(500),
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"station_id" UUID NOT NULL,
		"surface" float NOT NULL,
		"euphotic_zone" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableDensitySQL = `DROP TABLE IF EXISTS public."density";`

const addFKDensityAssetSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKDensityAssetSQL = `DROP CONSTRAINT IF EXISTS fk_density_asset;`

const addFKDensityPlatformSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKDensityPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_density_platform;`

const addFKDensityStationSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_station
	FOREIGN KEY ("station_id") REFERENCES public."station" ("station_id")
	;`
const dropFKDensityStationSQL = `DROP CONSTRAINT IF EXISTS fk_density_station;`

const addFKDensitySpeciesSQL = `
	ALTER TABLE public."density"
	ADD CONSTRAINT fk_density_species
	FOREIGN KEY ("species_id") REFERENCES public."species" ("species_id")
	;`
const dropFKDensitySpeciesSQL = `DROP CONSTRAINT IF EXISTS fk_density_species;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table density...")
		var scripts = [5]string{
			createTableDensitySQL,
			addFKDensityAssetSQL,
			addFKDensityPlatformSQL,
			addFKDensityStationSQL,
			addFKDensitySpeciesSQL,
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

		fmt.Println("[Migration] Seeding table density...")
		Densities, err := GetDensitiesData()
		if err != nil {
			fmt.Println("Cannot get even density data")
			return err
		}

		////No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
		// "density_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		// "year" varchar(500) NOT NULL,
		// "species_id" UUID,
		// "species_name" varchar(500),
		// "asset_id" UUID NOT NULL,
		// "platform_id" UUID NOT NULL,
		// "station_id" UUID NOT NULL,
		// "surface" float NOT NULL,
		// "euphotic_zone" float NOT NULL,

		// "created_at" TIMESTAMPTZ DEFAULT NOW(),
		// "updated_at" TIMESTAMPTZ DEFAULT NOW(),
		// "deleted_at" TIMESTAMPTZ DEFAULT NULL
		for _, density := range Densities {
			insertdensitySQL := fmt.Sprintf(`
			INSERT INTO public."density"("year",
			"asset_id", "platform_id","station_id","species_id","species_name","surface","euphotic_zone") VALUES('%s',
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			(SELECT station_id from public."station" WHERE station_name = '%s' LIMIT 1),
			(SELECT species_id from public."species" WHERE species_name = '%s' LIMIT 1),
			'%s','%s','%s');`,
				density.Year,
				density.Asset,
				density.Platform,
				density.Station,
				density.Species,
				density.Species,
				density.Surface,
				density.Euphotic_zone)
			_, err := db.Exec(insertdensitySQL)
			if err != nil {
				fmt.Println(insertdensitySQL)
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table density...")
		var scripts = [5]string{
			dropFKDensitySpeciesSQL,
			dropFKDensityStationSQL,
			dropFKDensityPlatformSQL,
			dropFKDensityAssetSQL,
			dropTableDensitySQL,
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
