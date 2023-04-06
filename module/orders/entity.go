package orders

// OrderHeader は、ORDER_HEADER テーブルに対応するエンティティ
type OrderHeader struct {
	OrderID    int
	CustomerID int
	OrderDate  string
}

// OrderDetail は、ORDER_DETAIL テーブルに対応するエンティティ
type OrderDetail struct {
	OrderID      int
	RowNumber    int
	ProductID    int
	Quantity     int
	PricePerUnit int
}

// NewOrderHeader は、OrderDto から OrderHeader を生成する
func NewOrderHeader(orderDto *OrderDto) *OrderHeader {
	return &OrderHeader{
		OrderID:    orderDto.OrderID,
		CustomerID: orderDto.CustomerID,
		OrderDate:  orderDto.OrderDate,
	}
}

// NewOrderDetails は、OrderDto から OrderDetail のリストを生成する
func NewOrderDetails(orderDto *OrderDto) []*OrderDetail {
	var list []*OrderDetail
	for _, item := range orderDto.Details {
		list = append(list, &OrderDetail{
			OrderID:      orderDto.OrderID,
			RowNumber:    item.RowNumber,
			ProductID:    item.ProductID,
			Quantity:     item.Quantity,
			PricePerUnit: item.PricePerUnit,
		})
	}
	return list
}
