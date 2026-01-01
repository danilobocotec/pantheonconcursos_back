-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS course_module_items (
    course_module_id UUID NOT NULL REFERENCES course_modules(id) ON DELETE CASCADE,
    course_item_id UUID NOT NULL REFERENCES course_items(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (course_module_id, course_item_id)
);

CREATE INDEX IF NOT EXISTS idx_course_module_items_module_id ON course_module_items(course_module_id);
CREATE INDEX IF NOT EXISTS idx_course_module_items_item_id ON course_module_items(course_item_id);

INSERT INTO course_module_items (course_module_id, course_item_id)
SELECT module_id, id
FROM course_items
WHERE module_id IS NOT NULL
ON CONFLICT DO NOTHING;

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS course_module_items;

COMMIT;
