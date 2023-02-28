CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE videos (
    id uuid default uuid_generate_v4() not null primary key,
    name text,
    description text,
    created_at timestamp without time zone not null,
    update_at timestamp without time zone not null
);
