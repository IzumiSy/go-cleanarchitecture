-- +migrate Up
create table todo (
    id string,
    name string,
    description string,
    user_id string
);

-- +migrate Down
drop table todo;
