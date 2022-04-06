-- +goose Up
CREATE TABLE IF NOT EXISTS visits (
   id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
   hits integer
);

INSERT INTO visits (hits) VALUES (0);

-- +goose Down
DROP TABLE visits;