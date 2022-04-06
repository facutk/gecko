-- +goose Up
CREATE TABLE IF NOT EXISTS user (
   id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
   hits integer
);

-- +goose Down
DROP TABLE user;