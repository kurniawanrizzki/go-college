-- name: CreateEnrollment
INSERT INTO enrollment (nim, course_code, semester) VALUES ({{ arg .NIM }}, {{ arg .Course }}, {{ arg .Semester }})
RETURNING id, grade::text, created_at, udpated_at

-- name: UpdateEnrollment
UPDATE enrollment
SET semester = {{ arg .Semester }}, grade = {{ arg .Grade }}::"char"
WHERE nim = {{ arg .NIM }} AND course_code = {{ arg .Course }}

-- name: DeleteEnrollment
DELETE FROM enrollment WHERE id = {{ arg .ID }}

-- name: FindEnrollmentDetailByNim
SELECT c.code, c.name, c.sks, c.created_at, c.updated_at,
       e.semester, e.grade::text, e.created_at, e.udpated_at
FROM enrollment e
JOIN course c ON c.code = e.course_code
WHERE e.nim = {{ arg .NIM }}
