-- +goose Up
ALTER TABLE course
ADD coursename VARCHAR(64) NOT NULL UNIQUE;

-- +goose Down
ALTER TABLE course
DROP COLUMN coursename;