-- +goose Up
CREATE TABLE messages(
    id SERIAL PRIMARY KEY,
    channel_uid TEXT NOT NULL,
    user_uid TEXT NOT NULL,
    content TEXT NOT NULL
);
ALTER TABLE messages
ADD CONSTRAINT fk_channel_uid
FOREIGN KEY (channel_uid)
REFERENCES channels(uid);

-- +goose Down
DROP TABLE IF EXISTS messages;