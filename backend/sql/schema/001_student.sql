-- +goose Up
CREATE TABLE student (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    emailID VARCHAR(64) NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE student;