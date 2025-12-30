-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS vade_mecum_constituicao (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    registro_id TEXT,
    idtipo TEXT,
    tipo TEXT,
    cabecalho TEXT,
    idtitulo TEXT,
    titulo TEXT,
    textodotitulo TEXT,
    idcapitulo TEXT,
    capitulo TEXT,
    textocapitulo TEXT,
    idsecao TEXT,
    secao TEXT,
    textosecao TEXT,
    idsubsecao TEXT,
    subsecao TEXT,
    subsecaotexto TEXT,
    "Normativo" TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS vade_mecum_constituicao_registro_id_key
    ON vade_mecum_constituicao (registro_id);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS vade_mecum_constituicao;

COMMIT;
