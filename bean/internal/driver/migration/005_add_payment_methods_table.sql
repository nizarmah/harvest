CREATE TABLE IF NOT EXISTS payment_methods (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  label VARCHAR(255) NOT NULL,
  last4 VARCHAR(4) NOT NULL,
  brand VARCHAR(255) NOT NULL,
  exp_month INTEGER NOT NULL,
  exp_year INTEGER NOT NULL,
  is_default BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX payment_methods_id_idx ON payment_methods (id);
CREATE INDEX payment_methods_user_id_idx ON payment_methods (user_id);

CREATE TRIGGER update_payment_methods_updated_at
BEFORE UPDATE ON payment_methods
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_col();

---- create above / drop below ----

DROP TRIGGER IF EXISTS update_payment_methods_updated_at ON payment_methods;

DROP INDEX IF EXISTS payment_methods_id_idx;
DROP INDEX IF EXISTS payment_methods_user_id_idx;

DROP TABLE IF EXISTS payment_methods;
