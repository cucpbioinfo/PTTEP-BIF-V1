package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createTableUserSQL = `
	CREATE TABLE IF NOT EXISTS public."user"
	(
		"user_id" UUID NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
		"email" varchar(500) UNIQUE NOT NULL,
		"password" text NOT NULL,
		"created_at" TIMESTAMPTZ DEFAULT NOW(),
		"updated_at" TIMESTAMPTZ DEFAULT NOW(),
		"deleted_at" TIMESTAMPTZ DEFAULT NULL
	);`
const dropTableUserSQL = `DROP TABLE IF EXISTS public."user";`

// func init() {
// 	migrations.MustRegisterTx(func(db migrations.DB) error {
// 		fmt.Println("[Migration] Creating table user...")
// 		_, err := db.Exec(createTableUserSQL)
// 		return err
// 	}, func(db migrations.DB) error {
// 		fmt.Println("[Migration] Droping table user...")
// 		_, err := db.Exec(dropTableUserSQL)
// 		return err
// 	})
// }

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table user...")
		var scripts = [1]string{
			createTableUserSQL,
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

		fmt.Println("[Migration] Seeding table user...")
		Users, err := GetUserData()
		if err != nil {
			fmt.Println("Cannot get even user data")
			return err
		}

		for _, usern := range Users {
			insertUserSQL := fmt.Sprintf(`
			INSERT INTO public."user"("email","password") VALUES(
				'%s','%s');`,
				usern.Email,
				usern.Password)
			// fmt.Println(insertUserSQL)
			_, err := db.Exec(insertUserSQL)
			if err != nil {
				fmt.Println(insertUserSQL)
				return err
			}
		}

		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table user...")
		var scripts = [1]string{
			dropTableUserSQL,
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
		return nil
	})
}
