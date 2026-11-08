-- +goose Up
CREATE TABLE users (
    uid TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    avatar_path TEXT
);

-- +goose Down
DROP TABLE IF EXISTS users;