CREATE TABLE subcategories (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL
);

ALTER TABLE feeds_categories
ADD COLUMN subcategory_id INTEGER REFERENCES subcategories(id);