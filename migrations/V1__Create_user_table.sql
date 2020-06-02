create table "User" (
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null
);