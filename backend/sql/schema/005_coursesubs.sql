-- +goose Up
CREATE TABLE coursesubs (
    id UUID PRIMARY KEY,
    tutor_id UUID NOT NULL,
    student_id UUID NOT NULL,
    FOREIGN KEY (tutor_id) REFERENCES tutor (id) ON DELETE CASCADE,
    FOREIGN KEY (student_id) REFERENCES student (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE coursesubs;