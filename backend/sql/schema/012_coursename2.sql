-- +goose Up
ALTER TABLE course DROP coursename;
ALTER TABLE course
ADD coursename VARCHAR(64);

-- +goose Down
ALTER TABLE course
DROP COLUMN coursename;