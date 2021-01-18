package checkStruct

import (
	"github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func AssertRecursive(t assert.TestingT, actual interface{}, expected behaviours.FieldBehaviour) bool {
	return expected.CheckField(t, actual, "")
}

func RequireRecursive(t require.TestingT, actual interface{}, expected behaviours.FieldBehaviour) {
	if AssertRecursive(t, actual, expected) {
		return
	}
	t.FailNow()
}
