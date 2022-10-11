package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableAssetSummarySQL = `
	CREATE TABLE IF NOT EXISTS public."assetsummary"
	(
		"summary_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"major_group_id" UUID NOT NULL,
		"identification_id" UUID NOT NULL,
		"asset_id" UUID NOT NULL,
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
const dropTableAssetSummarySQL = `DROP TABLE IF EXISTS public."assetsummary";`

const addFKAssetSummaryMajorGroupSQL = `
	ALTER TABLE public."major_group"
	ADD CONSTRAINT fk_assetsummary_major_group
	FOREIGN KEY ("major_group_id") REFERENCES public."major_group" ("major_group_id")
	;`
const dropFKAssetSummaryMajorGroupSQL = `DROP CONSTRAINT IF EXISTS fk_assetsummary_major_group;`

const addFKAssetSummaryIdentificationSQL = `
	ALTER TABLE public."identification"
	ADD CONSTRAINT fk_assetsummary_identification
	FOREIGN KEY ("identification_id") REFERENCES public."identification" ("identification_id")
	;`
const dropFKAssetSummaryIdentificationSQL = `DROP CONSTRAINT IF EXISTS fk_assetsummary_identification;`

const addFKAssetSummaryAssetSQL = `
	ALTER TABLE public."asset"
	ADD CONSTRAINT fk_assetsummary_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKAssetSummaryAssetSQL = `DROP CONSTRAINT IF EXISTS fk_assetsummary_asset;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table assetsummary...")
		var scripts = [4]string{
			createTableAssetSummarySQL,
			addFKAssetSummaryMajorGroupSQL,
			addFKAssetSummaryIdentificationSQL,
			addFKAssetSummaryAssetSQL,
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

		// fmt.Println("[Migration] Seeding table assetsummary...")
		// Summaries, err := GetSummariesData()
		// if err != nil {
		// 	fmt.Println("Cannot get even summary data")
		// 	return err
		// }

		// for _, summary := range Summaries {
		// 	insertSummarySQL := fmt.Sprintf(`
		// 	INSERT INTO public."summary"("year","major_group_id","identification_id",
		// 	"asset_id", "platform_id","station_id",
		// 	"surfaceshannon","surfacenumber","surfacemax","surfaceevenness",
		// 	"euphoticzoneshannon","euphoticzonenumber","euphoticzonemax","euphoticzoneevenness",
		// 	"averageshannon","averagenumber","averagemax","averageevenness"
		// 	) VALUES(
		// 	'%s',
		// 	(SELECT major_group_id from public."major_group" WHERE major_group_name = '%s' LIMIT 1),
		// 	(SELECT identification_id from public."identification" WHERE identification_name = '%s' LIMIT 1),
		// 	(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
		// 	(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
		// 	(SELECT station_id from public."station" WHERE station_name = '%s' LIMIT 1),
		// 	'%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');`,
		// 		summary.Year,
		// 		summary.Group, summary.Identification,
		// 		summary.Asset, summary.Platform, summary.Station,
		// 		summary.SurfaceShannon, summary.SurfaceNumber, summary.SurfaceMax, summary.SurfaceEvenness,
		// 		summary.Euphotic_zoneShannon, summary.Euphotic_zoneNumber, summary.Euphotic_zoneMax, summary.Euphotic_zoneEvenness,
		// 		summary.AverageShannon, summary.AverageNumber, summary.AverageMax, summary.AverageEvenness)
		// 	_, err := db.Exec(insertSummarySQL)
		// 	if err != nil {
		// 		fmt.Println(insertSummarySQL)
		// 		return err
		// 	}
		// }
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table assetsummary...")
		var scripts = [4]string{
			dropFKAssetSummaryAssetSQL,
			dropFKAssetSummaryIdentificationSQL,
			dropFKAssetSummaryMajorGroupSQL,
			dropTableAssetSummarySQL,
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
