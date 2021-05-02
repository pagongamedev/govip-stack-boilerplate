-- +migrate Up
create table auth.user (
    -- id and uuid
    id serial not null,
    uuid uuid default uuid_generate_v4(),
    -- ------------------------------------------------
    -- uk
    email varchar,
    phone varchar,
    -- normal
    -- fk
    -- ------------------------------------------------
    -- time
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp null,
    -- ------------------------------------------------
    -- foreign
    -- constraint
    constraint user_id_pk primary key (id),
    constraint user_uuid_uk unique (uuid),
    constraint user_email_uk unique (email)
);

-- +migrate Down
drop table auth.user;