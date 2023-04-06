package products

import (
	"context"
	"database/sql"
	"log"

	"github.com/lib/pq"
	"github.com/take0a/go-rest-sample/utils"
)

// SQL文
const (
	Insert = `
INSERT INTO PRODUCT (
PRODUCT_ID, NAME, PRICE_PER_UNIT
) VALUES (
$1, $2, $3
)`
	Select = `
SELECT
PRODUCT_ID, NAME, PRICE_PER_UNIT
FROM PRODUCT
WHERE PRODUCT_ID = $1`
	Update = `
UPDATE PRODUCT SET
PRODUCT_ID = $1, 
NAME = $2,
PRICE_PER_UNIT = $3
WHERE PRODUCT_ID = $1`
	Delete = `
DELETE FROM PRODUCT
WHERE PRODUCT_ID = $1`
)

// Dao は、Product の Table Data Gateway
type Dao struct{}

// Insert は、指定された Product を登録する。
func (d *Dao) Insert(ctx context.Context, tx *sql.Tx, entity *Entity) (*Entity, error) {
	_, err := tx.ExecContext(ctx, Insert,
		entity.ProductID,
		entity.Name,
		entity.PricePerUnit,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" { // unique_violation
				return nil, utils.ErrConflict
			}
		}
		log.Printf("%s %s\n", Insert, err)
		return nil, err
	}
	return entity, nil
}

// Select は、指定されたキーの Product を取得する。
func (d *Dao) Select(ctx context.Context, tx *sql.Tx, key *Key) (*Entity, error) {
	var entity Entity
	err := tx.QueryRowContext(ctx, Select,
		key.ProductID,
	).Scan(
		&entity.ProductID,
		&entity.Name,
		&entity.PricePerUnit,
	)
	if err != nil {
		// レコードが存在しない場合は、sql.ErrNoRows
		log.Printf("%s Query %s\n", Select, err)
		return nil, err
	}
	return &entity, nil
}

// Update は、指定された Product を更新する。
func (d *Dao) Update(ctx context.Context, tx *sql.Tx, entity *Entity) (*Entity, error) {
	result, err := tx.ExecContext(ctx, Update,
		entity.ProductID,
		entity.Name,
		entity.PricePerUnit,
	)
	if err != nil {
		log.Printf("%s Exec %s\n", Update, err)
		return nil, err
	}
	num, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s RowsAffected %s\n", Update, err)
		return nil, err
	}
	if num == 0 {
		log.Printf("Not Found %#v", entity)
		return nil, sql.ErrNoRows
	}
	return entity, nil
}

// Delete は、指定されたキーの Product を削除する。
func (d *Dao) Delete(ctx context.Context, tx *sql.Tx, key *Key) error {
	result, err := tx.ExecContext(ctx, Delete,
		key.ProductID,
	)
	if err != nil {
		log.Printf("%s %s\n", Delete, err)
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		log.Printf("%s RowsAffected %s\n", Delete, err)
		return err
	}
	if num == 0 {
		log.Printf("Not Found %#v", key)
		return sql.ErrNoRows
	}
	return nil
}
