package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableGenusSQL = `
	CREATE TABLE IF NOT EXISTS public."genus"
	(
		"genus_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"genus_name" varchar(500) UNIQUE NOT NULL,

		"family_id" UUID,
		"super_genus_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableGenusSQL = `DROP TABLE IF EXISTS public."genus";`

const addFKGenusFamilySQL = `
	ALTER TABLE public."genus"
	ADD CONSTRAINT fk_genus_family
	FOREIGN KEY ("family_id") REFERENCES public."family" ("family_id")
`
const dropFKGenusFamilySQL = `
	ALTER TABLE public."genus"
	DROP CONSTRAINT fk_genus_family
`

const addFKGenusSuperGenusSQL = `
	ALTER TABLE public."genus"
	ADD CONSTRAINT fk_genus_super_genus
	FOREIGN KEY ("super_genus_id") REFERENCES public."super_genus" ("super_genus_id")
`
const dropFKGenusSuperGenusSQL = `
	ALTER TABLE public."genus"
	DROP CONSTRAINT fk_genus_super_genus
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		var scripts = [3]string{
			createTableGenusSQL,
			addFKGenusFamilySQL,
			addFKGenusSuperGenusSQL,
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

		fmt.Println("[Migration] Seeding table genus...")
		genus, err := GetGenusData()
		if err != nil {
			return err
		}
		for _, g := range genus {

			if g.GenusName == "NULL" {
				continue
			}

			insertGenusSQL := fmt.Sprintf(`
			INSERT INTO public."genus" ("genus_name", "family_id", "super_genus_id") VALUES ('%s', 
			(SELECT "family_id" FROM public."family" WHERE "family_name" = '%s' LIMIT 1),
			(SELECT "super_genus_id" FROM public."super_genus" WHERE "super_genus_name" = '%s' LIMIT 1))`,
				g.GenusName, g.FamilyName, g.SuperGenusName)
			_, err := db.Exec(insertGenusSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table genus...")
		var scripts = [3]string{
			dropFKGenusSuperGenusSQL,
			dropFKGenusFamilySQL,
			dropTableGenusSQL,
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
