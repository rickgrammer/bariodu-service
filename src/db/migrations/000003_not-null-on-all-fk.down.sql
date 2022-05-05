SET search_path TO papita;

ALTER TABLE profile
ALTER COLUMN "authorId" DROP NOT NULL;

ALTER TABLE post
ALTER COLUMN "authorId" DROP NOT NULL;
