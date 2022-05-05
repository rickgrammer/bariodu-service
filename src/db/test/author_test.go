package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	"example.com/bariyale/src/db/gensql"
	"example.com/bariyale/src/utils"
	"github.com/jackc/pgx/v4"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/require"
)

var testQueries *gensql.Queries;
var ctx context.Context;

func createAuthor(t *testing.T) gensql.CreateAuthorRow {
    faker := faker.New()
    person := faker.Person()
    author := gensql.CreateAuthorParams{
        Email: utils.RandomEmail(),
        Password: faker.Internet().Password(),
        FirstName: person.FirstName(),
        LastName: person.LastName(),
        LastLogin: sql.NullTime{
            Time: utils.RandTimestamp(),
            Valid: true,
        },
        Updated: time.Now(),
    }
    result, err := testQueries.CreateAuthor(ctx, author)
    require.NotEmpty(t, result)
    require.Nil(t, err)
    require.NotEmpty(t, result.Id)
    require.Equal(t, author.Password, result.Password)
    require.Equal(t, author.FirstName, result.FirstName)
    require.Equal(t, author.LastName, result.LastName)
    require.NotEmpty(t, result.LastLogin)
    require.NotEmpty(t, result.Updated)
    return result
}

func TestCreateAuthor(t *testing.T) {
    createAuthor(t)
}

func TestGetAuthor(t *testing.T) {
    author := createAuthor(t)
    result, err := testQueries.GetAuthor(ctx, author.Id)
    require.Nil(t, err)
    require.NotEmpty(t, result)
    require.Equal(t, author.Password, result.Password)
    require.Equal(t, author.FirstName, result.FirstName)
    require.Equal(t, author.LastName, result.LastName)
    require.NotEmpty(t, result.LastLogin)
    require.NotEmpty(t, result.Updated)
    require.WithinDuration(t, time.Now(), result.Created, time.Second)
}

func TestUpdateAuthor(t *testing.T) {
    author := createAuthor(t)
    f := gensql.UpdateAuthorParams{Id: author.Id, FirstName: "chiranjeevi"}
    r, err := testQueries.UpdateAuthor(ctx, f)
    require.Equal(t, f.FirstName, r.FirstName)
    require.Nil(t, err)
}

func TestMain(m *testing.M) {
    fmt.Println("ek bar")
	ctx = context.Background()
    conn, _ := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost:5500/bariodu?search_path=papita")
    defer conn.Close(ctx)
    testQueries = gensql.New(conn)
    os.Exit(m.Run())
}
