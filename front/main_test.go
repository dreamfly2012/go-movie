package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {

	var a = 2
	var b = 5

	c := a / b

	fmt.Println(c)

	assert.Equal(t, 0, c)
}
