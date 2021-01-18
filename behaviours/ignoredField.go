package behaviours

import (
	"github.com/stretchr/testify/assert"
)

//IgnoredField  Ignore this value
type IgnoredField struct {
}

//CheckField Never returns an error
func (i IgnoredField) CheckField(_ assert.TestingT, _ interface{}, _ string) bool {
	return true
}
