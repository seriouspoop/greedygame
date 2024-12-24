-- +goose Up
-- +goose StatementBegin
CREATE TYPE campaign_status AS ENUM ('active', 'inactive');

CREATE TABLE campaign (
    cid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    cta VARCHAR(255) NOT NULL,
    status campaign_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER set_timestamp BEFORE
UPDATE ON campaign FOR EACH ROW EXECUTE FUNCTION trigger_set_timestamp ();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp ON campaign;
DROP TABLE IF EXISTS campaign;
DROP TYPE IF EXISTS campaign_status;
-- +goose StatementEnd
