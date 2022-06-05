alter table getme_db.public.offers
alter column status set default true;

alter table getme_db.public.plans
    alter column is_active set default true;