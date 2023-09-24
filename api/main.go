package api

import (
	_ "main/api/docs"
	"main/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/branches", h.CreateBranch)
	r.GET("/branches", h.GetAllBranch)
	r.GET("/branches/:id", h.GetBranch)
	r.PUT("/branches/:id", h.UpdateBranch)
	r.DELETE("/branches/:id", h.DeleteBranch)

	r.POST("/categories", h.CreateCategory)
	r.GET("/categories", h.GetAllCategory)
	r.GET("/categories/:id", h.GetCategory)
	r.PUT("/categories/:id", h.UpdateCategory)
	r.DELETE("/categories/:id", h.DeleteCategory)

	r.POST("/products", h.CreateProduct)
	r.GET("/products", h.GetAllProduct)
	r.GET("/products/:id", h.GetProduct)
	r.PUT("/products/:id", h.UpdateProduct)
	r.DELETE("/products/:id", h.DeleteProduct)

	r.POST("/comingtables", h.CreateComingTable)
	r.GET("/comingtables", h.GetAllComingTable)
	r.GET("/comingtables/:id", h.GetComingTable)
	r.PUT("/comingtables/:id", h.UpdateComingTable)
	r.DELETE("/comingtables/:id", h.DeleteComingTable)

	r.POST("/comingproducts", h.CreateComingTableProduct)
	r.GET("/comingproducts", h.GetAllComingTableProduct)
	r.GET("/comingproducts/:id", h.GetComingTableProduct)
	r.PUT("/comingproducts/:id", h.UpdateComingTableProduct)
	r.DELETE("/comingproducts/:id", h.DeleteComingTableProduct)

	r.POST("/remaining", h.CreateRemaining)
	r.GET("/remaining", h.GetAllRemaining)
	r.GET("/remaining/:id", h.GetRemaining)
	r.PUT("/remaining/:id", h.UpdateRemaining)
	r.DELETE("/remaining/:id", h.DeleteRemaining)

	r.POST("/comingtables/scan-barcode/:id", h.InsertOrUpdate)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
