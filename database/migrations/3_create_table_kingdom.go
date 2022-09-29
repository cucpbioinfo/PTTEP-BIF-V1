package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableKingdomSQL = `
	CREATE TABLE IF NOT EXISTS public."kingdom"
	(
		"kingdom_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"kingdom_name" varchar(500) UNIQUE NOT NULL,

		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableKingdomSQL = `DROP TABLE IF EXISTS public."kingdom";`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table kingdom...")
		var scripts = [1]string{
			createTableKingdomSQL,
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

		fmt.Println("[Migration] Seeding table kingdom...")
		kingdoms, err := GetKingdomData()
		if err != nil {
			return err
		}

		for _, kingdom := range kingdoms {

			if kingdom.KingdomName == "NULL" {
				continue
			}

			insertKingdomSQL := fmt.Sprintf(`
				INSERT INTO public."kingdom" ("kingdom_name") VALUES ('%s')`, kingdom.KingdomName)

			_, err := db.Exec(insertKingdomSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table kingdom...")
		var scripts = [1]string{
			dropTableKingdomSQL,
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
