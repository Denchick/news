CREATE TABLE articles (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    link VARCHAR NOT NULL UNIQUE, 
    title VARCHAR NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
);