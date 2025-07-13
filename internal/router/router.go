package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"segment-service/internal/handlers"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Segment routes
	r.Post("/segments", handlers.CreateSegmentHandler)
	r.Delete("/segments/{name}", handlers.DeleteSegmentHandler)
	r.Patch("/segments/{name}", handlers.UpdateSegmentHandler)

	// User routes
	r.Post("/users", handlers.CreateUserHandler)
	r.Post("/users/{userID}/segments", handlers.AddUserToSegmentHandler)
	r.Get("/users/{userID}/segments", handlers.GetUserSegmentsHandler)

	// Segment assignment routes
	r.Post("/segments/{name}/assign_random", handlers.AssignSegmentRandomlyHandler)

	return r
}
