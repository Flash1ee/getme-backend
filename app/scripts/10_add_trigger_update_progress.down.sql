drop trigger update_progress on task;

drop trigger insert_progress on task;

drop function update_progress();

alter table plans
ALTER progress SET default 0;

alter table task
    ALTER status DROP default;