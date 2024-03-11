package server

import (
	"errors"
	"re-go-challenge/pkg/packets"
)

var (
	// ErrInvalidPackSrv is the error message for an invalid pack service
	ErrInvalidPackSrv = errors.New("invalid pack service")
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

// New returns a new instance of the server
func New(packSrv packets.PackSrv) (*server, error) {

	if packSrv == nil {
		return nil, ErrInvalidPackSrv
	}

	s := server{
		packSrv: packSrv,
	}

	return &s, nil
}
