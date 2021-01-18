package behaviours

import (
	"github.com/stretchr/testify/assert"
)

//ZeroValueField Checks that the field value is the zero value for it's type
type ZeroValueField struct {
}

//CheckField Performs the assertion of zero value
func (z ZeroValueField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.Zerof(t, actualValueInterface, "Field {%s} should have zero value but has value %v", fieldName, actualValueInterface)
}
