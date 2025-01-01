-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS course_course_modules (
    course_id UUID NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    course_module_id UUID NOT NULL REFERENCES course_modules(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (course_id, course_module_id)
);

CREATE INDEX IF NOT EXISTS idx_course_course_modules_course_id ON course_course_modules(course_id);
CREATE INDEX IF NOT EXISTS idx_course_course_modules_module_id ON course_course_modules(course_module_id);

DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_name = 'course_modules'
          AND column_name = 'course_id'
    ) THEN
        EXECUTE '
            INSERT INTO course_course_modules (course_id, course_module_id)
            SELECT course_id, id
            FROM course_modules
            WHERE course_id IS NOT NULL
            ON CONFLICT DO NOTHING
        ';
    END IF;
END $$;

ALTER TABLE course_modules DROP COLUMN IF EXISTS course_id;

COMMIT;

-- +goose Down
BEGIN;

ALTER TABLE course_modules ADD COLUMN IF NOT EXISTS course_id UUID;
ALTER TABLE course_modules ADD CONSTRAINT course_modules_course_id_fkey
    FOREIGN KEY (course_id) REFERENCES courses(id) ON DELETE SET NULL;
CREATE INDEX IF NOT EXISTS idx_course_modules_course_id ON course_modules(course_id);

UPDATE course_modules AS m
SET course_id = sub.course_id
FROM (
    SELECT course_module_id, (array_agg(course_id ORDER BY course_id))[1] AS course_id
    FROM course_course_modules
    GROUP BY course_module_id
) AS sub
WHERE m.id = sub.course_module_id;

DROP TABLE IF EXISTS course_course_modules;

COMMIT;
