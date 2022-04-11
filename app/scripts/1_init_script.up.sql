CREATE
EXTENSION IF NOT EXISTS citext;

CREATE TABLE users
(
    id            bigserial not null,
    name          text      not null,
    about         text      not null,
    avatar        text      not null,
    is_searchable bool      not null default false,
    PRIMARY KEY (id)
);

CREATE TABLE plans
(
    id        bigserial,
    name      text    not null,
    about     text,
    is_active boolean not null default false,
    progress  numeric          default 0,
    mentor_id bigint,
    mentee_id bigint,
    PRIMARY KEY (id),
    FOREIGN KEY (mentee_id) REFERENCES users (id)
);

CREATE TABLE color
(
    name  citext,
    value int,
    PRIMARY KEY (name)
);

CREATE TABLE status
(
    name  citext,
    color citext,
    PRIMARY KEY (name),
    FOREIGN KEY (color) REFERENCES color (name)
);

CREATE TABLE task
(
    id          bigserial,
    name        citext not null,
    description text   not null,
    deadline    timestamptz default now()::timestamptz not null,
    status      text   not null,
    plan_id     bigint,
    PRIMARY KEY (id),
    FOREIGN KEY (plan_id) REFERENCES plans (id),
    FOREIGN KEY (status) REFERENCES status (name)
);

CREATE TABLE skills
(
    name  citext not null,
    color citext not null,
    PRIMARY KEY (name),
    FOREIGN KEY (color) REFERENCES color (name)
);

CREATE TABLE users_skills
(
    id         bigserial not null,
    user_id    bigint    not null,
    skill_name citext    not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (skill_name) REFERENCES skills (name)
);

CREATE TABLE offers
(
    id         bigserial not null,
    skill_name citext    not null,
    status     boolean   not null default false,
    mentor_id  bigint    not null,
    mentee_id  bigint    not null,
    PRIMARY KEY (id),
    FOREIGN KEY (skill_name) REFERENCES skills (name),
    FOREIGN KEY (mentor_id) REFERENCES users (id),
    FOREIGN KEY (mentee_id) REFERENCES users (id)
);

