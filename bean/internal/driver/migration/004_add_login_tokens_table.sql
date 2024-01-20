CREATE TABLE IF NOT EXISTS login_tokens (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  email VARCHAR(255) UNIQUE NOT NULL,
  hashed_token VARCHAR(60) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '10 minutes'
);

CREATE INDEX login_tokens_id_idx ON login_tokens (id);
CREATE INDEX login_tokens_email_idx ON login_tokens (email);

CREATE OR REPLACE FUNCTION overwrite_login_token_timestamps()
RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at = CURRENT_TIMESTAMP;
    NEW.expires_at = CURRENT_TIMESTAMP + INTERVAL '10 minutes';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER overwrite_login_token
BEFORE UPDATE ON login_tokens
FOR EACH ROW
EXECUTE PROCEDURE overwrite_login_token_timestamps();

---- create above / drop below ----

DROP TRIGGER IF EXISTS overwrite_login_token ON login_tokens;

DROP FUNCTION IF EXISTS overwrite_login_token_timestamps() CASCADE;

DROP INDEX IF EXISTS login_tokens_id_idx;
DROP INDEX IF EXISTS login_tokens_email_idx;

DROP TABLE IF EXISTS login_tokens;
