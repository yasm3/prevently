-- +goose Up
CREATE TYPE push_status AS ENUM (
    'pending',
    'processing',
    'sent',
    'failed'
);
CREATE TABLE pushes(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    attempts INT NOT NULL DEFAULT 0,
    last_error TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    sent_at TIMESTAMP WITH TIME ZONE
);
CREATE INDEX idx_pushes_user_id ON pushes(user_id);
CREATE INDEX idx_pushes_status ON pushes(status);

-- +goose Down
DROP TYPE push_status;
DROP TABLE pushes;
DROP INDEX idx_pushes_user_id;
DROP INDEX idx_pushes_status;