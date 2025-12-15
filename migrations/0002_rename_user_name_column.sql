-- +goose Up
BEGIN;

DO $rename_up$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users'
          AND column_name = 'name'
    ) THEN
        ALTER TABLE users
            RENAME COLUMN name TO full_name;
    END IF;
END;
$rename_up$;

COMMIT;

-- +goose Down
BEGIN;

DO $rename_down$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'users'
          AND column_name = 'full_name'
    ) THEN
        ALTER TABLE users
            RENAME COLUMN full_name TO name;
    END IF;
END;
$rename_down$;

COMMIT;
