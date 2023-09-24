package models

type Branch struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

type CreateUpdateBranch struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}

type IdRequest struct {
	Id string
}

type GetAllBranch struct {
	Branches []Branch
	Count    int
}

type GetAllBranchRequest struct {
	Page   int
	Limit  int
	Search string
}
