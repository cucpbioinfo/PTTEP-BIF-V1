package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableSubSpeciesSQL = `
	CREATE TABLE IF NOT EXISTS public."sub_species"
	(
		"sub_species_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"sub_species_name" varchar(500) UNIQUE NOT NULL,

		"species_id" UUID,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableSubSpeciesSQL = `DROP TABLE IF EXISTS public."sub_species";`

const addFKSubSpeciesSpeciesSQL = `
	ALTER TABLE public."sub_species"
	ADD CONSTRAINT fk_sub_species_species
	FOREIGN KEY ("species_id") REFERENCES public."species" ("species_id")
;`
const dropFKSubSpeciesSpeciesSQL = `
	ALTER TABLE public."sub_species"
	DROP CONSTRAINT fk_sub_species_species
;`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table sub_species...")
		var scripts = [2]string{
			createTableSubSpeciesSQL,
			addFKSubSpeciesSpeciesSQL,
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

		fmt.Println("[Migration] Seeding table sub_species...")
		subSpecies, err := GetSubSpeciesData()
		if err != nil {
			return err
		}

		for _, s := range subSpecies {

			if s.SubSpeciesName == "NULL" {
				continue
			}

			insertKingdomSQL := fmt.Sprintf(`
				INSERT INTO public."sub_species" ("sub_species_name", "species_id") VALUES ('%s', 
				(SELECT species_id FROM public."species" WHERE "species_name" = '%s' LIMIT 1))`,
				s.SubSpeciesName, s.SpeciesName)

			_, err := db.Exec(insertKingdomSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table sub_species...")
		var scripts = [2]string{
			dropFKSubSpeciesSpeciesSQL,
			dropTableSubSpeciesSQL,
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
