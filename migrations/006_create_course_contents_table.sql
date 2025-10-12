-- +goose Up
CREATE TABLE course_contents (
    id SERIAL PRIMARY KEY,
    chapter_id INTEGER NOT NULL REFERENCES course_chapters(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    content_type VARCHAR(50) NOT NULL, -- 'video', 'text', 'image', 'pdf', 'link', etc.
    file_url VARCHAR(500),
    content_text TEXT,
    content_order INTEGER NOT NULL DEFAULT 1,
    is_published BOOLEAN NOT NULL DEFAULT false,
    duration_minutes INTEGER DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_course_contents_chapter_id ON course_contents(chapter_id);
CREATE INDEX idx_course_contents_deleted_at ON course_contents(deleted_at);
CREATE INDEX idx_course_contents_content_type ON course_contents(content_type);
CREATE INDEX idx_course_contents_content_order ON course_contents(chapter_id, content_order);

-- +goose Down
DROP TABLE IF EXISTS course_contents;