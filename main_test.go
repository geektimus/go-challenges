package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntro(t *testing.T) {
	i := getIntro()
	e := "Golang Challenges"
	assert.Equal(t, e, i)
}
