-- +goose Up
ALTER TABLE course
DROP COLUMN price;
ALTER TABLE course
ADD COLUMN price INT NOT NULL;

-- +goose Down
ALTER TABLE course
DROP COLUMN price;
ALTER TABLE course
ADD COLUMN price DECIMAL(6,2) NOT NULL;