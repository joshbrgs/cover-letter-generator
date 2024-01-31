package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/docgen"
	"github.com/go-chi/render"
	"github.com/joshbrgs/cvg/cmds/coverletter"
	"github.com/joshbrgs/cvg/pkg/logger"
	"github.com/joshbrgs/cvg/pkg/server"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	log := logger.Init()
    flag.Parse()

    r := chi.NewRouter()

      // A good base middleware stack
  r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	})

    r.Mount("/cv", coverletter.CoverLetterRouter())

    if *routes {
		// fmt.Println(docgen.JSONRoutesDoc(r))
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/joshbrgs/cover-letter-generator",
			Intro:       "Welcome to the API CVG Chi generated docs.",
		}))
		return
	}

    s := server.NewServer(
        server.WithHost("localhost"),
    )

    log.Infof("Server Starting on http://%s:%s", s.Host, s.Port )

    if err := http.ListenAndServe(":" + s.Port, r); err != nil {
        log.Fatal(err)
    }
}
