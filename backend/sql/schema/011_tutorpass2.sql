-- +goose Up
ALTER TABLE tutor
DROP COLUMN passcode;
ALTER TABLE tutor
ADD passcode VARCHAR(64) NOT NULL DEFAULT (
    123456
);

-- +goose Down
ALTER TABLE tutor
DROP COLUMN passcode;