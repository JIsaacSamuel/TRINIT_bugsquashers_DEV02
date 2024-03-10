-- name: ListCoursesSubs :many
SELECT coursesubs.id, course.coursename from coursesubs left join course on coursesubs.course_id = course.id where student_id = $1;

-- name: ListStudents :many
SELECT coursesubs.id, student.name from coursesubs left join student on coursesubs.student_id = student.id 
where course_id = $1;

-- name: CreateCourseSub :one
INSERT INTO coursesubs(id, student_id, course_id)
VALUES ($1, $2, $3)
RETURNING *;