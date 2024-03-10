-- +goose Up
ALTER TABLE tutor
ADD passcode VARCHAR(64) NOT NULL UNIQUE DEFAULT (
    123456
);

-- +goose Down
ALTER TABLE tutor
DROP COLUMN passcode;