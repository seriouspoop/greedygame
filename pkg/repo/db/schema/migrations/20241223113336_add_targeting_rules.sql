-- +goose Up
-- +goose StatementBegin
CREATE TABLE targeting_rule (
    tid UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    cid UUID REFERENCES campaign ON DELETE CASCADE,
    app_include VARCHAR[],
    app_exclude VARCHAR[],
    country_include VARCHAR[],
    country_exclude VARCHAR[],
    os_include VARCHAR[],
    os_exclude VARCHAR[],
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER set_timestamp_targeting BEFORE
UPDATE ON campaign FOR EACH ROW EXECUTE FUNCTION trigger_set_timestamp ();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS set_timestamp_targeting ON targeting_rule;
DROP TABLE IF EXISTS targeting_rule;

-- +goose StatementEnd
