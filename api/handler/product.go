package handler

import (
	"errors"
	"fmt"
	"main/api/response"
	"main/models"
	"main/packages/helper"
	"main/packages/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router       /products [POST]
// @Summary      Create Product
// @Description  Create Product
// @Tags         PRODUCT
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateProduct  true  "product data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateProduct(c *gin.Context) {

	var product models.CreateUpdateProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Product().Create(product)
	if err != nil {
		fmt.Println("error Branch Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /products/{id} [put]
// @Summary      Update Product
// @Description  api for update product
// @Tags         PRODUCT
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of product" format(uuid)
// @Param        product    body     models.CreateUpdateProduct  true  "data of product"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateProduct(c *gin.Context) {

	var product models.CreateUpdateProduct
	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.Product().Update(product, id)
	if err != nil {
		fmt.Println("error Product Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /products/{id} [GET]
// @Summary      Get By Id
// @Description  get product by ID
// @Tags         PRODUCT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID" format(uuid)
// @Success      200  {object}  models.Product
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.Product().Get(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Product Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /products [get]
// @Summary      List Product
// @Description  get Product
// @Tags         PRODUCT
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  false  "limit for response"  Default(10)
// @Param        page    query     integer  false  "page of req"  Default(1)
// @Param        name    query     string  false  "filter by name"
// @Param        barcode    query     string  false  "filter by barcode"
// @Success      200  {array}   models.Product
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllProduct(c *gin.Context) {

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.strg.Product().GetAll(models.GetAllProductRequest{
		Page:    page,
		Limit:   limit,
		Name:    c.Query("name"),
		Barcode: c.Query("barcode"),
	})
	if err != nil {
		h.log.Error("error Product GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /products/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete product by ID
// @Tags         PRODUCT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Product ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}
	resp, err := h.strg.Product().Delete(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error Product Check:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
