set search_path to papita;
ALTER TABLE author
RENAME COLUMN "firstName" to first_name;

ALTER TABLE author
RENAME COLUMN "lastName" to last_name;

ALTER TABLE author
RENAME COLUMN "lastLogin" to last_login;

ALTER TABLE profile
RENAME COLUMN "authorId" to author_id;

ALTER TABLE author_session
RENAME COLUMN "authorId" to author_id;

ALTER TABLE post
RENAME COLUMN "authorId" to author_id;

ALTER TABLE _post_tag
RENAME COLUMN "postId" to post_id;

ALTER TABLE _post_tag
RENAME COLUMN "tagId" to tag_id;

ALTER TABLE author
RENAME COLUMN isdeleted to is_deleted;

ALTER TABLE post
RENAME COLUMN isdeleted to is_deleted;

ALTER TABLE profile
RENAME COLUMN isdeleted to is_deleted;
