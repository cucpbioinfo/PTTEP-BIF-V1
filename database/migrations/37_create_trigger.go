package main

import (
	"fmt"

	"github.com/go-pg/migrations/v8"
)

const createUpdatedAtTriggerQL = `
DO $$
DECLARE
    t text;
BEGIN
    FOR t IN 
        SELECT table_name FROM information_schema.columns
        WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER set_updated_at
                        BEFORE UPDATE ON %I
                        FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp()',
                        t);
    END LOOP;
END;
$$ LANGUAGE plpgsql;
`
const dropUpdatedAtTriggerSQL = `
DO $$
DECLARE
    t text;
BEGIN
    FOR t IN 
        SELECT table_name FROM information_schema.columns
        WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('DROP TRIGGER set_updated_at');
    END LOOP;
END;
$$ LANGUAGE plpgsql;
`

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("[Migration] Creating updated_at trigger...")
		_, err := db.Exec(createUpdatedAtTriggerQL)
		if err != nil {
			return err
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("[Migration] Droping updated_at trigger...")
		_, err := db.Exec(dropUpdatedAtTriggerSQL)
		return err
	})
}
