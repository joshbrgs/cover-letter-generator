package coverletter

import (
	"net/http"

	"github.com/go-chi/chi"
)

func CoverLetterRouter() http.Handler {
	r := chi.NewRouter()

	return r
}