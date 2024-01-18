CREATE TABLE IF NOT EXISTS tokens (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  email VARCHAR(255) UNIQUE NOT NULL,
  hashed_token VARCHAR(60) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '10 minutes'
);

CREATE INDEX tokens_id_idx ON tokens (id);
CREATE INDEX tokens_email_idx ON tokens (email);

CREATE OR REPLACE FUNCTION overwrite_token_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.created_at = CURRENT_TIMESTAMP;
    NEW.expires_at = CURRENT_TIMESTAMP + INTERVAL '10 minutes';
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER overwrite_token
BEFORE UPDATE ON tokens
FOR EACH ROW
EXECUTE PROCEDURE overwrite_token_timestamp();

---- create above / drop below ----

DROP TRIGGER IF EXISTS overwrite_token ON tokens;

DROP FUNCTION IF EXISTS overwrite_token_timestamp() CASCADE;

DROP INDEX IF EXISTS tokens_id_idx;
DROP INDEX IF EXISTS tokens_email_idx;

DROP TABLE IF EXISTS tokens;
