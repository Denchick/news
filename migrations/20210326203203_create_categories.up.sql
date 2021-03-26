CREATE TABLE categories (
    id SERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE feeds_categories (
    id SERIAL NOT NULL PRIMARY KEY,
    feed_id SERIAL REFERENCES feeds(id),
    category_id SERIAL REFERENCES categories(id)
);
