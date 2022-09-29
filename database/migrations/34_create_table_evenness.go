package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableEvennessSQL = `
	CREATE TABLE IF NOT EXISTS public."evenness"
	(
		"evenness_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"year" varchar(500) NOT NULL,
		"asset_id" UUID NOT NULL,
		"platform_id" UUID NOT NULL,
		"station_id" UUID NOT NULL,
		"surface" float DEFAULT NULL,
		"euphotic_zone" float DEFAULT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableEvennessSQL = `DROP TABLE IF EXISTS public."evenness";`

const addFKEvennessAssetSQL = `
	ALTER TABLE public."evenness"
	ADD CONSTRAINT fk_evenness_asset
	FOREIGN KEY ("asset_id") REFERENCES public."asset" ("asset_id")
	;`
const dropFKEvennessAssetSQL = `DROP CONSTRAINT IF EXISTS fk_evenness_asset;`

const addFKEvennessPlatformSQL = `
	ALTER TABLE public."evenness"
	ADD CONSTRAINT fk_evenness_platform
	FOREIGN KEY ("platform_id") REFERENCES public."platform" ("platform_id")
	;`
const dropFKEvennessPlatformSQL = `DROP CONSTRAINT IF EXISTS fk_evenness_platform;`

const addFKEvennessStationSQL = `
	ALTER TABLE public."evenness"
	ADD CONSTRAINT fk_evenness_station
	FOREIGN KEY ("station_id") REFERENCES public."station" ("station_id")
	;`
const dropFKEvennessStationSQL = `DROP CONSTRAINT IF EXISTS fk_evenness_station;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table evenness...")
		var scripts = [4]string{
			createTableEvennessSQL,
			addFKEvennessAssetSQL,
			addFKEvennessPlatformSQL,
			addFKEvennessStationSQL,
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

		fmt.Println("[Migration] Seeding table eveness...")
		Evennesses, err := GetEvennessesData()
		if err != nil {
			fmt.Println("Cannot get even evenness data")
			return err
		}

		for _, evenness := range Evennesses {
			insertevennessSQL := fmt.Sprintf(`
			INSERT INTO public."evenness"("year",
			"asset_id", "platform_id","station_id","surface","euphotic_zone") VALUES('%s',
			(SELECT asset_id from public."asset" WHERE asset_name = '%s' LIMIT 1),
			(SELECT platform_id from public."platform" WHERE platform_name = '%s' LIMIT 1),
			(SELECT station_id from public."station" WHERE station_name = '%s' LIMIT 1),
			'%s','%s');`,
				evenness.Year,
				evenness.Asset,
				evenness.Platform,
				evenness.Station,
				evenness.Surface,
				evenness.Euphotic_zone)
			_, err := db.Exec(insertevennessSQL)
			if err != nil {
				//fmt.Println(insertevennessSQL)
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table evenness...")
		var scripts = [4]string{
			dropFKEvennessStationSQL,
			dropFKEvennessPlatformSQL,
			dropFKEvennessAssetSQL,
			dropTableEvennessSQL,
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
