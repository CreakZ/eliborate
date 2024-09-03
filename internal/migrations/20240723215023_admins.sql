-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS admin_users (
    id SERIAL PRIMARY KEY,
    login VARCHAR,
    password VARCHAR,
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS admin_users;

-- +goose StatementEnd
