package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joshbrgs/cvg/cmds/coverletter"
	"github.com/joshbrgs/cvg/pkg/logger"
	"github.com/joshbrgs/cvg/pkg/server"
)

func main() {
	log := logger.Init()

    r := chi.NewRouter()

    r.Mount("/cv", coverletter.CoverLetterRouter())

    s := server.NewServer(
        server.WithHost("localhost"),
    )

    log.Infof("Server Starting on http://%s:%s", s.Host, s.Port )

    if err := http.ListenAndServe(":" + s.Port, r); err != nil {
        log.Fatal(err)
    }
}
