CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description TEXT NULL,
    category INTEGER,
    authors VARCHAR[] NOT NULL,
    is_foreign BOOLEAN NOT NULL,
    cover_url VARCHAR NULL,
    rack INTEGER,
    shelf INTEGER
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR,
    password VARCHAR,
    name VARCHAR
);


CREATE TABLE IF NOT EXISTS admin_users (
    id SERIAL PRIMARY KEY,
    login VARCHAR,
    password VARCHAR
);
