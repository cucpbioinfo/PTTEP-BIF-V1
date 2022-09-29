package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableInfraOrderSQL = `
	CREATE TABLE IF NOT EXISTS public."infra_order"
	(
		"infra_order_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"infra_order_name" varchar(500) UNIQUE NOT NULL,

		"order_id" UUID,
		"sub_order_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableInfraOrderSQL = `DROP TABLE IF EXISTS public."infra_order";`

const addFKInfraOrderOrderSQL = `
	ALTER TABLE public."infra_order"
	ADD CONSTRAINT fk_infra_order_order
	FOREIGN KEY ("order_id") REFERENCES public."order" ("order_id")
`
const dropFKInfraOrderOrderSQL = `
	ALTER TABLE public."infra_order"
	DROP CONSTRAINT fk_infra_order_order
`

const addFKInfraOrderSubOrderSQL = `
	ALTER TABLE public."infra_order"
	ADD CONSTRAINT fk_infra_order_sub_order
	FOREIGN KEY ("sub_order_id") REFERENCES public."sub_order" ("sub_order_id")
`
const dropFKInfraOrderSubOrderSQL = `
	ALTER TABLE public."infra_order"
	DROP CONSTRAINT fk_infra_order_sub_order
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table infra_order...")
		var scripts = [3]string{
			createTableInfraOrderSQL,
			addFKInfraOrderOrderSQL,
			addFKInfraOrderSubOrderSQL,
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

		fmt.Println("[Migration] Seeding table infra_order...")
		infraOrders, err := GetInfraOrderData()
		if err != nil {
			return err
		}

		for _, infraOrder := range infraOrders {

			if infraOrder.InfraOrderName == "NULL" {
				continue
			}

			insertInfraOrderSQL := fmt.Sprintf(`
			INSERT INTO public."infra_order" ("infra_order_name", "order_id", "sub_order_id") VALUES ('%s', 
			(SELECT "order_id" FROM public."order" WHERE "order_name" = '%s' LIMIT 1),
			(SELECT "sub_order_id" FROM public."sub_order" WHERE "sub_order_name" = '%s' LIMIT 1))`,
				infraOrder.InfraOrderName, infraOrder.OrderName, infraOrder.SubOrderName)

			_, err := db.Exec(insertInfraOrderSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table infra_order...")
		var scripts = [3]string{
			dropFKInfraOrderSubOrderSQL,
			dropFKInfraOrderOrderSQL,
			dropTableInfraOrderSQL,
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
