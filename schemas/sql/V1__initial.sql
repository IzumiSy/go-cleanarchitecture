create table `todo` (
    `id` varchar(255) primary key,
    `user_id` varchar(255) not null,
    `name` varchar(255) not null,
    `description` varchar(255) not null
) engine=InnoDB default charset=utf8mb4;

create unique index idx_todo_name_user_id on todo(name, user_id);

create table `todo_history` (
  `id` varchar(255) primary key,
  `todo_id` varchar(255) not null,
  `history_type` varchar(255) not null,
  `previous` varchar(255),
  `current` varchar(255),
  `created_at` timestamp not null,

  constraint fk_todo_id foreign key (todo_id)
    references todo (id) on delete cascade
) engine=InnoDB default charset=utf8mb4;

create index idx_todo_history_todo_id on todo_history(todo_id);

create table `todo_category` (
  `name` varchar(255) not null,
  `user_id` varchar(255) not null,

  primary key (name, user_id)
) engine=InnoDB default charset=utf8mb4;

create table `todo_todo_category` (
  `todo_id` varchar(255) not null,
  `category_name` varchar(255) not null,

  constraint fk_category_name foreign key (category_name)
    references todo_category (name) on delete cascade on update cascade
) engine=InnoDB default charset=utf8mb4;

create index idx_todo_todo_category on todo_todo_category(todo_id);

create table `session` (
  `id` varchar(255) primary key,
  `user_id` varchar(255) not null,
  `created_at` timestamp not null
) engine=InnoDB default charset=utf8mb4;

create index idx_session_user_id on session(user_id);

create table `user` (
  `id` varchar(255) primary key,
  `name` varchar(255) not null
) engine=InnoDB default charset=utf8mb4;

create table `authentication` (
  `email` varchar(255) primary key,
  `user_id` varchar(255) not null,
  `hash` varchar(255) not null,
  `created_at` timestamp not null,

  constraint fk_authn_user_id foreign key (user_id)
    references user (id) on delete cascade
) engine=InnoDB default charset=utf8mb4;

create index idx_authn_user_id on authentication(user_id);
