-- +goose Up
BEGIN;

ALTER TABLE vade_mecum
    ADD COLUMN IF NOT EXISTS header TEXT,
    ADD COLUMN IF NOT EXISTS title_id TEXT,
    ADD COLUMN IF NOT EXISTS title_name TEXT,
    ADD COLUMN IF NOT EXISTS title_text TEXT,
    ADD COLUMN IF NOT EXISTS chapter_id TEXT,
    ADD COLUMN IF NOT EXISTS chapter_name TEXT,
    ADD COLUMN IF NOT EXISTS chapter_text TEXT;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE vade_mecum
    DROP COLUMN IF EXISTS header,
    DROP COLUMN IF EXISTS title_id,
    DROP COLUMN IF EXISTS title_name,
    DROP COLUMN IF EXISTS title_text,
    DROP COLUMN IF EXISTS chapter_id,
    DROP COLUMN IF EXISTS chapter_name,
    DROP COLUMN IF EXISTS chapter_text;

COMMIT;
