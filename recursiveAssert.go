package checkStruct

import (
	"github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//AssertRecursive Performs a recursive assertion between an expected behaviour, and an actual value
func AssertRecursive(t assert.TestingT, actual interface{}, expected behaviours.FieldBehaviour) bool {
	return expected.CheckField(t, actual, "")
}

//RequireRecursive Same as AssertRecursive, but stops the test if failed
func RequireRecursive(t require.TestingT, actual interface{}, expected behaviours.FieldBehaviour) {
	if AssertRecursive(t, actual, expected) {
		return
	}
	t.FailNow()
}
