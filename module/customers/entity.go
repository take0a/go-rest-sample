package customers

// Entity は、CUSTOMER テーブルに対応するエンティティ
type Entity struct {
	CustomerID int
	Name       string
	Address    string
}

// NewEntity は、Dto から Entity を生成する
func NewEntity(dto *Dto) *Entity {
	return &Entity{
		CustomerID: dto.CustomerID,
		Name:       dto.Name,
		Address:    dto.Address,
	}
}
