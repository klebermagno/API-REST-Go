package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealth(t *testing.T) {
	returned := health("Ola")
	assert.Equal(t, "Ola", returned)
}
