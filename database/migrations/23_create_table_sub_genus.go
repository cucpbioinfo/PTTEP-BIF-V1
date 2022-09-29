package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubGenusSQL = `
	CREATE TABLE IF NOT EXISTS public."sub_genus"
	(
		"sub_genus_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_genus_name" varchar(500) UNIQUE NOT NULL,

		"genus_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubGenusSQL = `DROP TABLE IF EXISTS public."sub_genus";`

const addFKSubGenusGenusSQL = `
	ALTER TABLE public."sub_genus"
	ADD CONSTRAINT fk_sub_genus_genus
	FOREIGN KEY ("genus_id") REFERENCES public."genus" ("genus_id")
`
const dropFKSubGenusGenusSQL = `
	ALTER TABLE public."sub_genus"
	DROP CONSTRAINT fk_sub_genus_genus
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_genus...")
		var scripts = [2]string{
			createTableSubGenusSQL,
			addFKSubGenusGenusSQL,
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

		fmt.Println("[Migration] Seeding table sub_genus...")
		subGenus, err := GetSubGenusData()
		if err != nil {
			return err
		}

		for _, s := range subGenus {

			if s.SubGenusName == "NULL" {
				continue
			}

			insertSubGenusSQL := fmt.Sprintf(`
			INSERT INTO public."sub_genus" ("sub_genus_name", "genus_id") VALUES ('%s', 
			(SELECT "genus_id" FROM public."genus" WHERE "genus_name" = '%s' LIMIT 1))`,
				s.SubGenusName, s.GenusName)

			_, err := db.Exec(insertSubGenusSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_genus...")
		var scripts = [2]string{
			dropFKSubGenusGenusSQL,
			dropTableSubGenusSQL,
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
