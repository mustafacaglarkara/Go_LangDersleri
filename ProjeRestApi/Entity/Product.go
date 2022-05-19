package Entity

type Product struct {
	ID          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryID  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}
