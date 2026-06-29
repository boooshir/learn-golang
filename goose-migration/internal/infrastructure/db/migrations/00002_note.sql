-- +goose Up
-- +goose StatementBegin

-- CREATE EXTENSION IF NOT EXISTS "pg_uuidv7";

CREATE TABLE IF NOT EXISTS notes (
  id UUID DEFAULT uuid_generate_v7() PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  content VARCHAR(255) NOT NULL,
  user_id UUID,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_users FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX idx_notes_user_id ON notes(user_id)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notes;
-- remove the index
DROP INDEX IF EXISTS idx_notes_user_id;
-- DROP EXTENSION IF EXISTS "pg_uuidv7";
-- +goose StatementEnd
