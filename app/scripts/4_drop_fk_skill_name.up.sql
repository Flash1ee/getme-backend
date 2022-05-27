delete from getme_db.public.skills where name = '';

ALTER TABLE getme_db.public.offers
    alter  skill_name drop default;
