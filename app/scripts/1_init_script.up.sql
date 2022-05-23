# PostgreSQL
# CREATE
#     EXTENSION IF NOT EXISTS citext;
#
# CREATE TABLE users
# (
#     id            bigserial                              not null,
#     tg_id         bigint                                 not null unique,
#     first_name    citext                                 not null,
#     last_name     citext                                 not null,
#     nickname      citext                                 not null,
#     about         text        default '',
#     avatar        text                                   not null,
#     is_searchable bool                                   not null default false,
#     created_at    timestamptz default now()::timestamptz not null,
#     updated_at    timestamptz default now()::timestamptz not null,
#
#     PRIMARY KEY (id)
# );
#
# CREATE TABLE plans
# (
#     id         bigserial,
#     name       text    not null,
#     about      text,
#     is_active  boolean not null default false,
#     progress   numeric          default 0,
#     mentor_id  bigint,
#     mentee_id  bigint,
#     created_at timestamptz      default now()::timestamptz not null,
#     PRIMARY KEY (id),
#     FOREIGN KEY (mentee_id) REFERENCES users (id)
# );
#
# CREATE TABLE color
# (
#     name  citext,
#     value int,
#     PRIMARY KEY (name)
# );
#
# CREATE TABLE status
# (
#     name  citext,
#     color citext,
#     PRIMARY KEY (name),
#     FOREIGN KEY (color) REFERENCES color (name)
# );
#
# CREATE TABLE task
# (
#     id          bigserial,
#     name        citext                                 not null,
#     description text                                   not null,
#     deadline    timestamptz default now()::timestamptz not null,
#     status      citext                                 not null,
#     plan_id     bigint,
#     created_at  timestamptz default now()::timestamptz not null,
#     PRIMARY KEY (id),
#     FOREIGN KEY (plan_id) REFERENCES plans (id),
#     FOREIGN KEY (status) REFERENCES status (name)
# );
#
# CREATE TABLE skills
# (
#     name  citext not null,
#     color citext not null,
#     PRIMARY KEY (name),
#     FOREIGN KEY (color) REFERENCES color (name)
# );
#
# CREATE TABLE users_skills
# (
#     id         bigserial not null,
#     user_id    bigint    not null,
#     skill_name citext    not null,
#     PRIMARY KEY (id),
#     FOREIGN KEY (user_id) REFERENCES users (id),
#     FOREIGN KEY (skill_name) REFERENCES skills (name)
# );
#
# CREATE TABLE offers
# (
#     id         bigserial not null,
#     skill_name citext    not null,
#     status     boolean   not null default false,
#     mentor_id  bigint    not null,
#     mentee_id  bigint    not null,
#     created_at timestamptz        default now()::timestamptz not null,
#     PRIMARY KEY (id),
#     FOREIGN KEY (skill_name) REFERENCES skills (name),
#     FOREIGN KEY (mentor_id) REFERENCES users (id),
#     FOREIGN KEY (mentee_id) REFERENCES users (id)
# );

# MYSQL

CREATE TABLE users
(
    id            bigint not null auto_increment,
    tg_id         bigint not null unique,
    first_name    text   not null,
    last_name     text   not null,
    nickname      text   not null,
    about         text,
    avatar        text   not null,
    is_searchable bool   not null default false,
    created_at    time,
    updated_at    time,

    PRIMARY KEY (id)
);

CREATE TABLE plans
(
    id         bigint auto_increment,
    name       text(20) not null,
    about      text(30),
    is_active  boolean  not null default false,
    progress   numeric           default 0,
    mentor_id  bigint,
    mentee_id  bigint,
    created_at time,
    PRIMARY KEY (id),
    FOREIGN KEY (mentee_id) REFERENCES users (id)
);

CREATE TABLE color
(
    name  text(3),
    value bigint,
    PRIMARY KEY (name(2))
);

CREATE TABLE status
(
    name  text(20),
    color text(30),
    PRIMARY KEY (name(2))
);

CREATE TABLE task
(
    id          bigint auto_increment,
    name        text(30),
    description text(50) not null,
    deadline    time,
    st          text(20) not null,
    plan_id     bigint,
    created_at  time,
    PRIMARY KEY (id)
);

CREATE TABLE skills
(
    name  text(25) not null,
    color text(2)  not null,
    PRIMARY KEY (name(10))
);

CREATE TABLE users_skills
(
    id         bigint   not null auto_increment,
    user_id    bigint   not null,
    skill_name text(25) not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE offers
(
    id         bigint   not null auto_increment,
    skill_name text(25) not null,
    status     boolean  not null default false,
    mentor_id  bigint   not null,
    mentee_id  bigint   not null,
    created_at time,
    PRIMARY KEY (id),
    FOREIGN KEY (mentor_id) REFERENCES users (id),
    FOREIGN KEY (mentee_id) REFERENCES users (id)
);

