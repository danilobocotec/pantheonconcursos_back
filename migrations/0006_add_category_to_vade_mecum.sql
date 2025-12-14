-- +goose Up
BEGIN;

ALTER TABLE vade_mecum
    ADD COLUMN IF NOT EXISTS category VARCHAR(50) NOT NULL DEFAULT 'constituicao';

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE vade_mecum
    DROP COLUMN IF EXISTS category;

COMMIT;
