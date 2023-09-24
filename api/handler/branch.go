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

// @Router       /branches [POST]
// @Summary      Create Branch
// @Description  Create Branch
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateBranch  true  "branch data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateBranch(c *gin.Context) {

	var branch models.CreateUpdateBranch
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Branch().Create(branch)
	if err != nil {
		fmt.Println("error Branch Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /branches/{id} [put]
// @Summary      Update Branch
// @Description  api for update branch
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch" format(uuid)
// @Param        branch    body     models.CreateUpdateBranch  true  "data of branch"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateBranch(c *gin.Context) {

	var branch models.CreateUpdateBranch
	err := c.ShouldBindJSON(&branch)
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

	resp, err := h.strg.Branch().Update(branch, id)
	if err != nil {
		fmt.Println("error Branch Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /branches/{id} [GET]
// @Summary      Get By Id
// @Description  get branch by ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Branch ID" format(uuid)
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetBranch(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.Branch().Get(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Branch Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /branches [get]
// @Summary      List Branches
// @Description  get Branch
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Param        name    query     string  false  "filter by name"
// @Success      200  {array}   models.Branch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllBranch(c *gin.Context) {

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

	resp, err := h.strg.Branch().GetAll(models.GetAllBranchRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("name"),
	})
	if err != nil {
		h.log.Error("error Branch GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /branches/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete branch by ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Branch ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteBranch(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}
	resp, err := h.strg.Branch().Delete(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error Branch Check:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
