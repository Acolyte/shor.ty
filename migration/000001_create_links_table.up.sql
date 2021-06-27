create table links
(
    id         bigserial not null
        constraint links_pkey
            primary key,
    uuid       text,
    full_url   varchar(2048),
    scheme     varchar(8),
    host       varchar(253),
    port       smallint,
    path       varchar(2048),
    query      varchar(2048),
    expires_in varchar(32),
    expires_at timestamp with time zone default null,
    created_at timestamp with time zone default (now()),
    updated_at timestamp with time zone default (now()),
    deleted_at timestamp with time zone
);

alter table links
    owner to db;

create index idx_links_deleted_at on links (deleted_at);

create index idx_query on links (query);

create index idx_path on links (path);

create index idx_port on links (port);

