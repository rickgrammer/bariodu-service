SET search_path TO papita;

ALTER TABLE author
ALTER COLUMN "firstName" SET NOT NULL,
ALTER COLUMN "lastName" SET NOT NULL;
