-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "pg_uuidv7";

CREATE TABLE IF NOT EXISTS users (
  id UUID DEFAULT uuid_generate_v7() PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(50) UNIQUE,
  password_hash VARCHAR(50) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),  
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
  -- Automatically indexed by PostgreSQL:
-- - id (as PRIMARY KEY)
-- - username (as UNIQUE)
-- - email (as UNIQUE)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "pg_uuidv7";
-- +goose StatementEnd
