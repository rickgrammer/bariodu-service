package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type HandlerMap map[string]map[string]func(w http.ResponseWriter, r *http.Request)

type urlType struct {
    urlPrefix string;
    handlerMapper HandlerMap;
}

var urls = []urlType{
    {"/author", AuthorMap},
}

type Hustler struct {
    path string
}

type REST interface {
    Get(w http.ResponseWriter,r *http.Request)
    GetAll(w http.ResponseWriter,r *http.Request)
    Post(w http.ResponseWriter,r *http.Request)
    Put(w http.ResponseWriter,r *http.Request)
    Patch(w http.ResponseWriter,r *http.Request)
    Delete(w http.ResponseWriter,r *http.Request)
}

func BindRoutes(muxRouter *mux.Router) {
    for _, url_type := range urls {
        for http_method, sub_handlers := range url_type.handlerMapper {
            for sub_url, _handler := range sub_handlers {
                muxRouter.HandleFunc(url_type.urlPrefix+sub_url, _handler).Methods(http_method)
            }
        }
    }
}

var handlerCtx context.Context;
var conn *pgx.Conn;
var dbErr interface{};

func init() {
    handlerCtx = context.Background()
    conn, dbErr = pgx.Connect(handlerCtx, "postgresql://postgres:postgres@localhost:5500/bariodu?search_path=papita")
    if dbErr != nil {
        fmt.Printf("Db connection error\n%v", dbErr) 
    }
}
