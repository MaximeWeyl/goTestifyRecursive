package behaviours

import (
	"github.com/stretchr/testify/assert"
)

type NotEmptyField struct {
}

func (n NotEmptyField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.NotEmptyf(t, actualValueInterface, "Field {%s} should not be empty, but was", fieldName)
}
