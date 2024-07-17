package main

import (
	"fmt"
	logdy "github.com/logdyhq/logdy-core/logdy"
	"log"
	"net/http"
	"strings"
	"time"
)

type Logger struct {
	logdy   logdy.Logdy
	handler http.Handler
}

func (l Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	l.handler.ServeHTTP(w, r)
	if strings.HasPrefix(r.URL.Path, l.logdy.Config().HttpPathPrefix) {
		return
	}

	l.logdy.Log(logdy.Fields{
		"ua":     r.Header.Get("User-Agent"),
		"method": r.Method,
		"path":   r.URL.Path,
		"query":  r.URL.RawQuery,
		"time":   time.Since(start),
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	mux.HandleFunc("/v1/time", func(w http.ResponseWriter, r *http.Request) {
		curTime := time.Now().Format(time.Kitchen)
		w.Write([]byte(fmt.Sprintf("The Current Time is %v", curTime)))
	})

	logger := logdy.InitializeLogdy(logdy.Config{
		HttpPathPrefix: "/_logdy-ui",
	}, mux)

	addr := ":8002"
	log.Printf("server listen at %s", addr)
	log.Fatal(http.ListenAndServe(addr, &Logger{logdy: logger, handler: mux}))
}
