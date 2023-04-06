package customers

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

// NewRouter は、このモジュールのルーターを生成する。
func NewRouter(db *sql.DB) chi.Router {
	r := chi.NewRouter()
	c := NewController(db)

	r.Get("/{customerId}", c.Get)
	r.Post("/", c.Post)
	r.Put("/{customerId}", c.Put)
	r.Delete("/{customerId}", c.Delete)

	return r
}
