create table users(
    id serial primary key,
    email varchar(150),
    name varchar(100),
    display_name varchar(100),
    password varchar(150),
    created_on timestamp,
    updated_on timestamp
);