alter table getme_db.public.offers
    alter column status drop default;

alter table getme_db.public.plans
    alter column is_active drop default;
