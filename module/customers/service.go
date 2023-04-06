package customers

import (
	"context"
	"database/sql"
)

// Service は、Contorller からの要求を Dao を使用して解決する。
type Service struct {
	dao *Dao
}

// NewService は、Service を生成する。
func NewService() *Service {
	return &Service{
		dao: &Dao{},
	}
}

// Create は、指定された Dto を登録する。
func (s *Service) Create(ctx context.Context, tx *sql.Tx, dto *Dto) (*Dto, error) {
	entity := NewEntity(dto)
	entity, err := s.dao.Insert(ctx, tx, entity)
	if err != nil {
		return nil, err
	}
	return NewDto(entity), nil
}

// Read は、指定されたキーの Dto を返す。
func (s *Service) Read(ctx context.Context, tx *sql.Tx, key *Key) (*Dto, error) {
	entity, err := s.dao.Select(ctx, tx, key)
	if err != nil {
		return nil, err
	}
	return NewDto(entity), nil
}

// Update は、指定された Dto を更新する。
func (s *Service) Update(ctx context.Context, tx *sql.Tx, dto *Dto) (*Dto, error) {
	entity := NewEntity(dto)
	entity, err := s.dao.Update(ctx, tx, entity)
	if err != nil {
		return nil, err
	}
	return NewDto(entity), nil
}

// Delete は、指定されたキーの Dto を削除する。
func (s *Service) Delete(ctx context.Context, tx *sql.Tx, key *Key) error {
	return s.dao.Delete(ctx, tx, key)
}
