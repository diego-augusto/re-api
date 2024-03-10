package server

import (
	"net/http"
	"re-go-challenge/pkg/packets"

	"github.com/labstack/echo/v4"
)

type httpError struct {
	Message string `json:"message"`
}

type server struct {
	packSrv packets.PackSrv
}

type packRequest struct {
	Sizes []int `json:"sizes"`
	Items int   `json:"items"`
}

type packResponse struct {
	Size     int `json:"size"`
	Quantity int `json:"quantity"`
}

// GetPacks returns the number of packs needed to fulfill the order
func (s server) GetPacks(c echo.Context) error {
	var input packRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, httpError{Message: err.Error()})
	}
	packs := s.packSrv.GetPacks(input.Items, input.Sizes)
	response := make([]packResponse, 0)
	for _, p := range packs {
		response = append(response, packResponse{Size: p.Size, Quantity: p.Quantity})
	}
	return c.JSON(http.StatusOK, response)
}

// New returns a new instance of the server
func New(packSrv packets.PackSrv) *server {
	s := server{
		packSrv: packSrv,
	}

	return &s
}
