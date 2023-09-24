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

// @Router       /categories [POST]
// @Summary      Create Category
// @Description  Create Category
// @Tags         CATEGORY
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUpdateCategory  true  "category data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateCategory(c *gin.Context) {

	var category models.CreateUpdateCategory
	err := c.ShouldBindJSON(&category)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.strg.Category().Create(category)
	if err != nil {
		fmt.Println("error Category Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp})
}

// @Router       /categories/{id} [put]
// @Summary      Update Category
// @Description  api for update category
// @Tags         CATEGORY
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of category" format(uuid)
// @Param        category    body     models.CreateUpdateCategory  true  "data of category"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateCategory(c *gin.Context) {

	var category models.CreateUpdateCategory
	err := c.ShouldBindJSON(&category)
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

	resp, err := h.strg.Category().Update(category, id)
	if err != nil {
		fmt.Println("error Category Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /categories/{id} [GET]
// @Summary      Get By Id
// @Description  get category by ID
// @Tags         CATEGORY
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Category ID" format(uuid)
// @Success      200  {object}  models.Category
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetCategory(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}

	resp, err := h.strg.Category().Get(models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Category Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// @Router       /categories [get]
// @Summary      List Category
// @Description  get category
// @Tags         CATEGORY
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Param        name    query     string  false  "filter by name"
// @Success      200  {array}   models.Category
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllCategory(c *gin.Context) {

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

	resp, err := h.strg.Category().GetAll(models.GetAllCategoryRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("name"),
	})
	if err != nil {
		h.log.Error("error Category GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// @Router       /categories/{id} [DELETE]
// @Summary      Delete By Id
// @Description  delete category by ID
// @Tags         CATEGORY
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Category ID" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteCategory(c *gin.Context) {

	id := c.Param("id")

	if !helper.IsValidUUID(id) {

		h.log.Error("error Valid UUID:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return

	}
	resp, err := h.strg.Category().Delete(models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error Category Check:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}
