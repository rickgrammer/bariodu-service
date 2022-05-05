SET search_path TO papita;
ALTER TABLE author
ADD COLUMN "firstName" varchar(30),
ADD COLUMN "lastName" varchar(30);
