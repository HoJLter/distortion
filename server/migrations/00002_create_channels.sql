-- +goose Up
CREATE TABLE channels(
    uid TEXT PRIMARY KEY,
    type TEXT NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS channels;