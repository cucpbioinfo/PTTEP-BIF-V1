package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubKingdomSQL = `
	CREATE TABLE IF NOT EXISTS public."sub_kingdom"
	(
		"sub_kingdom_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_kingdom_name" varchar(500) UNIQUE NOT NULL,
		"kingdom_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubKingdomSQL = `DROP TABLE IF EXISTS public."sub_kingdom";`

const addFKSubKingdomKingdomSQL = `
	ALTER TABLE public."sub_kingdom"
	ADD CONSTRAINT fk_sub_kingdom_kingdom
	FOREIGN KEY ("kingdom_id") REFERENCES public."kingdom" ("kingdom_id")
`
const dropFKSubKingdomKingdomSQL = `
	ALTER TABLE public."sub_kingdom" 
	DROP CONSTRAINT fk_sub_kingdom_kingdom
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_kingdom...")
		var scripts = [2]string{
			createTableSubKingdomSQL,
			addFKSubKingdomKingdomSQL,
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

		fmt.Println("[Migration] Seeding table sub_kingdom...")
		subKingdoms, err := GetSubKingdomData()
		if err != nil {
			return err
		}

		for _, subKingdom := range subKingdoms {

			if subKingdom.SubKingdomName == "NULL" {
				continue
			}

			insertSubKingdomSQL := fmt.Sprintf(`
				INSERT INTO public."sub_kingdom" ("sub_kingdom_name", "kingdom_id") VALUES ('%s', 
				(SELECT "kingdom_id" FROM public."kingdom" WHERE "kingdom_name" = '%s' LIMIT 1))`,
				subKingdom.SubKingdomName, subKingdom.KingdomName)
			_, err := db.Exec(insertSubKingdomSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_kingdom...")
		var scripts = [2]string{
			dropFKSubKingdomKingdomSQL,
			dropTableSubKingdomSQL,
		}
		var firstError error
		for _, script := range scripts {
			_, err := db.Exec(script)
			if err != nil {
				firstError = err
				break
			}
		}
		return firstError
	})
}
