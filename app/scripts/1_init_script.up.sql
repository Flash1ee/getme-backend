CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users
(
    nickname citext not null unique primary key,
    fullname text   not null,
    about    text,
    email    citext not null unique
);