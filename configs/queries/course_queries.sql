-- name: GetAllCourses
SELECT * FROM course

-- name: CreateCourse
INSERT INTO course (code, name, sks) VALUES ({{ arg .Code }}, {{ arg .Name }}, {{ arg .SKS }})
RETURNING created_at, updated_at

-- name: UpdateCourse
UPDATE course
set name = {{ arg .Name }}, sks = {{ arg .SKS }} 
WHERE code = {{ arg .Code }}

-- name: DeleteCourse
DELETE FROM course WHERE code = {{ arg .Code }}

-- name: FindCourseByCode
SELECT * FROM course WHERE code = {{ arg .Code }};
