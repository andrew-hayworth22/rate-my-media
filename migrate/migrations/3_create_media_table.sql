create table media_type(
    id serial primary key,
    name varchar(20)
);

insert into media_type(name) values ('Movie'), ('Video Game'), ('Book'), ('TV Show');

create table media(
    id serial primary key,
    media_type_id serial references media_type(id),
    name varchar(300),
    description varchar(1000),
    release_date date
);