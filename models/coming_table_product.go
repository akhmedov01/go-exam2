package models

type ComingTableProduct struct {
	Id            string  `json:"id"`
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         float64 `json:"count"`
	TotalPrice    float64 `json:"total_price"`
	ComingTableId string  `json:"coming_table_id"`
	Created_at    string  `json:"created_at"`
	Updated_at    string  `json:"updated_at"`
}

type CreateUpdateComingTableProduct struct {
	CategoryId    string  `json:"category_id"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Barcode       string  `json:"barcode"`
	Count         float64 `json:"count"`
	ComingTableId string  `json:"coming_table_id"`
}

type GetAllComingTableProduct struct {
	ComingTableProducts []ComingTableProduct
	Count               int
}

type GetAllComingTableProductRequest struct {
	Page    int
	Limit   int
	Name    string
	Barcode string
}

type GetScanBarcodeRequest struct {
	ComingTableId string
	Barcode       string
}

type CreateScanBarcodeRequest struct {
}
