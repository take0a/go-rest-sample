package orders

// OrderKey は、注文リソースの主キー
type OrderKey struct {
	OrderID int
}

// OrderDto は、注文リソース
type OrderDto struct {
	OrderID    int               `json:"orderId"`
	CustomerID int               `json:"customerId"`
	OrderDate  string            `json:"orderDate"`
	Details    []*OrderDetailDto `json:"details"`
}

// OrderDetailDto は、注文リソースを構成する注文明細
type OrderDetailDto struct {
	RowNumber    int `json:"rowNumber"`
	ProductID    int `json:"productId"`
	Quantity     int `json:"quantity"`
	PricePerUnit int `json:"pricePerUnit"`
}

// NewOrderDto は、独立エンティティである OrderHeader から OrderDto を生成する
func NewOrderDto(orderHeader *OrderHeader) *OrderDto {
	return &OrderDto{
		OrderID:    orderHeader.OrderID,
		CustomerID: orderHeader.CustomerID,
		OrderDate:  orderHeader.OrderDate,
	}
}

// AddOrderDetails は、従属エンティティである OrderDetail を OrderDto に追加する
func (d *OrderDto) AddOrderDetails(orderDetails []*OrderDetail) {
	for _, item := range orderDetails {
		d.Details = append(d.Details, &OrderDetailDto{
			RowNumber:    item.RowNumber,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PricePerUnit: item.PricePerUnit,
		})
	}
}

// Key は、OrderDto の OrderKey を生成する
func (d *OrderDto) Key() *OrderKey {
	return &OrderKey{
		OrderID: d.OrderID,
	}
}
