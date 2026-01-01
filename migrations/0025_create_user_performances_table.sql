-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS user_performances (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    total_questions INTEGER NOT NULL,
    correct_questions INTEGER NOT NULL,
    wrong_questions INTEGER NOT NULL,
    accuracy_percent DOUBLE PRECISION NOT NULL,
    recorded_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_performances_user_id ON user_performances(user_id);
CREATE INDEX IF NOT EXISTS idx_user_performances_recorded_at ON user_performances(recorded_at);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS user_performances;

COMMIT;
