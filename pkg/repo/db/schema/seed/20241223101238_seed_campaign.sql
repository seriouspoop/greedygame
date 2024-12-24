-- +goose Up
-- +goose StatementBegin
INSERT INTO campaign (cid, name, image, cta, status) VALUES
    (to_uuid('c2d29867-3d0b-d497-9191-18a9d8ee7830'), 'Spotify - Music for everyone', 'https://somelink', 'Download', 'active'),
    (to_uuid('d2d29867-3d0b-d497-9191-18a9d8ee7830'), 'Duolingo: Best way to learn', 'https://somelink2', 'Install', 'active'),
    (to_uuid('e2d29867-3d0b-d497-9191-18a9d8ee7830'), 'Subway Surfer', 'https://somelink3', 'Play', 'active');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE campaign CASCADE;
-- +goose StatementEnd
