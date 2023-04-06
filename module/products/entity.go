package products

// Entity は、PRODUCT テーブルに対応するエンティティ
type Entity struct {
	ProductID    int
	Name         string
	PricePerUnit int
}

// NewEntity は、Dto から Entity を生成する
func NewEntity(dto *Dto) *Entity {
	return &Entity{
		ProductID:    dto.ProductID,
		Name:         dto.Name,
		PricePerUnit: dto.PricePerUnit,
	}
}
