package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

// //No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
const createTablePlatformDensitySQL = `
	CREATE TABLE IF NOT EXISTS public."platformdensity"
	(
		"density_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"species_id" UUID,
		"species_name" varchar(500),
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"surface" float NOT NULL,
		"euphotic_zone" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTablePlatformDensitySQL = `DROP TABLE IF EXISTS public."platformdensity";`

const addFKPlatformDensityAssetSQL = `
	ALTER TABLE public."platformdensity"
	ADD CONSTRAINT fk_platformdensity_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKPlatformDensityAssetSQL = `DROP CONSTRAINT IF EXISTS fk_platformdensity_asset;`

const addFKPlatformDensityPlatformSQL = `
	ALTER TABLE public."platformdensity"
	ADD CONSTRAINT fk_platformdensity_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKPlatformDensityPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_platformdensity_platform;`

const addFKPlatformDensitySpeciesSQL = `
	ALTER TABLE public."platformdensity"
	ADD CONSTRAINT fk_platformdensity_species
	FOREIGN KEY ("species_id") REFERENCES public."species" ("species_id")
	;`
const dropFKPlatformDensitySpeciesSQL = `DROP CONSTRAINT IF EXISTS fk_platformdensity_species;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table platformdensity...")
		var scripts = [4]string{
			createTablePlatformDensitySQL,
			addFKPlatformDensityAssetSQL,
			addFKPlatformDensityPlatformSQL,
			addFKPlatformDensitySpeciesSQL,
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

		fmt.Println("[Migration] Seeding table platformdensity...")
		Densities, err := GetPlatformDensitiesData()
		if err != nil {
			fmt.Println("Cannot get even platformdensity data")
			return err
		}

		for _, density := range Densities {
			insertplatformdensitySQL := fmt.Sprintf(`
			INSERT INTO public."platformdensity"("year",
			"asset_id", "platform_id","species_id","species_name","surface","euphotic_zone") VALUES('%s',
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			(SELECT species_id from public."species" WHERE species_name = '%s' LIMIT 1),
			'%s','%s','%s');`,
				density.Year,
				density.Asset,
				density.Platform,
				density.Species,
				density.Species,
				density.Surface,
				density.Euphotic_zone)
			_, err := db.Exec(insertplatformdensitySQL)
			if err != nil {
				fmt.Println(insertplatformdensitySQL)
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table assetdensity...")
		var scripts = [4]string{
			dropFKPlatformDensitySpeciesSQL,
			dropFKPlatformDensityPlatformSQL,
			dropFKPlatformDensityAssetSQL,
			dropTablePlatformDensitySQL,
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
