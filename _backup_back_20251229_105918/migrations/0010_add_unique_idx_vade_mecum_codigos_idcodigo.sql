-- +goose Up
BEGIN;

CREATE UNIQUE INDEX IF NOT EXISTS idx_vade_mecum_codigos_idcodigo
    ON vade_mecum_codigos (idcodigo);

COMMIT;

-- +goose Down
BEGIN;

DROP INDEX IF EXISTS idx_vade_mecum_codigos_idcodigo;

COMMIT;
