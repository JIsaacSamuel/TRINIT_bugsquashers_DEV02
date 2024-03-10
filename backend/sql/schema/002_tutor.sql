-- +goose Up
CREATE TABLE tutor (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    emailID VARCHAR(64) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE tutor;