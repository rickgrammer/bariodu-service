SET search_path TO papita;

ALTER TABLE author
ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE profile
ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE post
ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE tag
ALTER COLUMN id SET DEFAULT gen_random_uuid();
