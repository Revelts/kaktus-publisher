package Routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"kaktus-services/Controllers"
	"net/http"
)

func HandleRequests() (err error) {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Route("/kaktus", func(r chi.Router) {
			r.Get("/", Controllers.ViewAllThreads)
			r.Get("/thread/{id}", Controllers.ViewThreadDetails)
			r.Post("/", Controllers.CreateThread)
			r.Post("/likeThread", Controllers.LikeThread)
			r.Post("/commentThread", Controllers.CommentThread)
			r.Post("/replyComment", Controllers.ReplyComment)
		})
	})
	err = http.ListenAndServe(":3000", r)
	return
}
