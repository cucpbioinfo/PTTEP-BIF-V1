package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubFamilySQL = `
	CREATE TABLE IF NOT EXISTS public."sub_family"
	(
		"sub_family_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_family_name" varchar(500) UNIQUE NOT NULL,

		"family_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubFamilySQL = `DROP TABLE IF EXISTS public."sub_family";`

const addFKSubFamilyFamilySQL = `
	ALTER TABLE public."sub_family"
	ADD CONSTRAINT fk_sub_family_family
	FOREIGN KEY ("family_id") REFERENCES public."family" ("family_id")
`
const dropFKSubFamilyFamilySQL = `
	ALTER TABLE public."sub_family"
	DROP CONSTRAINT fk_sub_family_family
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_family...")
		var scripts = [2]string{
			createTableSubFamilySQL,
			addFKSubFamilyFamilySQL,
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

		fmt.Println("[Migration] Seeding table sub_family...")
		subFamilies, err := GetSubFamilyData()
		if err != nil {
			return err
		}

		for _, subFamily := range subFamilies {

			if subFamily.SubFamilyName == "NULL" {
				continue
			}

			insertSubFamilySQL := fmt.Sprintf(`
			INSERT INTO public."sub_family" ("sub_family_name", "family_id") VALUES ('%s', 
			(SELECT "family_id" FROM public."family" WHERE "family_name" = '%s' LIMIT 1))`,
				subFamily.SubFamilyName, subFamily.FamilyName)

			_, err := db.Exec(insertSubFamilySQL)
			if err != nil {
				fmt.Println(subFamily)
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_family...")
		var scripts = [2]string{
			dropFKSubFamilyFamilySQL,
			dropTableSubFamilySQL,
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
