package server

import (
	"re-go-challenge/pkg/packets"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	s, err := New(packets.New())
	assert.NoError(t, err)
	assert.NotNil(t, s)
}

func Test_New_Error(t *testing.T) {
	s, err := New(nil)
	assert.Equal(t, ErrInvalidPackSrv, err)
	assert.Nil(t, s)
}
