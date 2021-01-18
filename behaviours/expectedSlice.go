package behaviours

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
)

//ExpectedSlice Assert/Require a slice. The values are the behaviours we expect for each element
type ExpectedSlice []ExpectedStruct

//CheckField Assert that a value slice matches the expected slice
func (e ExpectedSlice) CheckField(t assert.TestingT, actualValueInterface interface{}, parentFieldName string) bool {
	actualValue := reflect.ValueOf(actualValueInterface)

	// A list of maps should correspond to a list of structs
	if actualValue.Kind() != reflect.Slice {
		return assert.Fail(
			t,
			"No slice found",
			"A slice was expected for field {%s}. We found : %v",
			parentFieldName,
			actualValueInterface,
		)
	}

	// The two slices lengths must match
	if actualValue.Len() != len(e) {
		return assert.Fail(
			t,
			"Wrong slice size",
			"For field {%s}, The actual slice has a len of %d, while a size of %d was expected",
			parentFieldName,
			actualValue.Len(),
			len(e),
		)
	}

	// All elements should match : recursion !
	var success = true
	for itemIndex, expectedItem := range e {
		itemValue := actualValue.Index(itemIndex)
		success = expectedItem.CheckField(
			t,
			itemValue.Interface(),
			fmt.Sprintf("%s[%d]", parentFieldName, itemIndex),
		) && success
	}
	return success
}
