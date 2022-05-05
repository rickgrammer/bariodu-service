package db

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"example.com/bariyale/src/db/gensql"
	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/require"
)


func createPost(author *gensql.CreateAuthorRow, t *testing.T) gensql.CreatePostRow {
    type AnyJson map[string]interface{};
    jsonb := AnyJson{
            "name":        "Passata",
            "ingredients": []string{"Tomatoes", "Onion", "Olive oil", "Garlic"},
            "organic":     true,
            "dimensions": map[string]interface{}{
                "weight": 250.00,
            },
        }
    a, _ := json.Marshal(jsonb)
    b := pgtype.JSONB{
        Bytes: a,
        Status: pgtype.Present,
    }
    post := gensql.CreatePostParams{
        AuthorID: author.Id,
        Content: b,
        Updated: time.Now(),
    }
    result, err := testQueries.CreatePost(ctx, post)
    require.NotEmpty(t, result)
    require.Nil(t, err)
    require.NotEmpty(t, result.Id)
    require.Equal(t, post.AuthorID, result.AuthorID)
    require.NotEmpty(t, post.Content)
    require.NotEmpty(t, result.Updated)
    return result
}

func TestCreatePost(t *testing.T) {
    author := createAuthor(t)
    createPost(&author, t)
}


func TestGetPost(t *testing.T) {
    author := createAuthor(t)

    post := createPost(&author, t)
    
    result, err := testQueries.GetPost(ctx, gensql.GetPostParams{AuthorID: author.Id, Id: post.Id})
    require.Nil(t, err)
    require.NotEmpty(t, result)
    require.Equal(t, post.AuthorID, result.AuthorID)
    require.NotEmpty(t, result.Content)
    require.NotEmpty(t, result.Updated)
    require.WithinDuration(t, time.Now(), result.Created, time.Second)
}

func TestPost(t *testing.T) {
    r, _ := testQueries.ListPosts(ctx, gensql.ListPostsParams{Offset: 0, Limit: 10}) 
    fmt.Printf("%v", r)
    fmt.Printf("tapka")
}
