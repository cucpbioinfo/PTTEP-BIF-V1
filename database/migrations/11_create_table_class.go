package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableClassSQL = `
	CREATE TABLE IF NOT EXISTS public."class"
	(
		"class_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"class_name" varchar(500) UNIQUE NOT NULL,

		"phylum_id" UUID,
		"super_class_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableClassSQL = `DROP TABLE IF EXISTS public."class";`

const addFKClassPhylumSQL = `
	ALTER TABLE public."class"
	ADD CONSTRAINT fk_class_phylum
	FOREIGN KEY ("phylum_id") REFERENCES public."phylum" ("phylum_id")
`
const dropFKClassPhylumSQL = `
	ALTER TABLE public."class"
	DROP CONSTRAINT fk_class_phylum
`

const addFKClassSuperClassSQL = `
	ALTER TABLE public."class"
	ADD CONSTRAINT fk_class_super_class
	FOREIGN KEY ("super_class_id") REFERENCES public."super_class" ("super_class_id")
`
const dropFKClassSuperClassSQL = `
	ALTER TABLE public."class"
	DROP CONSTRAINT fk_class_super_class
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table class...")
		var scripts = [3]string{
			createTableClassSQL,
			addFKClassPhylumSQL,
			addFKClassSuperClassSQL,
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

		fmt.Println("[Migration] Seeding table class...")
		classes, err := GetClassData()
		if err != nil {
			return err
		}

		for _, class := range classes {

			if class.ClassName == "NULL" {
				continue
			}

			insertClassSQL := fmt.Sprintf(`
			INSERT INTO public."class" ("class_name", "phylum_id", "super_class_id") VALUES ('%s', 
			(SELECT "phylum_id" FROM public."phylum" WHERE "phylum_name" = '%s' LIMIT 1),
			(SELECT "super_class_id" FROM public."super_class" WHERE "super_class_name" = '%s' LIMIT 1))`,
				class.ClassName, class.PhylumName, class.SuperClassName)
			_, err := db.Exec(insertClassSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table class...")
		var scripts = [3]string{
			dropFKClassPhylumSQL,
			dropTableClassSQL,
			dropFKClassSuperClassSQL,
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
