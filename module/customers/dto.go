package customers

// Key は、顧客リソースの主キー
type Key struct {
	CustomerID int
}

// Dto は、顧客リソース
type Dto struct {
	CustomerID int    `json:"customerId"`
	Name       string `json:"name"`
	Address    string `json:"address"`
}

// NewDto は、Customer から CustomerDto を生成する
func NewDto(customer *Entity) *Dto {
	return &Dto{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Address:    customer.Address,
	}
}

// Key は、Dto の Key を生成する。
func (c *Dto) Key() *Key {
	return &Key{
		CustomerID: c.CustomerID,
	}
}
