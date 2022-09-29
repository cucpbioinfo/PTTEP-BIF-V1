package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableInfraPhylumSQL = `
	CREATE TABLE IF NOT EXISTS public."infra_phylum"
	(
		"infra_phylum_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"infra_phylum_name" varchar(500) UNIQUE NOT NULL,

		"phylum_id" UUID,
		"sub_phylum_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableInfraPhylumSQL = `DROP TABLE IF EXISTS public."sub_phylum";`

const addFKInfraPhylumPhylumSQL = `
	ALTER TABLE public."infra_phylum"
	ADD CONSTRAINT fk_infra_phylum_phylum
	FOREIGN KEY ("phylum_id") REFERENCES public."phylum" ("phylum_id")
`
const dropFKInfraPhylumPhylumSQL = `
	ALTER TABLE public."infra_phylum"
	DROP CONSTRAINT fk_infra_phylum_phylum
`

const addFKInfraPhylumSubPhylumSQL = `
	ALTER TABLE public."infra_phylum"
	ADD CONSTRAINT fk_infra_phylum_sub_phylum
	FOREIGN KEY ("sub_phylum_id") REFERENCES public."sub_phylum" ("sub_phylum_id")
`
const dropFKInfraPhylumSubPhylumSQL = `
	ALTER TABLE public."infra_phylum"
	DROP CONSTRAINT fk_infra_phylum_sub_phylum
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_phylum...")
		var scripts = [3]string{
			createTableInfraPhylumSQL,
			addFKInfraPhylumPhylumSQL,
			addFKInfraPhylumSubPhylumSQL,
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

		fmt.Println("[Migration] Seeding table sub_phylum...")
		infraPhylums, err := GetInfraPhylumData()
		if err != nil {
			return err
		}

		for _, infraPhylum := range infraPhylums {

			if infraPhylum.InfraPhylumName == "NULL" {
				continue
			}

			insertInfraPhylumSQL := fmt.Sprintf(`
			INSERT INTO public."infra_phylum" ("infra_phylum_name", "phylum_id", "sub_phylum_id") VALUES ('%s', 
			(SELECT "phylum_id" FROM public."phylum" WHERE "phylum_name" = '%s' LIMIT 1),
			(SELECT "sub_phylum_id" FROM public."sub_phylum" WHERE "sub_phylum_name" = '%s' LIMIT 1))`,
				infraPhylum.InfraPhylumName, infraPhylum.PhylumName, infraPhylum.SubPhylumName)

			_, err := db.Exec(insertInfraPhylumSQL)
			if err != nil {
				return err
			}
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table infra_phylum...")
		var scripts = [3]string{
			dropFKInfraPhylumSubPhylumSQL,
			dropFKInfraPhylumPhylumSQL,
			dropTableInfraPhylumSQL,
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
