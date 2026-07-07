-- name: CreateCollege
INSERT INTO college (nim, name, semester, sks, active)
VALUES ({{ arg .NIM }}, {{ arg .Name }}, {{ arg .Semester }}, {{ arg .SKS }}, true)
RETURNING created_at, updated_at

-- name: FindColleges
SELECT * FROM college

-- name: FindCollegeByNim
SELECT * FROM college WHERE nim = {{ arg .NIM }}

-- name: FindCollegeBySemester
SELECT * FROM college WHERE semester = {{ arg .Semester }}

-- name: FindCollegeByName
SELECT * FROM college WHERE name ILIKE '%' || {{ arg .Name }} || '%'

-- name: UpdateCollege
UPDATE college
SET name = {{ arg .Name }}, semester = {{ arg .Semester }}, sks = {{ arg .SKS }}, active = {{ arg .Active }}, updated_at = {{ arg .UpdatedAt }}
WHERE nim = {{ arg .NIM }};

-- name: DeleteCollege
DELETE FROM college WHERE nim = {{ arg .NIM }};