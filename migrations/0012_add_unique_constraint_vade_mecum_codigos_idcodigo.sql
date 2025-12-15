-- +goose Up
BEGIN;

ALTER TABLE vade_mecum_codigos
    ADD CONSTRAINT vade_mecum_codigos_idcodigo_key
    UNIQUE (idcodigo);

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE vade_mecum_codigos
    DROP CONSTRAINT IF EXISTS vade_mecum_codigos_idcodigo_key;

COMMIT;
