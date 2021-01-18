package behaviours

import (
	"github.com/stretchr/testify/assert"
)

//SameValueField Assert/Require that the value is the same, even if the type differs
type SameValueField struct {
	Expected interface{}
}

//CheckField Performs the assertion of same value
func (s SameValueField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.EqualValues(t, s.Expected, actualValueInterface, "Field %s should have the same value as %v", fieldName, s.Expected)
}
