package db

import (
	"time"

	"testing"

	"example.com/bariyale/src/db/gensql"
	"example.com/bariyale/src/utils"
	"github.com/stretchr/testify/require"
)

func createTag(post *gensql.CreatePostRow, t *testing.T) gensql.Tag {
    tag := gensql.CreateTagParams{
        Name: utils.RandomString(7),
        Updated: time.Now(),
    }
    result, err := testQueries.CreateTag(ctx, tag)
    require.Nil(t, err)
    require.NotEmpty(t, result)
    require.Equal(t, tag.Name, result.Name)
    require.NotEmpty(t, tag.Updated)
    return result
}

func TestListTags(t *testing.T) {
    author := createAuthor(t)
    post := createPost(&author, t)
    createTag(&post, t)

    result, err := testQueries.ListTags(ctx, gensql.ListTagsParams{Offset: 0, Limit: 10})
    require.Nil(t, err)
    require.NotEmpty(t, result)
}

func TestListTagsForPost(t *testing.T) {
    author := createAuthor(t)
    post := createPost(&author, t)
    tag := createTag(&post, t)
    createTagForPost(&post, &tag, t)
    result, err := testQueries.ListTagsForPost(ctx, post.Id)
    require.Nil(t, err)
    require.NotEmpty(t, result)

}

func createTagForPost(post *gensql.CreatePostRow, tag *gensql.Tag, t *testing.T) gensql.PostTag {
    result, err := testQueries.CreateTagForPost(ctx, gensql.CreateTagForPostParams{PostID: post.Id, TagID: tag.Id})
    require.Nil(t, err)
    require.NotEmpty(t, result)
    return result
}

func TestCreateTagForPost(t *testing.T) {
    author := createAuthor(t)
    post := createPost(&author, t)
    tag := createTag(&post, t)
    result := createTagForPost(&post, &tag, t)
    require.NotEmpty(t, result)
}

func TestCreateTag(t *testing.T) {
    author := createAuthor(t)
    post := createPost(&author, t)
    createTag(&post, t)
}
