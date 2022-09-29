package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableMajorGroupSQL = `
	CREATE TABLE IF NOT EXISTS public."major_group"
	(
		"major_group_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"major_group_name" varchar(500) UNIQUE NOT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableMajorGroupSQL = `DROP TABLE IF EXISTS public."major_group";`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table major_group...")
		var scripts = [1]string{
			createTableMajorGroupSQL,
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

		fmt.Println("[Migration] Seeding table major_group...")
		majorGroups, err := GetMajorGroupData()
		if err != nil {
			return err
		}

		for _, majorGroup := range majorGroups {
			insertMajorGroupSQL := fmt.Sprintf(`
				INSERT INTO public."major_group" ("major_group_name") VALUES ('%s')`, majorGroup.MajorGroupName)
			_, err := db.Exec(insertMajorGroupSQL)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table major_group...")
		_, err := db.Exec(dropTableMajorGroupSQL)
		return err
	})
}
