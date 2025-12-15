-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS capa_vade_mecum_jurisprudencia (
    nomecodigo TEXT NOT NULL,
    "Cabecalho" TEXT,
    grupo TEXT,
    PRIMARY KEY (nomecodigo)
);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS capa_vade_mecum_jurisprudencia;

COMMIT;
