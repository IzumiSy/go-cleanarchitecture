-- +migrate Up
create table todos (id int, name string);

-- +migrate Down
drop table todos;
