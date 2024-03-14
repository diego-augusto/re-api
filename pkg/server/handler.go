package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetPacks returns the number of packs needed to fulfill the order
func (s server) GetPacks(c echo.Context) error {
	var input packRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httpError{Message: err.Error()})
	}
	packs := s.packSrv.GetPacks(input.Sizes, input.Items)
	response := make([]packResponse, 0)
	for k, v := range packs {
		response = append(response, packResponse{Size: k, Quantity: v})
	}
	return c.JSON(http.StatusOK, response)
}
