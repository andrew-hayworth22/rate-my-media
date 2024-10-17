create table media_types(
    id int primary key,
    name varchar(20)
);

insert into media_types(id, name) values (1, 'Movie'), (2, 'Video Game'), (3, 'Book'), (4, 'TV Show'), (5, 'Album');

create table media(
    id serial primary key,
    media_type_id serial references media_types(id),
    name varchar(300),
    description varchar(1000),
    release_date date
);

create table movies(
    id int primary key references media(id) on delete cascade,
    runtime_minutes int
);

create table video_games(
    id int primary key references media(id) on delete cascade
);

create table books(
    id int primary key references media(id) on delete cascade,
    pages int
);

create table tv_shows(
    id int primary key references media(id) on delete cascade,
    episode_runtime_minutes int
);