package models

type Remaining struct {
	Id         string  `json:"id"`
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
	TotalPrice float64 `json:"total_price"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}

type CreateUpdateRemaining struct {
	BranchId   string  `json:"branch_id"`
	CategoryId string  `json:"category_id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	Count      float64 `json:"count"`
}

type GetAllRemaining struct {
	Remainings []Remaining
	Count      int
}

type GetAllRemainingRequest struct {
	Page       int
	Limit      int
	BranchId   string
	CategoryId string
	Barcode    string
}
