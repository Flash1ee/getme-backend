CREATE TABLE users_simple_auth
(
    id                 bigserial not null,
    login              text      not null unique,
    encrypted_password text      not null,
    user_id            bigint,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
