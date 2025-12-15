-- +goose Up
CREATE TABLE IF NOT EXISTS vade_mecum_oab (
    id TEXT PRIMARY KEY,
    idtipo TEXT,
    tipo TEXT,
    nomecodigo TEXT NOT NULL,
    "Cabecalho" TEXT,
    titulo TEXT,
    titulotexto TEXT,
    titulo_label TEXT,
    capitulo TEXT,
    capitulotexto TEXT,
    capitulo_label TEXT,
    secao TEXT,
    secaotexto TEXT,
    secao_label TEXT,
    subsecao TEXT,
    subsecaotexto TEXT,
    subsecao_label TEXT,
    num_artigo TEXT,
    "Artigos" TEXT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS vade_mecum_oab_unique_idx ON vade_mecum_oab (nomecodigo, num_artigo, "Artigos");

-- +goose Down
DROP TABLE IF EXISTS vade_mecum_oab;
