package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSuperClassSQL = `
	CREATE TABLE IF NOT EXISTS public."super_class"
	(
		"super_class_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"super_class_name" varchar(500) UNIQUE NOT NULL,

		"phylum_id" UUID,
		"infra_phylum_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSuperClassSQL = `DROP TABLE IF EXISTS public."super_class";`

const addFKSuperClassPhylumSQL = `
	ALTER TABLE public."super_class"
	ADD CONSTRAINT fk_super_class_phylum
	FOREIGN KEY ("phylum_id") REFERENCES public."phylum" ("phylum_id")
`
const dropFKSuperClassPhylumSQL = `
	ALTER TABLE public."super_class"
	DROP CONSTRAINT fk_super_class_phylum
`

const addFKSuperClassInfraPhylumSQL = `
	ALTER TABLE public."super_class"
	ADD CONSTRAINT fk_super_class_infra_phylum
	FOREIGN KEY ("infra_phylum_id") REFERENCES public."infra_phylum" ("infra_phylum_id")
`
const dropFKSuperClassInfraPhylumSQL = `
	ALTER TABLE public."super_class"
	DROP CONSTRAINT fk_super_class_infra_phylum
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table super_class...")
		var scripts = [3]string{
			createTableSuperClassSQL,
			addFKSuperClassPhylumSQL,
			addFKSuperClassInfraPhylumSQL,
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

		fmt.Println("[Migration] Seeding table super_class...")
		superClasses, err := GetSuperClassData()
		if err != nil {
			return err
		}

		for _, superClass := range superClasses {

			if superClass.SuperClassName == "NULL" {
				continue
			}

			insertSuperClassSQL := fmt.Sprintf(`
			INSERT INTO public."super_class" ("super_class_name", "phylum_id", "infra_phylum_id") VALUES ('%s', 
			(SELECT "phylum_id" FROM public."phylum" WHERE "phylum_name" = '%s' LIMIT 1),
			(SELECT "infra_phylum_id" FROM public."infra_phylum" WHERE "infra_phylum_name" = '%s' LIMIT 1))`,
				superClass.SuperClassName, superClass.PhylumName, superClass.InfraPhylumName)

			_, err := db.Exec(insertSuperClassSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table super_class...")
		var scripts = [3]string{
			dropFKSuperClassPhylumSQL,
			dropTableSuperClassSQL,
			dropFKSuperClassInfraPhylumSQL,
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
