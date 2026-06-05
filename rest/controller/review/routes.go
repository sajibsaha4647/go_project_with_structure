package review

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) ReviewRoutes(mux *http.ServeMux, mm *middleware.MiddlewareManager) {

	// mux.Handle("GET /getAllReviews", mm.Apply(middleware.AuthenticationMiddleware(h.getAllReviews)))
	// mux.Handle("POST /createReview", mm.Apply(middleware.AuthenticationMiddleware(h.CreateReview)))
	// mux.Handle("PUT /updateReview", mm.Apply(middleware.AuthenticationMiddleware(h.UpdateReviews)))
	// mux.Handle("DELETE /deleteReview", mm.Apply(middleware.AuthenticationMiddleware(h.DeleteReview)))

}
