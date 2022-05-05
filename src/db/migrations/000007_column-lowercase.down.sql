set search_path to papita;
ALTER TABLE author
RENAME COLUMN first_name to "firstName";

ALTER TABLE author
RENAME COLUMN last_name to "lastName";

ALTER TABLE author
RENAME COLUMN last_login to "lastLogin";

ALTER TABLE profile
RENAME COLUMN author_id to "authorId";

ALTER TABLE author_session
RENAME COLUMN author_id to "authorId";

ALTER TABLE post
RENAME COLUMN author_id to "authorId";

ALTER TABLE _post_tag
RENAME COLUMN post_id to "postId";

ALTER TABLE _post_tag
RENAME COLUMN tag_id to "tagId";

ALTER TABLE author
RENAME COLUMN is_deleted to isdeleted;

ALTER TABLE post
RENAME COLUMN is_deleted to isdeleted;

ALTER TABLE profile
RENAME COLUMN is_deleted to isdeleted;
