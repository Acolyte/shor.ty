create table users
(
    id         bigserial not null
        constraint users_pkey
            primary key,
    name       varchar(128),
    username   varchar(255),
    password   varchar(128),
    created_at timestamp with time zone default(now()),
    updated_at timestamp with time zone default(now()),
    deleted_at timestamp with time zone default null
);

create index idx_users_deleted_at on users (deleted_at);
