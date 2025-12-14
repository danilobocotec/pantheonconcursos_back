-- +goose Up
BEGIN;

ALTER TABLE users
    RENAME COLUMN name TO full_name;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE users
    RENAME COLUMN full_name TO name;

COMMIT;
