BEGIN;

-- base function for setting updated_at columns
CREATE OR REPLACE FUNCTION on_update_timestamp()
  RETURNS trigger AS $$
  BEGIN
    NEW.updated_at = now();
    RETURN NEW;
  END;
$$ language 'plpgsql';


-- create wallets table
CREATE TABLE IF NOT EXISTS wallets (
  id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
  currency TEXT NOT NULL,
  balance DECIMAL(13,4) NOT NULL DEFAULT 0.0,
  account_id TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

-- add unique index on account_id
CREATE UNIQUE INDEX wallets_account_id ON wallets(account_id);

-- set updated_at field per update
CREATE TRIGGER wallets_updated_at 
  BEFORE UPDATE
    ON wallets FOR EACH ROW EXECUTE PROCEDURE on_update_timestamp();


COMMIT;