-- Администратор
CREATE ROLE administrators;
GRANT USAGE ON SCHEMA public TO administrators;
GRANT SELECT, INSERT, UPDATE, DELETE
    ON ALL TABLES
    IN SCHEMA public
    TO administrators;

ALTER DEFAULT PRIVILEGES IN SCHEMA public
    GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLES TO administrators;

CREATE USER admin
    WITH
    CREATEDB
    CREATEROLE
    ENCRYPTED PASSWORD 'qwerty'
    IN ROLE administrators;
-- Гость
CREATE ROLE guest;

CREATE USER visitor
    WITH ENCRYPTED PASSWORD 'qwerty'
    IN ROLE guest;

GRANT SELECT
    ON TABLE users, users_simple_auth, users_skills, telegram_auth,task,status,skills,plans_skills,plans,offers,color
    TO visitor;

-- Пользователь
CREATE ROLE users;

CREATE USER getme
    WITH ENCRYPTED PASSWORD 'getme-app'
    IN ROLE users;

GRANT SELECT, INSERT, UPDATE, DELETE
    ON TABLE users, users_simple_auth, users_skills, telegram_auth,task,status,skills,plans_skills,plans,offers,color
    TO getme;

