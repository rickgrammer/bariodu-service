package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"time"

	"example.com/bariyale/src/handlers"
	"github.com/gorilla/mux"
)


func forceJsonMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/json")
        next.ServeHTTP(w, r)
    })
}


var middlewares = []func(next http.Handler) http.Handler{
    forceJsonMiddleware, 
}
var addr = "127.0.0.1:8088"

func main() {
    muxRouter := mux.NewRouter()
    for _, middleware := range middlewares {
        muxRouter.Use(middleware)
    }

    handlers.BindRoutes(muxRouter)
    ctx, cancelCtx := context.WithCancel(context.Background())
    server := &http.Server{
        Handler: muxRouter,
        Addr: addr,
        WriteTimeout: 15*time.Second,
        ReadTimeout: 15*time.Second,
        BaseContext: func(l net.Listener) context.Context {
            ctx = context.WithValue(ctx, "aawargi", "saath mere")
            return ctx
        },
    }
    // Run our server in a goroutine so that it doesn't block.
    go func() {
        fmt.Printf("Listening on %s\n", addr)
        if err := server.ListenAndServe(); err != nil {
            // log.Println(err)
            if errors.Is(err, http.ErrServerClosed) {
                fmt.Println("server closed")
            } else if err != nil {
                fmt.Printf("errors starting server: %s\n", err)
            }
            cancelCtx()
        }
    }()
    
    <- ctx.Done()

}

