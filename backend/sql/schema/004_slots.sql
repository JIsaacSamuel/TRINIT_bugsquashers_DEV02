-- +goose Up
CREATE TABLE slots (
    id UUID PRIMARY KEY,
    tutor_id UUID NOT NULL,
    fromtime TIME NOT NULL,
    totime TIME NOT NULL,
    taken BOOL NOT NULL DEFAULT(false),
    FOREIGN KEY (tutor_id) REFERENCES tutor (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE slots;