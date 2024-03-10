-- name: CreateCourse :one
INSERT INTO course (id, langtaught, tutor_id, price, takenby, coursename)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: AllCourseLang :many
select course.langtaught, tutor.name, course.coursename, course.price as price, course.id
from course left join tutor on course.tutor_id = tutor.id
where  course.langtaught = $1 and price < $2;

-- name: AllcourseTutor :many
SELECT id, coursename from course where tutor_id = $1;

-- name: GetCourse :one
SELECT * from course where id = $1;