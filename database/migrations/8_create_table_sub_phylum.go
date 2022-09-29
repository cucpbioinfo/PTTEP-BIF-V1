package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubPhylumSQL = `
	CREATE TABLE IF NOT EXISTS public."sub_phylum"
	(
		"sub_phylum_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_phylum_name" varchar(500) UNIQUE NOT NULL,

		"phylum_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubPhylumSQL = `DROP TABLE IF EXISTS public."sub_phylum";`

const addFKSubPhylumPhylumSQL = `
	ALTER TABLE public."sub_phylum"
	ADD CONSTRAINT fk_sub_phylum_phylum
	FOREIGN KEY ("phylum_id") REFERENCES public."phylum" ("phylum_id")
`
const dropFKSubPhylumPhylumSQL = `
	ALTER TABLE public."sub_phylum"
	DROP CONSTRAINT fk_sub_phylum_phylum
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_phylum...")
		var scripts = [2]string{
			createTableSubPhylumSQL,
			addFKSubPhylumPhylumSQL,
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

		fmt.Println("[Migration] Seeding table sub_phylum...")
		subPhylums, err := GetSubPhylumData()
		if err != nil {
			return err
		}

		for _, subPhylum := range subPhylums {

			if subPhylum.SubPhylumName == "NULL" {
				continue
			}

			insertSubPhylumSQL := fmt.Sprintf(`
			INSERT INTO public."sub_phylum" ("sub_phylum_name", "phylum_id") VALUES ('%s', 
			(SELECT "phylum_id" FROM public."phylum" WHERE "phylum_name" = '%s' LIMIT 1))`,
				subPhylum.SubPhylumName, subPhylum.PhylumName)

			_, err := db.Exec(insertSubPhylumSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_phylum...")
		var scripts = [2]string{
			dropFKSubPhylumPhylumSQL,
			dropTableSubPhylumSQL,
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
