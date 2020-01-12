package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	service := NewIDService("123")
	id, err := service.EncodeID(123)
	fmt.Print(id)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}
func TestDecode(t *testing.T)  {
	service := NewIDService("123")
	id, err := service.DecodeID("E4gVbBwDy013oLPOxlMp5r7k6WnR9Z")
	fmt.Print(id)
	assert.Nil(t, err)
	assert.NotEmpty(t, id)
}