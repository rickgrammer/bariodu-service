CREATE SCHEMA IF NOT EXISTS papita;
-- Use papita as default schema, reset again to use public schema
SET search_path TO papita;

SET TIMEZONE TO 'Asia/Kolkata';

CREATE TABLE IF NOT EXISTS author (
    id uuid primary key,
    email varchar(30) unique,
    password varchar(30) not null,
    created timestamptz(3) DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz(3) not null,
    "lastLogin" timestamptz
);

CREATE TABLE IF NOT EXISTS profile (
    id uuid primary key,
    "authorId" uuid,
    image text,
    about jsonb,
    created timestamptz(3) DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz(3) not null
);

CREATE TABLE IF NOT EXISTS author_session (
    id text primary key,
    "authorId" uuid not null,
    created timestamptz(3) DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS post (
    id uuid primary key,
    "authorId" uuid,
    content jsonb,
    created timestamptz(3) DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz(3) not null
);

CREATE TABLE IF NOT EXISTS tag (
    id uuid primary key,
    name varchar(30) not null,
    created timestamptz(3) DEFAULT CURRENT_TIMESTAMP,
    updated timestamptz(3) not null
);

CREATE TABLE IF NOT EXISTS _post_tag (
    "postId" uuid not null,
    "tagId" uuid not null
);

ALTER TABLE _post_tag
ADD PRIMARY KEY ("postId", "tagId");

ALTER TABLE _post_tag
ADD FOREIGN KEY ("postId")
REFERENCES post(id)
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE _post_tag
ADD FOREIGN KEY ("tagId")
REFERENCES tag(id)
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE profile
ADD FOREIGN KEY ("authorId") 
REFERENCES author(id)
ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE author_session
ADD FOREIGN KEY ("authorId") 
REFERENCES author(id)
ON DELETE CASCADE ON UPDATE SET NULL;

ALTER TABLE post
ADD FOREIGN KEY ("authorId") 
REFERENCES author(id)
ON DELETE CASCADE ON UPDATE CASCADE;
