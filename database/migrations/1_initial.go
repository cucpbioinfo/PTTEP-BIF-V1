package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createUUIDExtensionSQL = `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`
const dropUUIDExtensionSQL = `DROP EXTENSION IF EXISTS "uuid-ossp";`

const createSpeciesTypeEnumSQL = `CREATE TYPE species_types AS ENUM ('zooplankton', 'phytoplankton', 'nekton', 'benthos', 'bacteria');`
const dropSpeciesTypeEnumSQL = `DROP TYPE species_types`

const createFunctionSetTimestampSQL = `CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;`
const dropFunctionSetTimestampSQL = `DROP FUNCTION trigger_set_timestamp();`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Init database...")
		var scripts = [3]string{
			createUUIDExtensionSQL,
			createSpeciesTypeEnumSQL,
			createFunctionSetTimestampSQL,
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
	}, func(db migrations.DB) error {
		var scripts = [3]string{
			dropFunctionSetTimestampSQL,
			dropSpeciesTypeEnumSQL,
			dropUUIDExtensionSQL,
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
