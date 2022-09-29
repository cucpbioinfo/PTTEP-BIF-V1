package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSuperGenusSQL = `
	CREATE TABLE IF NOT EXISTS public."super_genus"
	(
		"super_genus_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"super_genus_name" varchar(500) UNIQUE NOT NULL,

		"family_id" UUID,
		"sub_family_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSuperGenusSQL = `DROP TABLE IF EXISTS public."super_genus";`

const addFKSuperGenusFamilySQL = `
	ALTER TABLE public."super_genus"
	ADD CONSTRAINT fk_super_genus_family
	FOREIGN KEY ("family_id") REFERENCES public."family" ("family_id")
`
const dropFKSuperGenusFamilySQL = `
	ALTER TABLE public."super_genus"
	DROP CONSTRAINT fk_super_genus_family
`

const addFKSuperGenusSubFamilySQL = `
	ALTER TABLE public."super_genus"
	ADD CONSTRAINT fk_super_genus_sub_family
	FOREIGN KEY ("sub_family_id") REFERENCES public."sub_family" ("sub_family_id")
`
const dropFKSuperGenusSubFamilySQL = `
	ALTER TABLE public."super_genus"
	DROP CONSTRAINT fk_super_genus_sub_family
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table super_genus...")
		var scripts = [3]string{
			createTableSuperGenusSQL,
			addFKSuperGenusFamilySQL,
			addFKSuperGenusSubFamilySQL,
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

		fmt.Println("[Migration] Seeding table super_genus...")
		superGenusList, err := GetSuperGenusData()
		if err != nil {
			return err
		}

		for _, superGenus := range superGenusList {

			if superGenus.SuperGenusName == "NULL" {
				continue
			}

			insertSuperGenusSQL := fmt.Sprintf(`
			INSERT INTO public."super_genus" ("super_genus_name", "family_id", "sub_family_id") VALUES ('%s', 
			(SELECT "family_id" FROM public."family" WHERE "family_name" = '%s' LIMIT 1),
			(SELECT "sub_family_id" FROM public."sub_family" WHERE "sub_family_name" = '%s' LIMIT 1))`,
				superGenus.SuperGenusName, superGenus.FamilyName, superGenus.SubFamilyName)

			_, err := db.Exec(insertSuperGenusSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table super_genus...")
		var scripts = [3]string{
			dropFKSuperGenusSubFamilySQL,
			dropFKSuperGenusFamilySQL,
			dropTableSuperGenusSQL,
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
