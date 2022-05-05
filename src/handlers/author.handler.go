package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/bariyale/src/db/gensql"
)

var AuthorMap = HandlerMap{
    "GET": {
        "/": AuthorHustler.GetAll,
        "/{id}": AuthorHustler.Get,
    },
    "POST": {
        "/": AuthorHustler.Post,
    },
    "PUT": {
        "/": AuthorHustler.Put,
    },
    "PATCH": {
        "/{id}": AuthorHustler.Patch,
    },
    "DELETE": {
        "/{id}": AuthorHustler.Delete,
    },
}
var AuthorHustler = Hustler{"author"};

func (h Hustler) Get(w http.ResponseWriter, r *http.Request) {
    queries := gensql.New(conn)
    result, err := queries.ListAuthors(handlerCtx)
    if err != nil {
        fmt.Println("abort", err)
    }
    json.NewEncoder(w).Encode(result)
}

func (h Hustler) Post(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hi")
}

func (h Hustler) Delete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hi")
}
func (h Hustler) Patch(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hi")
}
func (h Hustler) Put(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "hi")
}
func (h Hustler) GetAll(w http.ResponseWriter, r *http.Request) {
    queries := gensql.New(conn)
    result, err := queries.ListAuthors(handlerCtx)
    if err != nil {
        fmt.Println("abort", err)
    }
    json.NewEncoder(w).Encode(result)
}

