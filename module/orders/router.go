package orders

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// NewRouter は、このモジュールのルーターを生成する。
func NewRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	c := NewController(db)

	r.Get("/{orderId}", c.Get)
	r.Post("/", c.Post)
	r.Put("/{orderId}", c.Put)
	r.Delete("/{orderId}", c.Delete)

	return r
}
