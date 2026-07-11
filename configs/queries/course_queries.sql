-- name: GetAllCourses
SELECT * FROM course
WHERE 1 = 1
{{ if .Code }}
    AND code ILIKE '%' || {{ arg .Code }} || '%'
{{ end }}
{{ if .Name }}
    AND name ILIKE '%' || {{ arg .Name }} || '%'
{{ end }}
{{ if .SKS }}
    AND sks = {{ arg .SKS }}
{{ end }}
{{ if .SortBy }}
    ORDER BY {{ raw .SortBy }} {{ raw .SortDir }}
{{ end }}
LIMIT {{ arg .Limit }} OFFSET {{ arg .Offset }};

-- name: CountCourses
SELECT COUNT(*) FROM course
WHERE 1 = 1
{{ if .Code }}
    AND code ILIKE '%' || {{ arg .Code }} || '%'
{{ end }}
{{ if .Name }}
    AND name ILIKE '%' || {{ arg .Name }} || '%'
{{ end }}
{{ if .SKS }}
    AND sks = {{ arg .SKS }}
{{ end }};

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
