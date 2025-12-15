-- +goose Up
BEGIN;

-- códigos
ALTER TABLE capa_vade_mecum_codigo ADD COLUMN IF NOT EXISTS id TEXT;
UPDATE capa_vade_mecum_codigo
SET id = md5(random()::text || clock_timestamp()::text)
WHERE id IS NULL OR id = '';
ALTER TABLE capa_vade_mecum_codigo ALTER COLUMN id SET DEFAULT md5(random()::text || clock_timestamp()::text);
ALTER TABLE capa_vade_mecum_codigo ALTER COLUMN id SET NOT NULL;
ALTER TABLE capa_vade_mecum_codigo DROP CONSTRAINT IF EXISTS capa_vade_mecum_codigo_pkey;
ALTER TABLE capa_vade_mecum_codigo ADD CONSTRAINT capa_vade_mecum_codigo_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX IF NOT EXISTS capa_vade_mecum_codigo_nomecodigo_key ON capa_vade_mecum_codigo (nomecodigo);
DROP INDEX IF EXISTS capa_vade_mecum_codigo_grupo_uniq;
CREATE UNIQUE INDEX IF NOT EXISTS capa_vade_mecum_codigo_grupo_uniq
    ON capa_vade_mecum_codigo (grupo) WHERE grupo IS NOT NULL AND btrim(grupo) <> '';

-- oab
ALTER TABLE capa_vade_mecum_oab ADD COLUMN IF NOT EXISTS id TEXT;
UPDATE capa_vade_mecum_oab
SET id = md5(random()::text || clock_timestamp()::text)
WHERE id IS NULL OR id = '';
ALTER TABLE capa_vade_mecum_oab ALTER COLUMN id SET DEFAULT md5(random()::text || clock_timestamp()::text);
ALTER TABLE capa_vade_mecum_oab ALTER COLUMN id SET NOT NULL;
ALTER TABLE capa_vade_mecum_oab DROP CONSTRAINT IF EXISTS capa_vade_mecum_oab_pkey;
ALTER TABLE capa_vade_mecum_oab ADD CONSTRAINT capa_vade_mecum_oab_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX IF NOT EXISTS capa_vade_mecum_oab_nomecodigo_key ON capa_vade_mecum_oab (nomecodigo);
DROP INDEX IF EXISTS capa_vade_mecum_oab_grupo_uniq;
CREATE UNIQUE INDEX IF NOT EXISTS capa_vade_mecum_oab_grupo_uniq
    ON capa_vade_mecum_oab (grupo) WHERE grupo IS NOT NULL AND btrim(grupo) <> '';

-- jurisprudência
ALTER TABLE capa_vade_mecum_jurisprudencia ADD COLUMN IF NOT EXISTS id TEXT;
UPDATE capa_vade_mecum_jurisprudencia
SET id = md5(random()::text || clock_timestamp()::text)
WHERE id IS NULL OR id = '';
ALTER TABLE capa_vade_mecum_jurisprudencia ALTER COLUMN id SET DEFAULT md5(random()::text || clock_timestamp()::text);
ALTER TABLE capa_vade_mecum_jurisprudencia ALTER COLUMN id SET NOT NULL;
ALTER TABLE capa_vade_mecum_jurisprudencia DROP CONSTRAINT IF EXISTS capa_vade_mecum_jurisprudencia_pkey;
ALTER TABLE capa_vade_mecum_jurisprudencia ADD CONSTRAINT capa_vade_mecum_jurisprudencia_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX IF NOT EXISTS capa_vade_mecum_jurisprudencia_nomecodigo_key ON capa_vade_mecum_jurisprudencia (nomecodigo);
DROP INDEX IF EXISTS capa_vade_mecum_jurisprudencia_grupo_uniq;
CREATE UNIQUE INDEX IF NOT EXISTS capa_vade_mecum_jurisprudencia_grupo_uniq
    ON capa_vade_mecum_jurisprudencia (grupo) WHERE grupo IS NOT NULL AND btrim(grupo) <> '';

COMMIT;

-- +goose Down
BEGIN;

-- códigos
ALTER TABLE capa_vade_mecum_codigo DROP CONSTRAINT IF EXISTS capa_vade_mecum_codigo_pkey;
ALTER TABLE capa_vade_mecum_codigo ADD CONSTRAINT capa_vade_mecum_codigo_pkey PRIMARY KEY (nomecodigo);
DROP INDEX IF EXISTS capa_vade_mecum_codigo_nomecodigo_key;
DROP INDEX IF EXISTS capa_vade_mecum_codigo_grupo_uniq;
ALTER TABLE capa_vade_mecum_codigo DROP COLUMN IF EXISTS id;

-- oab
ALTER TABLE capa_vade_mecum_oab DROP CONSTRAINT IF EXISTS capa_vade_mecum_oab_pkey;
ALTER TABLE capa_vade_mecum_oab ADD CONSTRAINT capa_vade_mecum_oab_pkey PRIMARY KEY (nomecodigo);
DROP INDEX IF EXISTS capa_vade_mecum_oab_nomecodigo_key;
DROP INDEX IF EXISTS capa_vade_mecum_oab_grupo_uniq;
ALTER TABLE capa_vade_mecum_oab DROP COLUMN IF EXISTS id;

-- jurisprudência
ALTER TABLE capa_vade_mecum_jurisprudencia DROP CONSTRAINT IF EXISTS capa_vade_mecum_jurisprudencia_pkey;
ALTER TABLE capa_vade_mecum_jurisprudencia ADD CONSTRAINT capa_vade_mecum_jurisprudencia_pkey PRIMARY KEY (nomecodigo);
DROP INDEX IF EXISTS capa_vade_mecum_jurisprudencia_nomecodigo_key;
DROP INDEX IF EXISTS capa_vade_mecum_jurisprudencia_grupo_uniq;
ALTER TABLE capa_vade_mecum_jurisprudencia DROP COLUMN IF EXISTS id;

COMMIT;
