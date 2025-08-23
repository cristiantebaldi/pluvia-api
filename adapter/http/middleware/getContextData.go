package middleware

import (
	"context"
	"net/http"

	"github.com/pluvia/pluvia-api/core/domain"
)

func setContextData(r *http.Request, d *domain.Administrador) (ro *http.Request) {
	ctx := r.Context()
	ctx = context.WithValue(ctx, 1, d)
	ro = r.WithContext(ctx)
	return
}

func GetContextData(r *http.Request) (d domain.Administrador) {
	d = *r.Context().Value(1).(*domain.Administrador)
	return
}
