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

// @Router       /comingproducts [POST]
// @Summary      Create ComingTableProduct
// @Description  Create ComingTableProduct
// @Tags         COMING_TABLE_PRODUCT
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateComingTableProduct true  "comingTableProduct data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateComingTableProduct(c *gin.Context) {

	var coming_table_product models.CreateUpdateComingTableProduct
	err := c.ShouldBindJSON(&coming_table_product)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.ComingTableProduct().Create(coming_table_product)
	if err != nil {
		fmt.Println("error ComingTableProduct Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /comingproducts/{id} [put]
// @Summary      Update ComingTableProduct
// @Description  api for update comingTableProduct
// @Tags         COMING_TABLE_PRODUCT
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTableProduct" format(uuid)
// @Param        coming_table_product    body     models.CreateUpdateComingTableProduct  true  "data of comingTableProduct"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateComingTableProduct(c *gin.Context) {

	var coming_table_product models.CreateUpdateComingTableProduct
	err := c.ShouldBindJSON(&coming_table_product)
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

	resp, err := h.strg.ComingTableProduct().Update(coming_table_product, id)
	if err != nil {
		fmt.Println("error ComingTableProduct Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /comingproducts/{id} [GET]
// @Summary      Get By Id
// @Description  get comingTableProduct by ID
// @Tags         COMING_TABLE_PRODUCT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "comingTableProduct ID" format(uuid)
// @Success      200  {object}  models.ComingTableProduct
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetComingTableProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.ComingTableProduct().Get(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error ComingTableProduct Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /comingproducts [get]
// @Summary      List ComingTableProduct
// @Description  get ComingTableProduct
// @Tags         COMING_TABLE_PRODUCT
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Param        name    query     string  false  "filter by name"
// @Param        barcode    query     string  false  "filter by barcode"
// @Success      200  {array}   models.ComingTableProduct
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllComingTableProduct(c *gin.Context) {

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

	resp, err := h.strg.ComingTableProduct().GetAll(models.GetAllComingTableProductRequest{
		Page:    page,
		Limit:   limit,
		Name:    c.Query("name"),
		Barcode: c.Query("barcode"),
	})
	if err != nil {
		h.log.Error("error ComingTableProduct GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /comingproducts/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete comingTableProduct by ID
// @Tags         COMING_TABLE_PRODUCT
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ComingTableProduct ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteComingTableProduct(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}
	resp, err := h.strg.ComingTableProduct().Delete(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error ComingTableProduct Check:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
