package behaviours

import (
	"github.com/stretchr/testify/assert"
)

type SameValueField struct {
	Expected interface{}
}

func (s SameValueField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.EqualValues(t, s.Expected, actualValueInterface, "Field %s should have the same value as %v", fieldName, s.Expected)
}
