package coverletter

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type CoverLetter struct {
	ID string
	Date time.Duration
	Company string
	Body string
	Variables map[string]string
	Hex string
}

func CoverLetterRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		// r.With(middleware.Paginate).Get("/", ListCoverLetters)
		r.Post("/", CreateCoverLetter)

		// r.Route("/{cvID}", func(r chi.Router) {
		// 	r.Use(CoverLetterCtx)            // Load the *CoverLetter on the request context
		// 	r.Get("/", GetCoverLetter)       // GET /CoverLetter/123
		// 	r.Put("/", UpdateCoverLetter)    // PUT /CoverLetter/123
		// 	r.Delete("/", DeleteCoverLetter) // DELETE /CoverLetter/123
		// })
	})

	return r
}

// func CoverLetterCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var coverLetter *CoverLetter
// 		var err error

// 		if coverLetterID := chi.URLParam(r, "CoverLetterID"); coverLetterID != "" {
// 			coverLetter, err = dbGetCoverLetter(coverLetterID)
// 		} else if CoverLetterSlug := chi.URLParam(r, "CoverLetterSlug"); CoverLetterSlug != "" {
// 			coverLetter, err = dbGetCoverLetterBySlug(CoverLetterSlug)
// 		} else {
// 			render.Render(w, r, ErrNotFound)
// 			return
// 		}
// 		if err != nil {
// 			render.Render(w, r, ErrNotFound)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "CoverLetter", coverLetter)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }