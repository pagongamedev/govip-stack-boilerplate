-- +migrate Up
create table auth.password (
    -- id and uuid
    id serial not null,
    uuid uuid default uuid_generate_v4(),
    -- ------------------------------------------------
    -- uk
    hashsalt varchar,
    -- normal
    -- fk
    user_id int,
    -- ------------------------------------------------
    -- time
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp null,
    -- ------------------------------------------------
    -- foreign
    foreign key (user_id) references auth.user (id),
    -- constraint
    constraint password_id_pk primary key (id),
    constraint password_uuid_uk unique (uuid)
);

-- +migrate Down
drop table auth.password;