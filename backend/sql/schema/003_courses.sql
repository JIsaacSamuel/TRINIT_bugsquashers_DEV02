-- +goose Up
CREATE TABLE course (
    id UUID PRIMARY KEY,
    langtaught TEXT NOT NULL,
    tutor_id UUID NOT NULL,
    price DECIMAL(6,2) NOT NULL,
    takenby INT NOT NULL DEFAULT(0),
    FOREIGN KEY (tutor_id) REFERENCES tutor (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE course;