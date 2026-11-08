-- +goose Up
CREATE TABLE net_neighbors(
     id SERIAL PRIMARY KEY,
     node_uid TEXT NOT NULL,
     node_ip TEXT NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS messages;