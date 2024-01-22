CREATE TABLE IF NOT EXISTS subscriptions (
  id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  user_id UUID REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  payment_method_id UUID REFERENCES payment_methods(id) ON DELETE CASCADE NOT NULL,
  label VARCHAR(255) NOT NULL,
  provider VARCHAR(255) NOT NULL,
  amount INTEGER NOT NULL,
  interval INTEGER NOT NULL,
  period VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX subscriptions_id_idx ON subscriptions (id);
CREATE INDEX subscriptions_user_id_idx ON subscriptions (user_id);
CREATE INDEX subscriptions_payment_method_id_idx ON subscriptions (payment_method_id);

CREATE TRIGGER update_subscriptions_updated_at
BEFORE UPDATE ON subscriptions
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at_col();

---- create above / drop below ----

DROP TRIGGER IF EXISTS update_subscriptions_updated_at ON subscriptions;

DROP INDEX IF EXISTS subscriptions_id_idx;
DROP INDEX IF EXISTS subscriptions_user_id_idx;
DROP INDEX IF EXISTS subscriptions_payment_method_id_idx;

DROP TABLE IF EXISTS subscriptions;
