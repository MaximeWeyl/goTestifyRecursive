package behaviours

import (
	"github.com/stretchr/testify/assert"
)

type FieldBehaviour interface {
	CheckField(t assert.TestingT, actualValueInterface interface{}, parentFieldName string) bool
}
