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

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating table user...")
		_, err := db.Exec(createTableUserSQL)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping table user...")
		_, err := db.Exec(dropTableUserSQL)
		return err
	})
}
