-- +goose Up
Create table flashcard (
    id UUID PRIMARY KEY,
    course_id UUID NOT NULL,
    word VARCHAR(64) NOT NULL,
    meaning VARCHAR(64) NOT NULL,
    FOREIGN KEY (course_id) REFERENCES course(id) ON DELETE CASCADE
);

-- +goose Down
DROP table flashcard;