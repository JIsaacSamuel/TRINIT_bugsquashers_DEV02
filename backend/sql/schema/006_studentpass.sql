-- +goose Up
ALTER TABLE student
ADD passcode VARCHAR(64) NOT NULL UNIQUE DEFAULT (
    123456
);

-- +goose Down
ALTER TABLE student
DROP COLUMN passcode;