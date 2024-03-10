-- +goose Up
ALTER TABLE coursesubs
DROP COLUMN tutor_id;
ALTER TABLE coursesubs
ADD course_id UUID NOT NULL;
ALTER TABLE coursesubs
ADD FOREIGN KEY (course_id) REFERENCES course(id);

-- +goose Down
ALTER TABLE coursesubs
DROP COLUMN course_id;
ALTER TABLE coursesubs
ADD tutor_id UUID;
