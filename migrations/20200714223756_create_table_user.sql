-- +goose Up
CREATE TABLE user (
username varchar(255) not null,
password text,
created_at text,
primary key (username)
);

-- +goose Down
DROP TABLE user;