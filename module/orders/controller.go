package orders

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/take0a/go-rest-sample/utils"
)

// Controller は、REST呼び出しを処理する。
type Controller struct {
	db      *sql.DB
	service *Service
}

// NewController は、Controller を生成する。
func NewController(db *sql.DB) *Controller {
	return &Controller{
		db:      db,
		service: NewService(),
	}
}

// Get は、指定されたキーのリソースを返す。
func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "orderId")
	id, err := strconv.Atoi(param)
	if err != nil {
		utils.BadRequest(w, err, fmt.Sprintf("Invalid URL Parameter %s", param))
		return
	}

	ctx := r.Context()
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		utils.ServerError(w, err, "BeginTx")
		return
	}
	defer tx.Rollback()

	dto, err := c.service.Read(ctx, tx, &OrderKey{OrderID: id})
	if err != nil {
		if err == sql.ErrNoRows {
			utils.NotFound(w, err, fmt.Sprintf("Invalid URL Parameter %s", param))
		} else {
			utils.ServerError(w, err, "service.Find")
		}
		return
	}

	render.JSON(w, r, dto)
	tx.Commit()
}

// Post は、指定されたリソースを保存する。
func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var dto OrderDto
	err := render.DecodeJSON(r.Body, &dto)
	if err != nil {
		utils.BadRequest(w, err, "render.DecodeJSON")
		return
	}

	ctx := r.Context()
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		utils.ServerError(w, err, "BeginTx")
		return
	}
	defer tx.Rollback()

	res, err := c.service.Create(ctx, tx, &dto)
	if err != nil {
		if err == utils.ErrConflict {
			utils.Conflict(w, err, "serice.Insert")
		} else {
			utils.ServerError(w, err, "service.Insert")
		}
		return
	}

	render.JSON(w, r, res)
	tx.Commit()
}

// Put は、指定されたリソースを更新する。
func (c *Controller) Put(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "orderId")
	id, err := strconv.Atoi(param)
	if err != nil {
		utils.BadRequest(w, err, fmt.Sprintf("Invalid URL Parameter %s", param))
		return
	}

	var dto OrderDto
	err = render.DecodeJSON(r.Body, &dto)
	if err != nil {
		utils.BadRequest(w, err, "render.DecodeJSON")
		return
	}
	if id != dto.Key().OrderID {
		utils.BadRequest(w, err, fmt.Sprintf("Unmatch URL Parameter %s", param))
		return
	}

	ctx := r.Context()
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		utils.ServerError(w, err, "BeginTx")
		return
	}
	defer tx.Rollback()

	res, err := c.service.Update(ctx, tx, &dto)
	if err != nil {
		utils.ServerError(w, err, "service.Insert")
		return
	}

	render.JSON(w, r, res)
	tx.Commit()
}

// Delete は、指定されたキーのリソースを削除する。
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "orderId")
	id, err := strconv.Atoi(param)
	if err != nil {
		utils.BadRequest(w, err, fmt.Sprintf("Invalid URL Parameter %s", param))
		return
	}

	ctx := r.Context()
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		utils.ServerError(w, err, "BeginTx")
		return
	}
	defer tx.Rollback()

	err = c.service.Delete(ctx, tx, &OrderKey{OrderID: id})
	if err == sql.ErrNoRows {
		utils.NotFound(w, err, fmt.Sprintf("Invalid URL Parameter %s", param))
		return
	}
	if err != nil {
		utils.ServerError(w, err, "service.Delete")
		return
	}
	w.WriteHeader(http.StatusOK)
	tx.Commit()
}
