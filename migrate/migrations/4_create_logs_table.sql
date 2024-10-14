create table logs(
    id serial primary key,
    user_id serial references users(id),
    media_id serial references media(id),
    rating int,
    started_on timestamp,
    finished_on timestamp,
    created_on timestamp,
    updated_on timestamp
);