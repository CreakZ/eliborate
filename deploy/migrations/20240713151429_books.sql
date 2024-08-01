-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT NULL,
    category INTEGER,
    authors VARCHAR[] NOT NULL,
    is_foreign BOOLEAN NOT NULL,
    logo VARCHAR NULL,
    rack INTEGER,
    shelf INTEGER
);

ALTER SEQUENCE id RESTART WITH 1;

-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS books;

-- +goose StatementEnd
