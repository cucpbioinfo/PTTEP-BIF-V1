package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

// //No.,Species,Year,Asset,Platform,Station,Surface,Euphotic_zone
const createTableAssetDensitySQL = `
	CREATE TABLE IF NOT EXISTS public."assetdensity"
	(
		"density_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"species_id" UUID,
		"species_name" varchar(500),
		"asset_id" UUID NOT NULL,
		"surface" float NOT NULL,
		"euphotic_zone" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableAssetDensitySQL = `DROP TABLE IF EXISTS public."assetdensity";`

const addFKAssetDensityAssetSQL = `
	ALTER TABLE public."assetdensity"
	ADD CONSTRAINT fk_assetdensity_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKAssetDensityAssetSQL = `DROP CONSTRAINT IF EXISTS fk_assetdensity_asset;`

const addFKAssetDensitySpeciesSQL = `
	ALTER TABLE public."assetdensity"
	ADD CONSTRAINT fk_assetdensity_species
	FOREIGN KEY ("species_id") REFERENCES public."species" ("species_id")
	;`
const dropFKAssetDensitySpeciesSQL = `DROP CONSTRAINT IF EXISTS fk_assetdensity_species;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table assetdensity...")
		var scripts = [3]string{
			createTableAssetDensitySQL,
			addFKAssetDensityAssetSQL,
			addFKAssetDensitySpeciesSQL,
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

		fmt.Println("[Migration] Seeding table assetdensity...")
		Densities, err := GetAssetDensitiesData()
		if err != nil {
			fmt.Println("Cannot get even assetdensity data")
			return err
		}

		for _, density := range Densities {
			insertassetdensitySQL := fmt.Sprintf(`
			INSERT INTO public."assetdensity"("year",
			"asset_id","species_id","species_name","surface","euphotic_zone") VALUES('%s',
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT species_id from public."species" WHERE species_name = '%s' LIMIT 1),
			'%s','%s','%s');`,
				density.Year,
				density.Asset,
				density.Species,
				density.Species,
				density.Surface,
				density.Euphotic_zone)
			_, err := db.Exec(insertassetdensitySQL)
			if err != nil {
				fmt.Println(insertassetdensitySQL)
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table assetdensity...")
		var scripts = [3]string{
			dropFKAssetDensitySpeciesSQL,
			dropFKAssetDensityAssetSQL,
			dropTableAssetDensitySQL,
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
