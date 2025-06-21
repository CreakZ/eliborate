CREATE TABLE IF NOT EXISTS books (
    id          SERIAL    PRIMARY KEY,
    title       VARCHAR   NOT NULL,
    description TEXT,
    category_id INTEGER,
    authors     VARCHAR[] NOT NULL,
    cover_urls  VARCHAR[] NULL,
    rack        INTEGER,
    shelf       INTEGER,

	FOREIGN KEY (category_id) REFERENCES categories(id)
);
