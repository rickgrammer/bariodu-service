package handlers

import (
	// "encoding/json"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/bariyale/src/db/gensql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type MinimalCRUDer interface {
    Get(w http.ResponseWriter,r *http.Request)
    Post(w http.ResponseWriter,r *http.Request)
    Delete(w http.ResponseWriter,r *http.Request)
}

type CRUDer interface {
    Get(w http.ResponseWriter,r *http.Request)
    GetAll(w http.ResponseWriter,r *http.Request)
    Post(w http.ResponseWriter,r *http.Request)
    Put(w http.ResponseWriter,r *http.Request)
    Patch(w http.ResponseWriter,r *http.Request)
    Delete(w http.ResponseWriter,r *http.Request)
}


type urlType struct {
    url string;
    method string;
    handler func(w http.ResponseWriter,r *http.Request);
}

var urls = []urlType{
    {"/",      "GET", bazi},
    {"/posts", "GET", posts},
    {"/posts/{id}", "GET", getPost},
    {"/bro",   "GET", hard},
}

func getPost(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    conn, _ := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost:5500/bariodu?search_path=papita")
    queries := gensql.New(conn)
    defer conn.Close(ctx)
    params := mux.Vars(r)
    id := params["id"]
    _id, err := uuid.Parse(id)
    if err != nil {
        http.Error(w, "Invalid id", 401)
        return
    }
    u, _ := uuid.Parse("6360214e-ceeb-426b-88de-95518d4a0e5f")
    result, err := queries.GetPost(ctx, gensql.GetPostParams{AuthorID: u, Id: _id})
    json.NewEncoder(w).Encode(result)
}

func posts(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    conn, _ := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost:5500/bariodu?search_path=papita")
    queries := gensql.New(conn)
    defer conn.Close(ctx)
    
    result, err := queries.ListPosts(ctx, gensql.ListPostsParams{AuthorID: uuid.MustParse("6360214e-ceeb-426b-88de-95518d4a0e5f"), Offset: 0, Limit: 10})
    // result, err := queries.ListPosts(ctx, gensql.ListPostsParams{AuthorID: "b2aaa057-61ad-44bb-be87-49c43d4687b8"})
    if err != nil {
        fmt.Printf("%v",err)
    }
    if len(result) == 0 {
        json.NewEncoder(w).Encode(make([]interface{}, 0))
    } else {
        json.NewEncoder(w).Encode(result)
    }
}

func bazi(w http.ResponseWriter, r *http.Request) {
    ctx := context.Background()
    conn, _ := pgx.Connect(ctx, "postgresql://postgres:postgres@localhost:5500/bariodu?search_path=papita")
    queries := gensql.New(conn)
    defer conn.Close(ctx)
    result, err := queries.ListAuthors(ctx)
    if err != nil {
        fmt.Printf("%v",err)
    }
    json.NewEncoder(w).Encode(result)
}

func hard(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func BindRoutes(muxRouter *mux.Router) {
    for _, url := range urls {
        muxRouter.HandleFunc(url.url, url.handler).Methods(url.method)
    }
}
