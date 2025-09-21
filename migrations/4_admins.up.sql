CREATE TABLE IF NOT EXISTS admins (
    id       SERIAL  PRIMARY KEY,
    login    VARCHAR UNIQUE,
    password VARCHAR
);
