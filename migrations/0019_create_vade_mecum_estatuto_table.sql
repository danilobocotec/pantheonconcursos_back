-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS vade_mecum_estatuto (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    idtipo TEXT,
    tipo TEXT,
    idcodigo TEXT,
    nomecodigo TEXT,
    "Cabecalho" TEXT,
    "PARTE" TEXT,
    idlivro TEXT,
    livro TEXT,
    livrotexto TEXT,
    idtitulo TEXT,
    titulo TEXT,
    titulotexto TEXT,
    idsubtitulo TEXT,
    subtitulo TEXT,
    subtitulotexto TEXT,
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
    "Artigos" VARCHAR(5000),
    "Ordem" TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS vade_mecum_estatuto_idcodigo_key
    ON vade_mecum_estatuto (idcodigo);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS vade_mecum_estatuto;

COMMIT;
