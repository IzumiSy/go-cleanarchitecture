-- +migrate Up
create table todos (id string, name string, description string);

-- +migrate Down
drop table todos;
