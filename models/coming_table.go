package models

type ComingTable struct {
	Id         string `json:"id"`
	ComingId   string `json:"coming_id"`
	BranchId   string `json:"branch_id"`
	DateTime   string `json:"date_time"`
	Status     string `json:"status"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateUpdateComingTable struct {
	ComingId string `json:"coming_id"`
	BranchId string `json:"branch_id"`
	DateTime string `json:"date_time"`
	Status   string `json:"status"`
}

type GetAllComingTable struct {
	ComingTables []ComingTable
	Count        int
}

type GetAllComingTableRequest struct {
	Page     int
	Limit    int
	BranchId string
	ComingId string
}
