package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableIdentificationSQL = `
	CREATE TABLE IF NOT EXISTS public."identification"
	(
		"identification_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"identification_name" varchar(500) UNIQUE NOT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableIdentificationSQL = `DROP TABLE IF EXISTS public."identification";`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table identification...")
		var scripts = [1]string{
			createTableIdentificationSQL,
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

		fmt.Println("[Migration] Seeding table identification...")
		identifications, err := GetIdentificationData()
		if err != nil {
			return err
		}

		for _, identification := range identifications {
			insertIdentificationSQL := fmt.Sprintf(`
				INSERT INTO public."identification" ("identification_name") VALUES ('%s')`, identification.IdentificationTechniqueName)
			_, err := db.Exec(insertIdentificationSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table identification...")
		_, err := db.Exec(dropTableIdentificationSQL)
		return err
	})
}
