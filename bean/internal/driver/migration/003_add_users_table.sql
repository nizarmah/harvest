CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  email VARCHAR(255) UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX users_id_idx ON users (id);
CREATE INDEX users_email_idx ON users (email);

CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_col();

---- create above / drop below ----

DROP TRIGGER IF EXISTS update_users_updated_at ON users;

DROP INDEX IF EXISTS users_id_idx;
DROP INDEX IF EXISTS users_email_idx;

DROP TABLE IF EXISTS users;
