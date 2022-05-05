SET search_path TO papita;

ALTER TABLE profile
ALTER COLUMN "authorId" SET NOT NULL;

ALTER TABLE post
ALTER COLUMN "authorId" SET NOT NULL;
