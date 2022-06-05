-- POSTGRESQL
CREATE TABLE users_simple_auth
(
    id                 bigserial not null,
    login              text      not null unique,
    encrypted_password text      not null,
    user_id            bigint,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
-- MySQL
-- CREATE TABLE users_simple_auth
-- (
--     id                 bigint  not null auto_increment,
--     login              text(20),
--     encrypted_password text not null,
--     user_id            bigint,
--     PRIMARY KEY (id),
--     FOREIGN KEY (user_id) REFERENCES users (id)
-- );
