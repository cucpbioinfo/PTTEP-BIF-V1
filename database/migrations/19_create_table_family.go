package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableFamilySQL = `
	CREATE TABLE IF NOT EXISTS public."family"
	(
		"family_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"family_name" varchar(500) UNIQUE NOT NULL,

		"order_id" UUID,
		"super_family_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableFamilySQL = `DROP TABLE IF EXISTS public."family";`

const addFKFamilyOrderSQL = `
	ALTER TABLE public."family"
	ADD CONSTRAINT fk_family_order
	FOREIGN KEY ("order_id") REFERENCES public."order" ("order_id")
`
const dropFKFamilyOrderSQL = `
	ALTER TABLE public."family"
	DROP CONSTRAINT fk_family_order
`

const addFKFamilySuperFamilySQL = `
	ALTER TABLE public."family"
	ADD CONSTRAINT fk_family_super_family
	FOREIGN KEY ("super_family_id") REFERENCES public."super_family" ("super_family_id")
`
const dropFKFamilySuperFamilySQL = `
	ALTER TABLE public."family"
	DROP CONSTRAINT fk_family_super_family
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table family...")
		var scripts = [3]string{
			createTableFamilySQL,
			addFKFamilyOrderSQL,
			addFKFamilySuperFamilySQL,
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

		fmt.Println("[Migration] Seeding table family...")
		families, err := GetFamilyData()
		if err != nil {
			return err
		}
		for _, family := range families {

			if family.FamilyName == "NULL" {
				continue
			}

			insertFamilySQL := fmt.Sprintf(`
			INSERT INTO public."family" ("family_name", "order_id", "super_family_id") VALUES ('%s', 
			(SELECT "order_id" FROM public."order" WHERE "order_name" = '%s' LIMIT 1),
			(SELECT "super_family_id" FROM public."super_family" WHERE "super_family_name" = '%s' LIMIT 1))`,
				family.FamilyName, family.OrderName, family.SuperFamilyName)
			_, err := db.Exec(insertFamilySQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table family...")
		var scripts = [3]string{
			dropFKFamilySuperFamilySQL,
			dropFKFamilyOrderSQL,
			dropTableFamilySQL,
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
