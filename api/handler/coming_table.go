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

// @Router       /comingtables [POST]
// @Summary      Create ComingTable
// @Description  Create ComingTable
// @Tags         COMINGTABLE
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateComingTable  true  "comingTable data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateComingTable(c *gin.Context) {

	var coming_table models.CreateUpdateComingTable
	err := c.ShouldBindJSON(&coming_table)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.ComingTable().Create(coming_table)
	if err != nil {
		fmt.Println("error ComingTable Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /comingtables/{id} [put]
// @Summary      Update ComingTable
// @Description  api for update comingTable
// @Tags         COMINGTABLE
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of comingTable" format(uuid)
// @Param        coming_table    body     models.CreateUpdateComingTable  true  "data of comingTable"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateComingTable(c *gin.Context) {

	var coming_table models.CreateUpdateComingTable
	err := c.ShouldBindJSON(&coming_table)
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

	resp, err := h.strg.ComingTable().Update(coming_table, id)
	if err != nil {
		fmt.Println("error ComingTable Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /comingtables/{id} [GET]
// @Summary      Get By Id
// @Description  get comingTable by ID
// @Tags         COMINGTABLE
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "comingTable ID" format(uuid)
// @Success      200  {object}  models.ComingTable
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetComingTable(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.ComingTable().Get(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error ComingTable Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /comingtables [get]
// @Summary      List ComingTable
// @Description  get ComingTable
// @Tags         COMINGTABLE
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Param        coming_id    query     string  false  "filter by comingId"
// @Param        branch_id    query     string  false  "filter by branchId"
// @Success      200  {array}   models.ComingTable
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllComingTable(c *gin.Context) {

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

	resp, err := h.strg.ComingTable().GetAll(models.GetAllComingTableRequest{
		Page:     page,
		Limit:    limit,
		ComingId: c.Query("coming_id"),
		BranchId: c.Query("branch_id"),
	})
	if err != nil {
		h.log.Error("error ComingTable GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /comingtables/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete comingTable by ID
// @Tags         COMINGTABLE
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ComingTable ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteComingTable(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}
	resp, err := h.strg.ComingTable().Delete(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error ComingTable Check:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
