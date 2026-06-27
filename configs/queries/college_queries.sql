-- name: CreateCollege
INSERT INTO college (nim, name, semester, sks, active)
VALUES ({{ arg .NIM }}, {{ arg .Name }}, {{ arg .Semester }}, {{ sks }}, true)
RETURNING created_at, updated_at

--name: FindCollegeByNim
SELECT * FROM college WHERE nim = {{ arg .NIM }}

--name: FindCollegeByName
SELECT * FROM college WHERE name LIKE '%' || {{ arg .Name }} || '%'