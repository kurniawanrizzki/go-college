-- name: CreateCollege
INSERT INTO college (nim, name, semester, sks, active)
VALUES ({{ arg .NIM }}, {{ arg .Name }}, {{ arg .Semester }}, {{ arg .SKS }}, true)
RETURNING created_at, updated_at

-- name: FindColleges
SELECT * FROM college
WHERE 1 = 1
{{ if .NIM }}
    AND nim ILIKE '%' || {{ arg .NIM }} || '%'
{{ end }}
{{ if .Name }}
    AND name ILIKE '%' || {{ arg .Name }} || '%'
{{ end }}
{{ if .Semester }}
    AND semester = {{ arg .Semester }}
{{ end }}
{{ if .SortBy }}
    ORDER BY {{ raw .SortBy }} {{ raw .SortDir }}
{{ end }}
LIMIT {{ arg .Limit }} OFFSET {{ arg .Offset }};

-- name: CountColleges
SELECT COUNT(*) FROM college
WHERE 1 = 1
{{ if .NIM }}
    AND nim ILIKE '%' || {{ arg .NIM }} || '%'
{{ end }}
{{ if .Name }}
    AND name ILIKE '%' || {{ arg .Name }} || '%'
{{ end }}
{{ if .Semester }}
    AND semester = {{ arg .Semester }}
{{ end }};

-- name: FindCollegeByNim
SELECT * FROM college WHERE nim = {{ arg .NIM }}

-- name: UpdateCollege
UPDATE college
SET name = {{ arg .Name }}, semester = {{ arg .Semester }}, sks = {{ arg .SKS }}, active = {{ arg .Active }}, updated_at = {{ arg .UpdatedAt }}
WHERE nim = {{ arg .NIM }};

-- name: DeleteCollege
DELETE FROM college WHERE nim = {{ arg .NIM }};