package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTablePhylumSQL = `
	CREATE TABLE IF NOT EXISTS public."phylum"
	(
		"phylum_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"phylum_name" varchar(500) UNIQUE NOT NULL,

		"kingdom_id" UUID,
		"super_phylum_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTablePhylumSQL = `DROP TABLE IF EXISTS public."phylum";`

const addFKPhylumKingdomSQL = `
	ALTER TABLE public."phylum"
	ADD CONSTRAINT fk_phylum_kingdom
	FOREIGN KEY ("kingdom_id") REFERENCES public."kingdom" ("kingdom_id")
`
const dropFKPhylumKingdomSQL = `
	ALTER TABLE public."phylum"
	DROP CONSTRAINT fk_phylum_kingdom
`

const addFKPhylumSuperPhylumSQL = `
	ALTER TABLE public."phylum"
	ADD CONSTRAINT fk_phylum_super_phylum
	FOREIGN KEY ("super_phylum_id") REFERENCES public."super_phylum" ("super_phylum_id")
`
const dropFKPhylumSuperPhylumSQL = `
	ALTER TABLE public."phylum"
	DROP CONSTRAINT fk_phylum_super_phylum
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table phylum...")
		var scripts = [3]string{
			createTablePhylumSQL,
			addFKPhylumKingdomSQL,
			addFKPhylumSuperPhylumSQL,
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

		fmt.Println("[Migration] Seeding table phylum...")
		phylums, err := GetPhylumData()
		if err != nil {
			return err
		}

		for _, phylum := range phylums {

			if phylum.PhylumName == "NULL" {
				continue
			}

			insertPhylumSQL := fmt.Sprintf(`
			INSERT INTO public."phylum" ("phylum_name", "kingdom_id", "super_phylum_id") VALUES ('%s', 
			(SELECT "kingdom_id" FROM public."kingdom" WHERE "kingdom_name" = '%s' LIMIT 1),
			(SELECT "super_phylum_id" FROM public."super_phylum" WHERE "super_phylum_name" = '%s' LIMIT 1))`,
				phylum.PhylumName, phylum.KingdomName, phylum.SuperPhylumName)

			_, err := db.Exec(insertPhylumSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table phylum...")
		var scripts = [3]string{
			dropFKPhylumSuperPhylumSQL,
			dropFKPhylumKingdomSQL,
			dropTablePhylumSQL,
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
