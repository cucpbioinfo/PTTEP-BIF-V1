package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubOrderSQL = `
	CREATE TABLE IF NOT EXISTS public."sub_order"
	(
		"sub_order_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_order_name" varchar(500) UNIQUE NOT NULL,

		"order_id" UUID NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubOrderSQL = `DROP TABLE IF EXISTS public."sub_order";`

const addFKSubOrderOrderSQL = `
	ALTER TABLE public."sub_order"
	ADD CONSTRAINT fk_sub_order_order
	FOREIGN KEY ("order_id") REFERENCES public."order" ("order_id")
`
const dropFKSubOrderOrderSQL = `
	ALTER TABLE public."sub_order"
	DROP CONSTRAINT fk_sub_order_order
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_order...")
		var scripts = [2]string{
			createTableSubOrderSQL,
			addFKSubOrderOrderSQL,
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

		fmt.Println("[Migration] Seeding table sub_class...")
		subOrders, err := GetSubOrderData()
		if err != nil {
			return err
		}

		for _, subOrder := range subOrders {

			if subOrder.SubOrderName == "NULL" {
				continue
			}

			insertSubOrderSQL := fmt.Sprintf(`
			INSERT INTO public."sub_order" ("sub_order_name", "order_id") VALUES ('%s', 
			(SELECT "order_id" FROM public."order" WHERE "order_name" = '%s' LIMIT 1))`,
				subOrder.SubOrderName, subOrder.OrderName)

			_, err := db.Exec(insertSubOrderSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_order...")
		var scripts = [2]string{
			dropFKSubOrderOrderSQL,
			dropTableSubOrderSQL,
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
