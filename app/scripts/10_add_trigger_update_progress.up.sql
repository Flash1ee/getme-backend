create or replace function update_progress() returns trigger as
$psql$
begin
    UPDATE plans
    SET progress =
            COALESCE(
                        ((select count(*) from task where plan_id = new.plan_id and status = 'Выполнена')::float *
                         100::float) /
                        NULLIF(
                                (select count(*)
                                 from task
                                 where plan_id = new.plan_id),
                                0
                            )::float,
                        100
                )::numeric
    WHERE id = new.plan_id;
    return new;
end;
$psql$ language plpgsql;

create trigger update_progress
    after update
    on task
    for each row
execute procedure update_progress();

create trigger insert_progress
    after insert
    on task
    for each row
execute procedure update_progress();

alter table plans
    ALTER progress SET default 100;

alter table task
    ALTER status SET default 'В процессе';