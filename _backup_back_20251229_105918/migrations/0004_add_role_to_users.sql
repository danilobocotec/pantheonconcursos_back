-- +goose Up
BEGIN;

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS role VARCHAR(20) NOT NULL DEFAULT 'user';

UPDATE users
SET role = 'user'
WHERE role IS NULL
   OR role = '';

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE users
    DROP COLUMN IF EXISTS role;

COMMIT;
