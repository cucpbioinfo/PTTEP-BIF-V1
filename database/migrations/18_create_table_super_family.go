package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSuperFamilySQL = `
	CREATE TABLE IF NOT EXISTS public."super_family"
	(
		"super_family_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"super_family_name" varchar(500) UNIQUE NOT NULL,

		"order_id" UUID,
		"infra_order_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSuperFamilySQL = `DROP TABLE IF EXISTS public."super_family";`

const addFKSuperFamilyOrderSQL = `
	ALTER TABLE public."super_family"
	ADD CONSTRAINT fk_super_family_order
	FOREIGN KEY ("order_id") REFERENCES public."order" ("order_id")
`
const dropFKSuperFamilyOrderSQL = `
	ALTER TABLE public."super_family"
	DROP CONSTRAINT fk_super_family_order
`

const addFKSuperFamilyInfraOrderSQL = `
	ALTER TABLE public."super_family"
	ADD CONSTRAINT fk_super_family_infra_order
	FOREIGN KEY ("infra_order_id") REFERENCES public."infra_order" ("infra_order_id")
`
const dropFKSuperFamilyInfraOrderSQL = `
	ALTER TABLE public."super_family"
	DROP CONSTRAINT fk_super_family_infra_order
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table super_family...")
		var scripts = [3]string{
			createTableSuperFamilySQL,
			addFKSuperFamilyOrderSQL,
			addFKSuperFamilyInfraOrderSQL,
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

		fmt.Println("[Migration] Seeding table super_family...")
		superFamilies, err := GetSuperFamilyData()
		if err != nil {
			return err
		}

		for _, superFamily := range superFamilies {

			if superFamily.SuperFamilyName == "NULL" {
				continue
			}

			insertSuperFamilySQL := fmt.Sprintf(`
			INSERT INTO public."super_family" ("super_family_name", "order_id", "infra_order_id") VALUES ('%s', 
			(SELECT "order_id" FROM public."order" WHERE "order_name" = '%s' LIMIT 1),
			(SELECT "infra_order_id" FROM public."infra_order" WHERE "infra_order_name" = '%s' LIMIT 1))`,
				superFamily.SuperFamilyName, superFamily.OrderName, superFamily.InfraOrderName)

			_, err := db.Exec(insertSuperFamilySQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table super_family...")
		var scripts = [3]string{
			dropFKSuperFamilyInfraOrderSQL,
			dropFKSuperFamilyOrderSQL,
			dropTableSuperFamilySQL,
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
