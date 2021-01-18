package behaviours

import (
	"github.com/stretchr/testify/assert"
)

//NotEmptyField  Assert/Require a non empty value
type NotEmptyField struct {
}

//CheckField Do the non empty assertion on a value
func (n NotEmptyField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.NotEmptyf(t, actualValueInterface, "Field {%s} should not be empty, but was", fieldName)
}
