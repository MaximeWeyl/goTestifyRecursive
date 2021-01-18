package behaviours

import (
	"fmt"
	"github.com/MaximeWeyl/goTestifyRecursive/formatting"
	"github.com/stretchr/testify/assert"
	"reflect"
)

//ExpectedStruct Assert/Require a struct. The values are the behaviours we expect for each element
type ExpectedStruct map[string]interface{}

//CheckField Assert that a value struct matches the expected slice
func (e ExpectedStruct) CheckField(t assert.TestingT, actualValueInterface interface{}, parentFieldName string) bool {
	actualType := reflect.TypeOf(actualValueInterface)
	actualValue := reflect.ValueOf(actualValueInterface)

	// This utility is for testing structs !
	if actualType.Kind() != reflect.Struct {
		return assert.Fail(t, "Non struct", "Tried to check an expected struct against a non struct object, consider fixing your test for field {%s}", parentFieldName)
	}

	var success = true

	// We browse all expected fields and test if they correspond to the actual struct or not
	for fieldKey, expectedFieldInterface := range e {
		_, structFieldFound := actualType.FieldByName(fieldKey)
		fieldName := formatting.FormatFieldName(parentFieldName, fieldKey)

		if !structFieldFound {
			success = assert.Fail(t,
				"Field not found", "Expected field not found : {%s}",
				fieldName,
			) && success
			continue
		}

		actualFieldValue := actualValue.FieldByName(fieldKey)
		actualFieldInterface := actualFieldValue.Interface()

		switch castedExpectedField := expectedFieldInterface.(type) {
		case FieldBehaviour:
			success = castedExpectedField.CheckField(
				t,
				actualFieldInterface,
				fieldName,
			) && success
		default:
			success = assert.Equalf(
				t,
				expectedFieldInterface,
				actualFieldInterface,
				"Check of Equality for field {%s}",
				fieldName,
			) && success
		}
	}

	// If some fields were not reviewed, we fail and tell what fields were missing
	allFieldsReviewed := len(e) == actualType.NumField()
	if !allFieldsReviewed {
		missingFields := make([]string, 0)
		for i := 0; i < actualType.NumField(); i++ {
			name := actualType.Field(i).Name
			_, found := e[name]
			if !found {
				missingFields = append(missingFields, name)
			}
		}

		precision := ""
		if parentFieldName != "" {
			precision = fmt.Sprintf("of field {%s} ", parentFieldName)
		}
		if len(missingFields) != 0 {
			missingFieldsString := ""
			for i, s := range missingFields {
				if i != 0 {
					missingFieldsString += ", "
				}
				missingFieldsString += fmt.Sprintf("{%s}", s)
			}
			success = assert.Fail(
				t,
				"Fields are missing",
				"Some fields %sare missing : %s", precision,
				missingFieldsString,
			) && success
		}

	}

	return success
}
