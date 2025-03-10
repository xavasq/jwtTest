-- +goose Up
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(70) NOT NULL
);
-- +goose Down
DROP TABLE users;


