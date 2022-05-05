SET search_path TO papita;
ALTER TABLE author
ADD COLUMN isDeleted boolean default false;

ALTER TABLE profile
ADD COLUMN isDeleted boolean default false;

ALTER TABLE post
ADD COLUMN isDeleted boolean default false;
