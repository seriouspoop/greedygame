-- +goose Up
-- +goose StatementBegin
INSERT INTO targeting_rule (cid, app_include, app_exclude, country_include, country_exclude, os_include, os_exclude) VALUES 
    (to_uuid('c2d29867-3d0b-d497-9191-18a9d8ee7830'), ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[], ARRAY['india', 'us'], ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[]),
    (to_uuid('d2d29867-3d0b-d497-9191-18a9d8ee7830'), ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[], ARRAY['us'], ARRAY['android', 'ios'], ARRAY[]::VARCHAR[]),
    (to_uuid('e2d29867-3d0b-d497-9191-18a9d8ee7830'), ARRAY['com.gametion.ludokinggame'], ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[], ARRAY[]::VARCHAR[], ARRAY['android'], ARRAY[]::VARCHAR[]);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE targeting_rule CASCADE;
-- +goose StatementEnd
