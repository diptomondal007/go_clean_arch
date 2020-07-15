-- +goose Up
CREATE TABLE user (
id int not null,
username text,
password text,
created_at text,
primary key (id)
);

-- +goose Down
DROP TABLE user;