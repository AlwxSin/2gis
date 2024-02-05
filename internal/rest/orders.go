package rest

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"applicationDesignTest/internal/logic"
)

func (s *Server) createOrder(w http.ResponseWriter, r *http.Request) {
	newOrder := &logic.NewOrderRequest{}
	if err := render.Bind(r, newOrder); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	order, err := logic.CreateOrder(r.Context(), s.DB, newOrder)
	if err != nil {
		logicErr := &logic.Error{}
		if errors.As(err, logicErr) {
			render.Render(w, r, ErrInvalidRequest(err))
			return
		}
		render.Render(w, r, ErrInternalServer(err))
		return
	}

	render.SetContentType(render.ContentTypeJSON)
	render.Status(r, http.StatusCreated)
	render.Render(w, r, order)
}
