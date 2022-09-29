package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableInfraKingdomSQL = `
	CREATE TABLE IF NOT EXISTS public."infra_kingdom"
	(
		"infra_kingdom_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"infra_kingdom_name" varchar(500) UNIQUE NOT NULL,

		"kingdom_id" UUID,
		"sub_kingdom_id" UUID,
		
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableInfraKingdomSQL = `DROP TABLE IF EXISTS public."infra_kingdom";`

const addFKInfraKingdomKingdomSQL = `
	ALTER TABLE public."infra_kingdom"
	ADD CONSTRAINT fk_infra_kingdom_kingdom
	FOREIGN KEY ("kingdom_id") REFERENCES public."kingdom" ("kingdom_id")
`
const dropFKInfraKingdomKingdomSQL = `
	ALTER TABLE public."infra_kingdom" 
	DROP CONSTRAINT fk_infra_kingdom_kingdom
`

const addFKInfraKingdomSubKingdomSQL = `
	ALTER TABLE public."infra_kingdom"
	ADD CONSTRAINT fk_infra_kingdom_sub_kingdom
	FOREIGN KEY ("sub_kingdom_id") REFERENCES public."sub_kingdom" ("sub_kingdom_id")
`
const dropFKInfraKingdomSubKingdomSQL = `
	ALTER TABLE public."infra_kingdom" 
	DROP CONSTRAINT fk_infra_kingdom_sub_kingdom
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table infra_kingdom...")
		var scripts = [3]string{
			createTableInfraKingdomSQL,
			addFKInfraKingdomKingdomSQL,
			addFKInfraKingdomSubKingdomSQL,
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

		fmt.Println("[Migration] Seeding table infra_kingdom...")
		infraKingdoms, err := GetInfraKingdomData()
		if err != nil {
			return err
		}

		for _, infraKingdom := range infraKingdoms {

			if infraKingdom.InfraKingdomName == "NULL" {
				continue
			}

			insertInfraKingdomSQL := fmt.Sprintf(`
				INSERT INTO public."infra_kingdom" ("infra_kingdom_name", "kingdom_id", "sub_kingdom_id") VALUES ('%s', 
				(SELECT "kingdom_id" FROM public."kingdom" WHERE "kingdom_name" = '%s' LIMIT 1),
				(SELECT "sub_kingdom_id" FROM public."sub_kingdom" WHERE "sub_kingdom_name" = '%s' LIMIT 1))`,
				infraKingdom.InfraKingdomName, infraKingdom.KingdomName, infraKingdom.SubKingdomName)

			_, err := db.Exec(insertInfraKingdomSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table infra_kingdom...")
		var scripts = [3]string{
			dropFKInfraKingdomSubKingdomSQL,
			dropFKInfraKingdomKingdomSQL,
			dropTableInfraKingdomSQL,
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
