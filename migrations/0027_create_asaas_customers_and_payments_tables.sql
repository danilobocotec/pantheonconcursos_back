-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS asaas_customers (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asaas_id VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    cpf_cnpj VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    response_json TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS asaas_payments (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asaas_id VARCHAR(255) NOT NULL UNIQUE,
    customer_id VARCHAR(255) NOT NULL,
    value NUMERIC(12,2) NOT NULL,
    due_date VARCHAR(50) NOT NULL,
    billing_type VARCHAR(50) NOT NULL,
    request_json TEXT,
    response_json TEXT,
    confirmation_json TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS asaas_payments;
DROP TABLE IF EXISTS asaas_customers;

COMMIT;
