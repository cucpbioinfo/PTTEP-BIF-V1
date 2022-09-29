package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSuperPhylumSQL = `
	CREATE TABLE IF NOT EXISTS public."super_phylum"
	(
		"super_phylum_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"super_phylum_name" varchar(500) UNIQUE NOT NULL,

		"kingdom_id" UUID,
		"infra_kingdom_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSuperPhylumSQL = `DROP TABLE IF EXISTS public."super_phylum";`

const addFKSuperPhylumKingdomSQL = `
	ALTER TABLE public."super_phylum"
	ADD CONSTRAINT fk_super_phylum_kingdom
	FOREIGN KEY ("kingdom_id") REFERENCES public."kingdom" ("kingdom_id")
`
const dropFKSuperPhylumKingdomSQL = `
	ALTER TABLE public."super_phylum"
	DROP CONSTRAINT fk_super_phylum_kingdom
`

const addFKSuperPhylumInfraKingdomSQL = `
	ALTER TABLE public."super_phylum"
	ADD CONSTRAINT fk_super_phylum_infra_kingdom
	FOREIGN KEY ("infra_kingdom_id") REFERENCES public."infra_kingdom" ("infra_kingdom_id")
`
const dropFKSuperPhylumInfraKingdomSQL = `
	ALTER TABLE public."super_phylum"
	DROP CONSTRAINT fk_super_phylum_infra_kingdom
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table super_phylum...")
		var scripts = [3]string{
			createTableSuperPhylumSQL,
			addFKSuperPhylumKingdomSQL,
			addFKSuperPhylumInfraKingdomSQL,
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

		fmt.Println("[Migration] Seeding table super_phylum...")
		superPhylums, err := GetSuperPhylumData()
		if err != nil {
			return err
		}

		for _, superPhylum := range superPhylums {

			if superPhylum.SuperPhylumName == "NULL" {
				continue
			}

			insertSuperPhylumSQL := fmt.Sprintf(`
				INSERT INTO public."super_phylum" ("super_phylum_name", "kingdom_id", "infra_kingdom_id") VALUES ('%s', 
				(SELECT "kingdom_id" FROM public."kingdom" WHERE "kingdom_name" = '%s' LIMIT 1),
				(SELECT "infra_kingdom_id" FROM public."infra_kingdom" WHERE "infra_kingdom_name" = '%s' LIMIT 1))`,
				superPhylum.SuperPhylumName, superPhylum.KingdomName, superPhylum.InfraKingdomName)

			_, err := db.Exec(insertSuperPhylumSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table super_phylum...")
		var scripts = [3]string{
			dropFKSuperPhylumInfraKingdomSQL,
			dropFKSuperPhylumKingdomSQL,
			dropTableSuperPhylumSQL,
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
