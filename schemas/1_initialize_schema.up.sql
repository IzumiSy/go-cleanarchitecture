-- +migrate Up
create table todo (
    id string,
    name string,
    description string,
    user_id string,

    primary key (id)
);

create table authentication (
    email string,
    hash string,
    user_id string,
    created_at timestamp,

    foreign key (user_id) references user(id),
    primary key (email)
);

create table session (
    id string,
    user_id string,
    created_at timestamp,

    foreign key (user_id) references user(id),
    primary key (id)
);

create table user (
    id string,
    name string,

    primary key (id)
);

-- +migrate Down
drop table todo;
drop table authentication;
drop table user;
