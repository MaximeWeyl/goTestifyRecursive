package behaviours

import (
	"github.com/stretchr/testify/assert"
)

type ZeroValueField struct {
}

func (z ZeroValueField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.Zerof(t, actualValueInterface, "Field {%s} should have zero value but has value %v", fieldName, actualValueInterface)
}
