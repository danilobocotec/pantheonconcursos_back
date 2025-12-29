-- +goose Up
BEGIN;

CREATE TABLE IF NOT EXISTS courses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    image TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_courses_user_id ON courses(user_id);

CREATE TABLE IF NOT EXISTS course_categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    image TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_course_categories_user_id ON course_categories(user_id);

ALTER TABLE courses ADD COLUMN IF NOT EXISTS category_id UUID;
ALTER TABLE courses ADD CONSTRAINT courses_category_id_fkey
    FOREIGN KEY (category_id) REFERENCES course_categories(id) ON DELETE SET NULL;
CREATE INDEX IF NOT EXISTS idx_courses_category_id ON courses(category_id);

CREATE TABLE IF NOT EXISTS course_modules (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    course_id UUID REFERENCES courses(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_course_modules_user_id ON course_modules(user_id);
CREATE INDEX IF NOT EXISTS idx_course_modules_course_id ON course_modules(course_id);

CREATE TABLE IF NOT EXISTS course_items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    module_id UUID REFERENCES course_modules(id) ON DELETE SET NULL,
    title VARCHAR(255) NOT NULL,
    tipo VARCHAR(100),
    conteudo TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_course_items_user_id ON course_items(user_id);
CREATE INDEX IF NOT EXISTS idx_course_items_module_id ON course_items(module_id);

COMMIT;

-- +goose Down
BEGIN;

DROP TABLE IF EXISTS course_items;
DROP TABLE IF EXISTS course_modules;

ALTER TABLE courses DROP CONSTRAINT IF EXISTS courses_category_id_fkey;
DROP INDEX IF EXISTS idx_courses_category_id;
DROP TABLE IF EXISTS course_categories;
DROP TABLE IF EXISTS courses;

COMMIT;
