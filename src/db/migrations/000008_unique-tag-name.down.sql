SET search_path TO papita;

ALTER TABLE tag
DROP constraint tag_name_unique;
