package behaviours

import (
	"github.com/stretchr/testify/assert"
)

//FieldBehaviour The interface to implement in order to have new behaviours
type FieldBehaviour interface {
	CheckField(t assert.TestingT, actualValueInterface interface{}, parentFieldName string) bool
}
