package orders

import (
	"context"
	"database/sql"
)

// Service は、Contorller からの要求を Table Data Gateway を使用して解決する。
type Service struct {
	orderHeaderDao *OrderHeaderDao
	orderDetailDao *OrderDetailDao
}

// NewService は、Service を生成する。
func NewService() *Service {
	return &Service{
		orderHeaderDao: &OrderHeaderDao{},
		orderDetailDao: &OrderDetailDao{},
	}
}

// Create は、指定された OrderDto を登録する。
func (s *Service) Create(ctx context.Context, tx *sql.Tx, dto *OrderDto) (*OrderDto, error) {
	orderHeader := NewOrderHeader(dto)
	orderHeader, err := s.orderHeaderDao.Insert(ctx, tx, orderHeader)
	if err != nil {
		return nil, err
	}
	orderDetails := NewOrderDetails(dto)
	orderDetails, err = s.orderDetailDao.Insert(ctx, tx, orderDetails)
	if err != nil {
		return nil, err
	}
	res := NewOrderDto(orderHeader)
	res.AddOrderDetails(orderDetails)
	return res, nil
}

// Read は、指定されたキーの OrderDto を返す。
func (s *Service) Read(ctx context.Context, tx *sql.Tx, key *OrderKey) (*OrderDto, error) {
	orderHeader, err := s.orderHeaderDao.Select(ctx, tx, key)
	if err != nil {
		return nil, err
	}
	orderDetails, err := s.orderDetailDao.Select(ctx, tx, key)
	if err != nil {
		return nil, err
	}
	dto := NewOrderDto(orderHeader)
	dto.AddOrderDetails(orderDetails)
	return dto, nil
}

// Update は、指定された OrderDto を更新する。
func (s *Service) Update(ctx context.Context, tx *sql.Tx, dto *OrderDto) (*OrderDto, error) {
	orderHeader := NewOrderHeader(dto)
	orderHeader, err := s.orderHeaderDao.Update(ctx, tx, orderHeader)
	if err != nil {
		return nil, err
	}
	key := dto.Key()
	err = s.orderDetailDao.Delete(ctx, tx, key)
	if err != nil {
		return nil, err
	}
	orderDetails := NewOrderDetails(dto)
	orderDetails, err = s.orderDetailDao.Insert(ctx, tx, orderDetails)
	if err != nil {
		return nil, err
	}
	res := NewOrderDto(orderHeader)
	res.AddOrderDetails(orderDetails)
	return res, nil
}

// Delete は、指定されたキーの OrderDto を削除する。
func (s *Service) Delete(ctx context.Context, tx *sql.Tx, key *OrderKey) error {
	err := s.orderHeaderDao.Delete(ctx, tx, key)
	if err != nil {
		return err
	}
	err = s.orderDetailDao.Delete(ctx, tx, key)
	if err != nil {
		return err
	}
	return nil
}
