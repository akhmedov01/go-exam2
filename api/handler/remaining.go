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

// @Router       /remaining [POST]
// @Summary      Create Remaining
// @Description  Create Remaining
// @Tags         REMAINING
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateRemaining  true  "Remaining data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateRemaining(c *gin.Context) {

	var remaining models.CreateUpdateRemaining
	err := c.ShouldBindJSON(&remaining)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Remaining().Create(remaining)
	if err != nil {
		fmt.Println("error Remaining Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /remaining/{id} [put]
// @Summary      Update Remaining
// @Description  api for update Remaining
// @Tags         REMAINING
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of Remaining" format(uuid)
// @Param        remaining    body     models.CreateUpdateRemaining  true  "data of Remaining"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateRemaining(c *gin.Context) {

	var remaining models.CreateUpdateRemaining
	err := c.ShouldBindJSON(&remaining)
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

	resp, err := h.strg.Remaining().Update(remaining, id)
	if err != nil {
		fmt.Println("error Remaining Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /remaining/{id} [GET]
// @Summary      Get By Id
// @Description  get Remaining by ID
// @Tags         REMAINING
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Remaining ID" format(uuid)
// @Success      200  {object}  models.Remaining
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetRemaining(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.Remaining().Get(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Product Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /remaining [get]
// @Summary      List Remaining
// @Description  get Remaining
// @Tags         REMAINING
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  false  "limit for response"  Default(10)
// @Param        page    query     integer  false  "page of req"  Default(1)
// @Param        category_id    query     string  false  "filter by categorId"
// @Param        branch_id    query     string  false  "filter by branchId"
// @Param        barcode    query     string  false  "filter by barcode"
// @Success      200  {array}   models.Remaining
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllRemaining(c *gin.Context) {

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

	resp, err := h.strg.Remaining().GetAll(models.GetAllRemainingRequest{
		Page:       page,
		Limit:      limit,
		CategoryId: c.Query("category_id"),
		BranchId:   c.Query("branch_id"),
		Barcode:    c.Query("barcode"),
	})
	if err != nil {
		h.log.Error("error Remaining GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /remaining/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete Remaining by ID
// @Tags         REMAINING
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Remaining ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteRemaining(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}
	resp, err := h.strg.Remaining().Delete(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error Remaining Check:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
