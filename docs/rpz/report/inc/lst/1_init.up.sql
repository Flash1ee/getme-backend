CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users
(
    id bigserial not null,
    first_name citext,
    last_name citext,
    nickname citext not null unique,
    about         text          default '',
    avatar        text,
    is_searchable bool not null default false,
    created_at timestamptz default now()::timestamptz not null,
    updated_at timestamptz default now()::timestamptz not null,
    PRIMARY KEY (id)
);
CREATE TABLE telegram_auth
(
    tg_id   bigint not null,
    created_at timestamptz default now()::timestamptz not null,
    last_auth timestamptz default now()::timestamptz not null,
    user_id bigint,
    PRIMARY KEY (tg_id),
    FOREIGN KEY (user_id) references users (id)
);

CREATE TABLE plans
(
    id bigserial,
    name      text    not null,
    about     text,
    is_active boolean not null default false,
    progress  numeric          default 0,
    mentor_id bigint,
    mentee_id bigint,
    created_at timestamptz default now()::timestamptz not null,
    PRIMARY KEY (id),
    FOREIGN KEY (mentee_id) REFERENCES users (id)
);

CREATE TABLE color
(
    name citext,
    value int,
    PRIMARY KEY (name)
);

CREATE TABLE status
(
    name citext,
    color citext,
    PRIMARY KEY (name),
    FOREIGN KEY (color) REFERENCES color (name)
);

CREATE TABLE task
(
    id bigserial,
    name citext not null,
    description text not null,
    deadline timestamptz default now()::timestamptz not null,
    status citext not null,
    plan_id     bigint,
    created_at timestamptz default now()::timestamptz not null,
    PRIMARY KEY (id),
    FOREIGN KEY (plan_id) REFERENCES plans (id),
    FOREIGN KEY (status) REFERENCES status (name)
);

CREATE TABLE skills
(
    name citext not null,
    color citext,
    PRIMARY KEY (name),
    FOREIGN KEY (color) REFERENCES color (name)
);

CREATE TABLE users_skills
(
    id bigserial not null,
    user_id bigint not null,
    skill_name citext,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (skill_name) REFERENCES skills (name)
);

CREATE TABLE offers
(
    id bigserial not null,
    skill_name citext not null,
    status    boolean not null default false,
    mentor_id bigint  not null,
    mentee_id bigint  not null,
    created_at timestamptz default now()::timestamptz not null,
    PRIMARY KEY (id),
    FOREIGN KEY (skill_name) REFERENCES skills (name),
    FOREIGN KEY (mentor_id) REFERENCES users (id),
    FOREIGN KEY (mentee_id) REFERENCES users (id)
);