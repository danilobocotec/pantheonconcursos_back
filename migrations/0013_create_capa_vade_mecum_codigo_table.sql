-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS capa_vade_mecum_codigo (
    id TEXT PRIMARY KEY DEFAULT md5(random()::text || clock_timestamp()::text),
    nomecodigo TEXT NOT NULL,
    "Cabecalho" TEXT,
    grupo TEXT
);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS capa_vade_mecum_codigo;

COMMIT;
