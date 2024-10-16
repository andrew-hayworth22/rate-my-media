create table logs(
    id serial primary key,
    user_id serial references users(id),
    media_id serial references media(id),
    rating int check (rating > 0 and rating <= 100),
    started_on timestamp,
    finished_on timestamp,
    created_on timestamp,
    updated_on timestamp
);