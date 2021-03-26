CREATE TABLE categories (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE feeds_categories (
    id SERIAL NOT NULL PRIMARY KEY,
    feed_id INTEGER REFERENCES feeds(id),
    category_id INTEGER REFERENCES categories(id)
);
