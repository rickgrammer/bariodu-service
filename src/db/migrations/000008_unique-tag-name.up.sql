SET search_path TO papita;

ALTER TABLE tag
ADD CONSTRAINT tag_name_unique UNIQUE (name);
