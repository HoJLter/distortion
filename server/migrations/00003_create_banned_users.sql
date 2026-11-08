-- +goose Up
CREATE TABLE banned_users(
    user_uid TEXT PRIMARY KEY,
    reason TEXT
);
ALTER TABLE banned_users
ADD CONSTRAINT fk_user_uid
FOREIGN KEY (user_uid)
REFERENCES users(uid);

-- +goose Down
DROP TABLE IF EXISTS banned_users;