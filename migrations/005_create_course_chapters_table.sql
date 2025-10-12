-- +goose Up
CREATE TABLE course_chapters (
    id SERIAL PRIMARY KEY,
    course_id INTEGER NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    chapter_order INTEGER NOT NULL DEFAULT 1,
    is_published BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_course_chapters_course_id ON course_chapters(course_id);
CREATE INDEX idx_course_chapters_deleted_at ON course_chapters(deleted_at);
CREATE INDEX idx_course_chapters_chapter_order ON course_chapters(course_id, chapter_order);

-- +goose Down
DROP TABLE IF EXISTS course_chapters;