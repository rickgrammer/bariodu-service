SET search_path TO papita;
ALTER TABLE author
DROP COLUMN IF EXISTS "isDeleted";
ALTER TABLE post
DROP COLUMN IF EXISTS "isDeleted";
ALTER TABLE profile
DROP COLUMN IF EXISTS "isDeleted";
