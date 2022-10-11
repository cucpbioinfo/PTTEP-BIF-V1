package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTablePlatformSummarySQL = `
	CREATE TABLE IF NOT EXISTS public."platformsummary"
	(
		"summary_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"major_group_id" UUID NOT NULL,
		"identification_id" UUID NOT NULL,
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"surfaceshannon" float NOT NULL,
		"surfacenumber" float NOT NULL,
		"surfacemax" float NOT NULL,
		"surfaceevenness" float NOT NULL,
		"euphoticzoneshannon" float NOT NULL,
		"euphoticzonenumber" float NOT NULL,
		"euphoticzonemax" float NOT NULL,
		"euphoticzoneevenness" float NOT NULL,
		"averageshannon" float NOT NULL,
		"averagenumber" float NOT NULL,
		"averagemax" float NOT NULL,
		"averageevenness" float NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTablePlatformSummarySQL = `DROP TABLE IF EXISTS public."platformsummary";`

const addFKPlatformSummaryMajorGroupSQL = `
	ALTER TABLE public."major_group"
	ADD CONSTRAINT fk_platformsummary_major_group
	FOREIGN KEY ("major_group_id") REFERENCES public."major_group" ("major_group_id")
	;`
const dropFKPlatformSummaryMajorGroupSQL = `DROP CONSTRAINT IF EXISTS fk_platformsummary_major_group;`

const addFKPlatformSummaryIdentificationSQL = `
	ALTER TABLE public."identification"
	ADD CONSTRAINT fk_platformsummary_identification
	FOREIGN KEY ("identification_id") REFERENCES public."identification" ("identification_id")
	;`
const dropFKPlatformSummaryIdentificationSQL = `DROP CONSTRAINT IF EXISTS fk_platformsummary_identification;`

const addFKPlatformSummaryAssetSQL = `
	ALTER TABLE public."asset"
	ADD CONSTRAINT fk_platformsummary_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKPlatformSummaryAssetSQL = `DROP CONSTRAINT IF EXISTS fk_platformsummary_asset;`

const addFKPlatformSummaryPlatformSQL = `
	ALTER TABLE public."platform"
	ADD CONSTRAINT fk_platformsummary_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKPlatformSummaryPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_platformsummary_platform;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table platformsummary...")
		var scripts = [5]string{
			createTablePlatformSummarySQL,
			addFKPlatformSummaryMajorGroupSQL,
			addFKPlatformSummaryIdentificationSQL,
			addFKPlatformSummaryAssetSQL,
			addFKPlatformSummaryPlatformSQL,
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

		fmt.Println("[Migration] Seeding table platformsummary...")
		Summaries, err := GetPlatformSummariesData()
		if err != nil {
			fmt.Println("Cannot get even platformsummary data")
			return err
		}

		for _, summary := range Summaries {
			insertSummarySQL := fmt.Sprintf(`
			INSERT INTO public."platformsummary"("year","major_group_id","identification_id",
			"asset_id", "platform_id",
			"surfaceshannon","surfacenumber","surfacemax","surfaceevenness",
			"euphoticzoneshannon","euphoticzonenumber","euphoticzonemax","euphoticzoneevenness",
			"averageshannon","averagenumber","averagemax","averageevenness"
			) VALUES(
			'%s',
			(SELECT major_group_id from public."major_group" WHERE major_group_name = '%s' LIMIT 1),
			(SELECT identification_id from public."identification" WHERE identification_name = '%s' LIMIT 1),
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');`,
				summary.Year,
				summary.Group, summary.Identification,
				summary.Asset, summary.Platform,
				summary.SurfaceShannon, summary.SurfaceNumber, summary.SurfaceMax, summary.SurfaceEvenness,
				summary.Euphotic_zoneShannon, summary.Euphotic_zoneNumber, summary.Euphotic_zoneMax, summary.Euphotic_zoneEvenness,
				summary.AverageShannon, summary.AverageNumber, summary.AverageMax, summary.AverageEvenness)
			_, err := db.Exec(insertSummarySQL)
			if err != nil {
				fmt.Println(insertSummarySQL)
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table platformsummary...")
		var scripts = [5]string{
			dropFKPlatformSummaryPlatformSQL,
			dropFKPlatformSummaryAssetSQL,
			dropFKPlatformSummaryIdentificationSQL,
			dropFKPlatformSummaryMajorGroupSQL,
			dropTablePlatformSummarySQL,
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
