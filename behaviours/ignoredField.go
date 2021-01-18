package behaviours

import (
	"github.com/stretchr/testify/assert"
)

type IgnoredField struct {
}

func (i IgnoredField) CheckField(_ assert.TestingT, _ interface{}, _ string) bool {
	return true
}
