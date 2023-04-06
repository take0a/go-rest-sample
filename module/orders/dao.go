package orders

import (
	"context"
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/take0a/go-rest-sample/utils"
)

// SQL文
const (
	InsertOrderHeader = `
INSERT INTO ORDER_HEADER (
ORDER_ID, CUSTOMER_ID, ORDER_DATE
) VALUES (
$1, $2, $3
)`
	SelectOrderHeader = `
SELECT
ORDER_ID, CUSTOMER_ID, ORDER_DATE
FROM ORDER_HEADER
WHERE ORDER_ID = $1`
	UpdateOrderHeader = `
UPDATE ORDER_HEADER SET
ORDER_ID = $1, 
CUSTOMER_ID = $2,
ORDER_DATE = $3
WHERE ORDER_ID = $1`
	DeleteOrderHeader = `
DELETE FROM ORDER_HEADER
WHERE ORDER_ID = $1`

	InsertOrderDetail = `
INSERT INTO ORDER_DETAIL (
ORDER_ID, ROW_NUMBER, PRODUCT_ID, QUANTITY, PRICE_PER_UNIT
) VALUES (
$1, $2, $3, $4, $5
)`
	SelectOrderDetail = `
SELECT
ORDER_ID, ROW_NUMBER, PRODUCT_ID, QUANTITY, PRICE_PER_UNIT
FROM ORDER_DETAIL 
WHERE ORDER_ID = $1`
	DeleteOrderDetail = `
DELETE FROM ORDER_DETAIL
WHERE ORDER_ID = $1`
)

// OrderHeaderDao は、OrderHeader の Table Data Gateway。
type OrderHeaderDao struct{}

// Insert は、指定された OrderHeader を登録する。
func (d *OrderHeaderDao) Insert(ctx context.Context, tx *sql.Tx, entity *OrderHeader) (*OrderHeader, error) {
	_, err := tx.ExecContext(ctx, InsertOrderHeader,
		entity.OrderID,
		entity.CustomerID,
		entity.OrderDate,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				return nil, utils.ErrConflict
			}
		}
		log.Printf("%s %s\n", InsertOrderHeader, err)
		return nil, err
	}
	return entity, nil
}

// Select は、指定されたキーの OrderHeader を取得する。
func (d *OrderHeaderDao) Select(ctx context.Context, tx *sql.Tx, key *OrderKey) (*OrderHeader, error) {
	var entity OrderHeader
	err := tx.QueryRowContext(ctx, SelectOrderHeader,
		key.OrderID,
	).Scan(
		&entity.OrderID,
		&entity.CustomerID,
		&entity.OrderDate,
	)
	if err != nil {
		// レコードが存在しない場合は、sql.ErrNoRows
		log.Printf("%s Query %s\n", SelectOrderHeader, err)
		return nil, err
	}
	return &entity, nil
}

// Update は、指定された OrderHeader を更新する。
func (d *OrderHeaderDao) Update(ctx context.Context, tx *sql.Tx, entity *OrderHeader) (*OrderHeader, error) {
	result, err := tx.ExecContext(ctx, UpdateOrderHeader,
		entity.OrderID,
		entity.CustomerID,
		entity.OrderDate,
	)
	if err != nil {
		log.Printf("%s Exec %s\n", UpdateOrderHeader, err)
		return nil, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s RowsAffected %s\n", UpdateOrderHeader, err)
		return nil, err
	}
	if num == 0 {
		log.Printf("OrderHeader Not Found %#v", entity)
		return nil, sql.ErrNoRows
	}
	return entity, nil
}

// Delete は、指定されたキーの OrderHeader を削除する。
func (d *OrderHeaderDao) Delete(ctx context.Context, tx *sql.Tx, key *OrderKey) error {
	result, err := tx.ExecContext(ctx, DeleteOrderHeader,
		key.OrderID,
	)
	if err != nil {
		log.Printf("%s %s\n", DeleteOrderHeader, err)
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s RowsAffected %s\n", DeleteOrderHeader, err)
		return err
	}
	if num == 0 {
		log.Printf("OrderHeader Not Found %#v", key)
		return sql.ErrNoRows
	}
	return nil
}

// OrderDetailDao は、OrderDetail の Table Data Gateway。
type OrderDetailDao struct{}

// Insert は、指定された OrderDetail のリストを登録する。
func (d *OrderDetailDao) Insert(ctx context.Context, tx *sql.Tx, list []*OrderDetail) ([]*OrderDetail, error) {
	stmt, err := tx.PrepareContext(ctx, InsertOrderDetail)
	if err != nil {
		log.Printf("%s Prepare %s\n", InsertOrderDetail, err)
		return nil, err
	}
	defer stmt.Close()

	for _, item := range list {
		_, err := stmt.ExecContext(ctx,
			item.OrderID,
			item.RowNumber,
			item.ProductID,
			item.Quantity,
			item.PricePerUnit,
		)
		if err != nil {
			log.Printf("%s Exec %s\n", InsertOrderDetail, err)
			return nil, err
		}
	}
	return list, nil
}

// Select は、指定されたキーの OrderDetail のリストを取得する。
func (d *OrderDetailDao) Select(ctx context.Context, tx *sql.Tx, key *OrderKey) ([]*OrderDetail, error) {
	var list []*OrderDetail
	rows, err := tx.QueryContext(ctx, SelectOrderDetail,
		key.OrderID,
	)
	if err != nil {
		log.Printf("%s Query %s\n", SelectOrderDetail, err)
		return nil, err
	}
	for rows.Next() {
		var entity OrderDetail
		err = rows.Scan(
			&entity.OrderID,
			&entity.RowNumber,
			&entity.ProductID,
			&entity.Quantity,
			&entity.PricePerUnit,
		)
		if err != nil {
			log.Printf("%s Scan %s\n", SelectOrderDetail, err)
			return nil, err
		}
		list = append(list, &entity)
	}
	return list, nil
}

// Delete は、指定されたキーの OrderDetail を削除する。
func (d *OrderDetailDao) Delete(ctx context.Context, tx *sql.Tx, key *OrderKey) error {
	_, err := tx.ExecContext(ctx, DeleteOrderDetail,
		key.OrderID,
	)
	if err != nil {
		log.Printf("%s %s\n", DeleteOrderDetail, err)
		return err
	}
	return nil
}
