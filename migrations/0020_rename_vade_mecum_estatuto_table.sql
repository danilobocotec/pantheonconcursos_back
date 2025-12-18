-- +goose Up
BEGIN;

-- Rename table to match GORM pluralization
ALTER TABLE IF EXISTS vade_mecum_estatuto
    RENAME TO vade_mecum_estatutos;

-- Rename unique index if it exists
ALTER INDEX IF EXISTS vade_mecum_estatuto_idcodigo_key
    RENAME TO vade_mecum_estatutos_idcodigo_key;

-- Ensure Artigos column can store large texts
ALTER TABLE IF EXISTS vade_mecum_estatutos
    ALTER COLUMN "Artigos" TYPE TEXT;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE IF EXISTS vade_mecum_estatutos
    RENAME TO vade_mecum_estatuto;

ALTER INDEX IF EXISTS vade_mecum_estatutos_idcodigo_key
    RENAME TO vade_mecum_estatuto_idcodigo_key;

ALTER TABLE IF EXISTS vade_mecum_estatuto
    ALTER COLUMN "Artigos" TYPE VARCHAR(5000);

COMMIT;
