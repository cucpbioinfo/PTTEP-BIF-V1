package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSuperOrderSQL = `
	CREATE TABLE IF NOT EXISTS public."super_order"
	(
		"super_order_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"super_order_name" varchar(500) UNIQUE NOT NULL,

		"class_id" UUID NOT NULL,
		"infra_class_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSuperOrderSQL = `DROP TABLE IF EXISTS public."super_order";`

const addFKSuperOrderClassSQL = `
	ALTER TABLE public."super_order"
	ADD CONSTRAINT fk_super_order_class
	FOREIGN KEY ("class_id") REFERENCES public."class" ("class_id")
`
const dropFKSuperOrderClassSQL = `
	ALTER TABLE public."super_order"
	DROP CONSTRAINT fk_super_order_class
`

const addFKSuperOrderInfraClassSQL = `
	ALTER TABLE public."super_order"
	ADD CONSTRAINT fk_super_order_infra_class
	FOREIGN KEY ("infra_class_id") REFERENCES public."infra_class" ("infra_class_id")
`
const dropFKSuperOrderInfraClassSQL = `
	ALTER TABLE public."super_order"
	DROP CONSTRAINT fk_super_order_infra_class
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table super_order...")
		var scripts = [3]string{
			createTableSuperOrderSQL,
			addFKSuperOrderInfraClassSQL,
			addFKSuperOrderClassSQL,
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

		fmt.Println("[Migration] Seeding table super_order...")
		superOrders, err := GetSuperOrderData()
		if err != nil {
			return err
		}

		for _, superOrder := range superOrders {

			if superOrder.SuperOrderName == "NULL" {
				continue
			}

			insertSuperOrderSQL := fmt.Sprintf(`
			INSERT INTO public."super_order" ("super_order_name", "class_id", "infra_class_id") VALUES ('%s', 
			(SELECT "class_id" FROM public."class" WHERE "class_name" = '%s' LIMIT 1),
			(SELECT "infra_class_id" FROM public."infra_class" WHERE "infra_class_name" = '%s' LIMIT 1))`,
				superOrder.SuperOrderName, superOrder.ClassName, superOrder.InfraClassName)

			_, err := db.Exec(insertSuperOrderSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table super_order...")
		var scripts = [3]string{
			dropFKSuperOrderInfraClassSQL,
			dropFKSuperOrderClassSQL,
			dropTableSuperOrderSQL,
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
