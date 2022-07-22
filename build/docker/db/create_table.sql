drop table if exists users;

create table users (
    id int,
    name varchar(30),
    age int,
    password varchar(100),
    last_login timestamp with time zone,
    deleted boolean,
    create_user varchar(30),
    create_timestamp timestamp with time zone,
    update_user varchar(30),
    update_timestamp timestamp with time zone
);

insert into users values (
    1,
    'test',
    30,
    'test',
    current_timestamp,
    false,
    'a.oshima',
    current_timestamp,
    'a.oshima',
    current_timestamp
);

drop table if exists "sessions";

create table sessions (
    id varchar(64),
    user_id int,
    max_age int,
    deleted boolean,
    create_user varchar(30),
    create_timestamp timestamp with time zone,
    update_user varchar(30),
    update_timestamp timestamp with time zone
);

