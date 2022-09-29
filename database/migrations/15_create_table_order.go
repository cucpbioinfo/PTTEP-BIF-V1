package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableOrderSQL = `
	CREATE TABLE IF NOT EXISTS public."order"
	(
		"order_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"order_name" varchar(500) UNIQUE NOT NULL,

		"class_id" UUID NOT NULL,
		"super_order_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableOrderSQL = `DROP TABLE IF EXISTS public."order";`

const addFKOrderClassSQL = `
	ALTER TABLE public."order"
	ADD CONSTRAINT fk_order_class
	FOREIGN KEY ("class_id") REFERENCES public."class" ("class_id")
`
const dropFKOrderClassSQL = `
	ALTER TABLE public."order"
	DROP CONSTRAINT fk_order_class
`

const addFKOrderSuperOrderSQL = `
	ALTER TABLE public."order"
	ADD CONSTRAINT fk_order_super_order
	FOREIGN KEY ("super_order_id") REFERENCES public."super_order" ("super_order_id")
`
const dropFKOrderSuperOrderSQL = `
	ALTER TABLE public."order"
	DROP CONSTRAINT fk_order_super_order
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table order...")
		var scripts = [3]string{
			createTableOrderSQL,
			addFKOrderClassSQL,
			addFKOrderSuperOrderSQL,
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

		fmt.Println("[Migration] Seeding table order...")
		orders, err := GetOrderData()
		if err != nil {
			return err
		}

		for _, order := range orders {

			if order.OrderName == "NULL" {
				continue
			}

			insertOrderSQL := fmt.Sprintf(`
			INSERT INTO public."order" ("order_name", "class_id", "super_order_id") VALUES ('%s', 
			(SELECT "class_id" FROM public."class" WHERE "class_name" = '%s' LIMIT 1),
			(SELECT "super_order_id" FROM public."super_order" WHERE "super_order_name" = '%s' LIMIT 1))`,
				order.OrderName, order.ClassName, order.SuperOrderName)
			_, err := db.Exec(insertOrderSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table order...")
		var scripts = [3]string{
			dropFKOrderSuperOrderSQL,
			dropFKOrderClassSQL,
			dropTableOrderSQL,
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
