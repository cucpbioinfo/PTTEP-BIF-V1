package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableInfraClassSQL = `
	CREATE TABLE IF NOT EXISTS public."infra_class"
	(
		"infra_class_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"infra_class_name" varchar(500) UNIQUE NOT NULL,

		"class_id" UUID,
		"sub_class_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableInfraClassSQL = `DROP TABLE IF EXISTS public."infra_class";`

const addFKInfraClassClassSQL = `
	ALTER TABLE public."infra_class"
	ADD CONSTRAINT fk_infra_class_class
	FOREIGN KEY ("class_id") REFERENCES public."class" ("class_id")
`
const dropFKInfraClassClassSQL = `
	ALTER TABLE public."infra_class"
	DROP CONSTRAINT fk_infra_class_class
`

const addFKInfraClassSubClassSQL = `
	ALTER TABLE public."infra_class"
	ADD CONSTRAINT fk_infra_class_sub_class
	FOREIGN KEY ("sub_class_id") REFERENCES public."sub_class" ("sub_class_id")
`
const dropFKInfraClassSubClassSQL = `
	ALTER TABLE public."infra_class"
	DROP CONSTRAINT fk_infra_class_sub_class
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table infra_class...")
		var scripts = [3]string{
			createTableInfraClassSQL,
			addFKInfraClassClassSQL,
			addFKInfraClassSubClassSQL,
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

		fmt.Println("[Migration] Seeding table infra_class...")
		infraClasses, err := GetInfraClassData()
		if err != nil {
			return err
		}

		for _, infraClass := range infraClasses {

			if infraClass.InfraClassName == "NULL" {
				continue
			}

			insertInfraClassSQL := fmt.Sprintf(`
			INSERT INTO public."infra_class" ("infra_class_name", "class_id", "sub_class_id") VALUES ('%s', 
			(SELECT "class_id" FROM public."class" WHERE "class_name" = '%s' LIMIT 1),
			(SELECT "sub_class_id" FROM public."sub_class" WHERE "sub_class_name" = '%s' LIMIT 1))`,
				infraClass.InfraClassName, infraClass.ClassName, infraClass.SubClassName)

			_, err := db.Exec(insertInfraClassSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table infra_class...")
		var scripts = [3]string{
			dropFKInfraClassSubClassSQL,
			dropFKInfraClassClassSQL,
			dropTableInfraClassSQL,
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
