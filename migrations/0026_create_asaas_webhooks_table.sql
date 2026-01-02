-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS asaas_webhooks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asaas_id VARCHAR(255) UNIQUE,
    name VARCHAR(255) NOT NULL,
    url TEXT NOT NULL,
    email VARCHAR(255),
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    interrupted BOOLEAN NOT NULL DEFAULT FALSE,
    auth_token TEXT,
    send_type VARCHAR(50),
    events JSONB NOT NULL,
    response_json JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS asaas_webhooks;

COMMIT;
