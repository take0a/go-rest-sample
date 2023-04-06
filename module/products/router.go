package products

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// NewRouter は、このモジュールのルーターを生成する。
func NewRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	c := NewController(db)

	r.Get("/{productId}", c.Get)
	r.Post("/", c.Post)
	r.Put("/{productId}", c.Put)
	r.Delete("/{productId}", c.Delete)

	return r
}
