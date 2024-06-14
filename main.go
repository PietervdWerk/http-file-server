package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	mux := http.NewServeMux()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	mux.Handle("/*", loggerMiddleware(http.FileServer(http.Dir(wd))))

	log.Info().Msg("starting server on :8888")
	if err := http.ListenAndServe(":8888", mux); err != nil {
		log.Fatal().Err(err).Msg("http server failed")
	}
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("method", r.Method).Str("path", r.URL.Path).Msg("request received")
		next.ServeHTTP(w, r)
	})
}
