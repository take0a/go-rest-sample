package products

// Key は、製品リソースの主キー
type Key struct {
	ProductID int
}

// Dto は、製品リソース
type Dto struct {
	ProductID    int    `json:"productId"`
	Name         string `json:"name"`
	PricePerUnit int    `json:"pricePerUnit"`
}

// NewDto は、Product から ProductDto を生成する
func NewDto(product *Entity) *Dto {
	return &Dto{
		ProductID:    product.ProductID,
		Name:         product.Name,
		PricePerUnit: product.PricePerUnit,
	}
}

// Key は、Dto の Key を生成する
func (p *Dto) Key() *Key {
	return &Key{
		ProductID: p.ProductID,
	}
}
