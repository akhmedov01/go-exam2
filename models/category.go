package models

type Category struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	ParentId   string `json:"parent_id"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateUpdateCategory struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type GetAllCategory struct {
	Categories []Category
	Count      int
}

type GetAllCategoryRequest struct {
	Page  int
	Limit int
	Name  string
}
