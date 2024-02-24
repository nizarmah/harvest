CREATE TABLE IF NOT EXISTS memberships (
  user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP NULL
);

CREATE INDEX memberships_user_id_idx ON memberships (user_id);

---- create above / drop below ----

DROP INDEX IF EXISTS memberships_user_id_idx;

DROP TABLE IF EXISTS memberships;
