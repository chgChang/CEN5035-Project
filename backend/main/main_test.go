package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	service := ...
	firstName, lastName := service.find(someParams)
	assert.Equal(t, "John", firstName)
	assert.Equal(t, "Dow", lastName)
}

func StringEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}