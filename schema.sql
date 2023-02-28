CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE videos (
    id uuid default uuid_generate_v4() not null primary key,
    name text not null,
    description text not null,
    url text not null,
    duration bigint not null,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);
