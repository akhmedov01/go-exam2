package storage

import (
	"main/models"
)

type StorageI interface {
	Branch() BranchesI
	Category() CategoriesI
	Product() ProductI
	ComingTable() ComingTableI
	ComingTableProduct() ComingTableProductI
	Remaining() RemainingI
}

type BranchesI interface {
	Create(models.CreateUpdateBranch) (string, error)
	Update(models.CreateUpdateBranch, string) (string, error)
	Get(models.IdRequest) (models.Branch, error)
	GetAll(models.GetAllBranchRequest) (models.GetAllBranch, error)
	Delete(models.IdRequest) (string, error)
}

type CategoriesI interface {
	Create(models.CreateUpdateCategory) (string, error)
	Update(models.CreateUpdateCategory, string) (string, error)
	Get(models.IdRequest) (models.Category, error)
	GetAll(models.GetAllCategoryRequest) (models.GetAllCategory, error)
	Delete(models.IdRequest) (string, error)
}

type ProductI interface {
	Create(models.CreateUpdateProduct) (string, error)
	Update(models.CreateUpdateProduct, string) (string, error)
	Get(models.IdRequest) (models.Product, error)
	GetAll(models.GetAllProductRequest) (models.GetAllProduct, error)
	Delete(models.IdRequest) (string, error)
	GetScanProduct(string) (models.Product, error)
}

type ComingTableI interface {
	Create(models.CreateUpdateComingTable) (string, error)
	Update(models.CreateUpdateComingTable, string) (string, error)
	Get(models.IdRequest) (models.ComingTable, error)
	GetAll(models.GetAllComingTableRequest) (models.GetAllComingTable, error)
	Delete(models.IdRequest) (string, error)
}

type ComingTableProductI interface {
	Create(models.CreateUpdateComingTableProduct) (string, error)
	Update(models.CreateUpdateComingTableProduct, string) (string, error)
	Get(models.IdRequest) (models.ComingTableProduct, error)
	GetAll(models.GetAllComingTableProductRequest) (models.GetAllComingTableProduct, error)
	Delete(models.IdRequest) (string, error)
	ScanUpdate(models.ComingTableProduct, float64) (string, error)
	GetScanBarcode(models.GetScanBarcodeRequest) (models.ComingTableProduct, error)
}

type RemainingI interface {
	Create(models.CreateUpdateRemaining) (string, error)
	Update(models.CreateUpdateRemaining, string) (string, error)
	Get(models.IdRequest) (models.Remaining, error)
	GetAll(models.GetAllRemainingRequest) (models.GetAllRemaining, error)
	Delete(models.IdRequest) (string, error)
}
