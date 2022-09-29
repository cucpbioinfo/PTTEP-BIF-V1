package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableMethodSQL = `
	CREATE TABLE IF NOT EXISTS public."method"
	(
		"method_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"method_name" varchar(500) NOT NULL,
		"identification_id" UUID NOT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableMethodSQL = `DROP TABLE IF EXISTS public."method";`

const addFKMethodIdentificationSQL = `
	ALTER TABLE public."method"
	ADD CONSTRAINT fk_method_identification
	FOREIGN KEY ("identification_id") REFERENCES public."identification" ("identification_id")
	;`
const dropFKMethodIdentificationSQL = `DROP CONSTRAINT IF EXISTS fk_method_identification;`

const addMethodUniqueConstraint = `
ALTER TABLE public."method" ADD CONSTRAINT "method_unique" UNIQUE ("method_name")
`
const dropMethodUniqueConstraint = `DROP CONSTRAINT IF EXISTS method_unique`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table method...")
		var scripts = [3]string{
			createTableMethodSQL,
			addFKMethodIdentificationSQL,
			addMethodUniqueConstraint,
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

		fmt.Println("[Migration] Seeding table method...")
		Methods, err := GetMethodData()
		if err != nil {
			fmt.Println("Cannot get method data")
			return err
		}

		for _, method := range Methods {
			insertMethodSQL := fmt.Sprintf(`
			INSERT INTO public."method"("method_name",
			"identification_id") VALUES('%s',
			(SELECT identification_id from public."identification" WHERE identification_name = '%s' LIMIT 1))`,
				method.MethodName,
				method.IdentificationTechniqueName)
			_, err := db.Exec(insertMethodSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table method...")
		var scripts = [3]string{
			dropMethodUniqueConstraint,
			dropFKMethodIdentificationSQL,
			dropTableMethodSQL,
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
		return nil
	})
}
