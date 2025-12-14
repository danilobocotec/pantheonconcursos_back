-- +goose Up
BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS plans (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(150) NOT NULL UNIQUE,
    description TEXT,
    phase VARCHAR(50) NOT NULL,
    duration VARCHAR(50) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS plan_id UUID NULL REFERENCES plans(id);

CREATE INDEX IF NOT EXISTS idx_users_plan_id ON users(plan_id);

INSERT INTO plans (name, description, phase, duration)
VALUES
    ('Plano 1 — Fase 1 (Anual)', 'Acesso anual completo à 1ª fase da OAB.', 'fase_1', 'anual'),
    ('Plano 2 — Fase 1 (Vitalício)', 'Acesso vitalício à 1ª fase da OAB.', 'fase_1', 'vitalicio'),
    ('Plano 3 — Fase 1 + Fase 2 (Anual)', 'Acesso anual à 1ª e 2ª fases da OAB.', 'fase_1_2', 'anual')
ON CONFLICT (name) DO NOTHING;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE users
    DROP COLUMN IF EXISTS plan_id;

DROP TABLE IF EXISTS plans;

COMMIT;
