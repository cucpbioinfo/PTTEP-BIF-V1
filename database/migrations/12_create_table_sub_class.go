package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubClassSQL = `
	CREATE TABLE IF NOT EXISTS public."sub_class"
	(
		"sub_class_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_class_name" varchar(500) UNIQUE NOT NULL,

		"class_id" UUID,
		
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubClassSQL = `DROP TABLE IF EXISTS public."sub_class";`

const addFKSubClassClassSQL = `
	ALTER TABLE public."sub_class"
	ADD CONSTRAINT fk_sub_class_class
	FOREIGN KEY ("class_id") REFERENCES public."class" ("class_id")
`
const dropFKSubClassClassSQL = `
	ALTER TABLE public."sub_class"
	DROP CONSTRAINT fk_sub_class_class
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_class...")
		var scripts = [2]string{
			createTableSubClassSQL,
			addFKSubClassClassSQL,
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
		subClasses, err := GetSubClassData()
		if err != nil {
			return err
		}

		for _, subClass := range subClasses {

			if subClass.SubClassName == "NULL" {
				continue
			}

			insertSubClassSQL := fmt.Sprintf(`
			INSERT INTO public."sub_class" ("sub_class_name", "class_id") VALUES ('%s', 
			(SELECT "class_id" FROM public."class" WHERE "class_name" = '%s' LIMIT 1))`,
				subClass.SubClassName, subClass.ClassName)

			_, err := db.Exec(insertSubClassSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_class...")
		var scripts = [2]string{
			dropFKSubClassClassSQL,
			dropTableSubClassSQL,
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
