package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func failOnError(t *testing.T, err error, msg string) {
	if !assert.NoError(t, err, msg) {
		t.FailNow()
	}
}
