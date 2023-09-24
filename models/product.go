package models

type Product struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	CategoryId string  `json:"category_id"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}

type CreateUpdateProduct struct {
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Barcode    string  `json:"barcode"`
	CategoryId string  `json:"category_id"`
}

type GetAllProduct struct {
	Products []Product
	Count    int
}

type GetAllProductRequest struct {
	Page    int
	Limit   int
	Name    string
	Barcode string
}

type GetAllScanProductRequest struct {
	Barcode string
}
