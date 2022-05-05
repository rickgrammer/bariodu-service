SET search_path TO papita;
ALTER TABLE author
DROP COLUMN IF EXISTS "firstName",
DROP COLUMN IF EXISTS "lastName";
