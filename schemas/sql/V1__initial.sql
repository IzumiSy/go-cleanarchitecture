create table todo (
    id varchar(255) primary key,
    user_id varchar(255) not null,
    name varchar(255) not null,
    description varchar(255) not null
);
