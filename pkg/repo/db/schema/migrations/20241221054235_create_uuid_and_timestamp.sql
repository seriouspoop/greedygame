-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET intervalstyle = 'iso_8601';

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION to_uuid(raw text)
  RETURNS uuid IMMUTABLE STRICT
AS $$
  BEGIN
    RETURN raw::uuid;
  EXCEPTION WHEN invalid_text_representation THEN
    RETURN uuid_in(overlay(overlay(md5(raw) placing '4' from 13) placing '8' from 17)::cstring);
  END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP FUNCTION IF EXISTS trigger_set_timestamp() CASCADE;
DROP FUNCTION IF EXISTS to_uuid() CASCADE;
-- +goose StatementEnd
