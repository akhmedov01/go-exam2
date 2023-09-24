package handler

import (
	"errors"
	"fmt"
	"main/models"
	"main/packages/helper"
	"main/packages/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router       /comingtables/scan-barcode/{id} [post]
// @Summary      Post From Scan
// @Description  insert ComingTable
// @Tags         SCAN_BARCODE
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ComingTable ID" format(uuid)
// @Param		 count      query     number   false "Count of Product" Default(1)
// @Param        barcode    query     string  false  "filter by barcode"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) InsertOrUpdate(c *gin.Context) {

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

	if resp.Status == "Finished" {
		fmt.Println("Coming_Table already finished")
		return
	}

	respProduct, err := h.strg.Product().GetScanProduct(c.Query("barcode"))

	if err != nil {
		fmt.Println("Product Not Found")
		h.log.Error("error Product GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	respComingProduct, err := h.strg.ComingTableProduct().GetScanBarcode(models.GetScanBarcodeRequest{
		Barcode:       c.Query("barcode"),
		ComingTableId: id,
	})

	if err != nil {
		h.log.Error("error ComingTableProduct GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	countS := c.Query("count")
	count, err := strconv.Atoi(countS)

	if err != nil {
		h.log.Error("error get count:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid count param")
		return
	}

	if respComingProduct.Count > 0 {

		_, err := h.strg.ComingTableProduct().ScanUpdate(respComingProduct, float64(count))
		if err != nil {
			fmt.Println("error ComingTableProduct Update:", err.Error())
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

	} else {

		_, err = h.strg.ComingTableProduct().Create(models.CreateUpdateComingTableProduct{
			CategoryId:    respProduct.CategoryId,
			Name:          respProduct.Name,
			Price:         respProduct.Price,
			Barcode:       respProduct.Barcode,
			Count:         float64(count),
			ComingTableId: id,
		})

		if err != nil {
			fmt.Println("error ComingTableProduct Create:", err.Error())
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

	}

	c.JSON(http.StatusOK, resp)

}
