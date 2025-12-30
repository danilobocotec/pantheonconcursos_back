-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS vade_mecum_leis (
    id TEXT PRIMARY KEY,
    idtipo TEXT,
    tipo TEXT,
    nomecodigo TEXT NOT NULL,
    "Cabecalho" TEXT,
    "idPARTE" TEXT,
    "PARTE" TEXT,
    "PARTETEXTO" TEXT,
    idtitulo TEXT,
    titulo TEXT,
    titulotexto TEXT,
    idcapitulo TEXT,
    capitulo TEXT,
    capitulotexto TEXT,
    idsecao TEXT,
    secao TEXT,
    secaotexto TEXT,
    idsubsecao TEXT,
    subsecao TEXT,
    subsecaotexto TEXT,
    num_artigo TEXT,
    "Artigos" TEXT,
    "Ordem" TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS vade_mecum_leis;

COMMIT;
