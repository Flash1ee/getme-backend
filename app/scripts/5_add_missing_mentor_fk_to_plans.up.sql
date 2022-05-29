alter table getme_db.public.plans add
    constraint fk_mentor_id foreign key (mentor_id) references getme_db.public.users(id);