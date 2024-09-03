-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR,
    password VARCHAR,
    name VARCHAR
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS users;

-- +goose StatementEnd
