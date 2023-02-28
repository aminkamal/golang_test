CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE videos (
    id uuid default uuid_generate_v4() not null primary key,
    name text not null,
    description text not null,
    url text not null,
    duration integer not null,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);

CREATE TABLE annotations (
    id uuid default uuid_generate_v4() not null primary key,
    video_id uuid references videos (id) on delete cascade not null,
    "type" text not null,
    note text not null,
    start_marker integer not null,
    end_marker integer not null,
    created_at timestamp without time zone not null,
    updated_at timestamp without time zone not null
);

CREATE INDEX idx_annotation_video_id ON annotations(video_id);
