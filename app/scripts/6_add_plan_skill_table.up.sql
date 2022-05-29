create table plans_skills (
                              id bigserial not null,
                              plan_id bigint not null,
                              skill_name citext,
                              PRIMARY KEY (id),
                              FOREIGN KEY (plan_id) REFERENCES plans (id),
                              FOREIGN KEY (skill_name) REFERENCES skills (name)
);