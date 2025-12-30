-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS vade_mecum_jurisprudencia (
    id TEXT PRIMARY KEY,
    idtipo TEXT,
    tipo TEXT,
    idcodigo TEXT,
    nomecodigo TEXT,
    "Cabecalho" TEXT,
    "Tipo" TEXT,
    idramo TEXT,
    ramotexto TEXT,
    idassunto TEXT,
    assuntotexto TEXT,
    idenunciado TEXT,
    "Enunciado" TEXT,
    idsecao TEXT,
    secao TEXT,
    secaotexto TEXT,
    idsubsecao TEXT,
    subsecao TEXT,
    subsecaotexto TEXT,
    num_artigo TEXT,
    "Normativo" TEXT,
    "Ordem" TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS vade_mecum_jurisprudencia_unique_idx
    ON vade_mecum_jurisprudencia (nomecodigo, "Normativo", num_artigo, "Enunciado");

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS vade_mecum_jurisprudencia;

COMMIT;
