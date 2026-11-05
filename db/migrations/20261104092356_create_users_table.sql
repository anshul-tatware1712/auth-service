-- +goose Up
CREATE TABLE users (
    id INT PRIMARY KEY,
    name TEXT,
    email TEXT,
    password TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;