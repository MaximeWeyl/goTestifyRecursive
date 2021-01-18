package checkStruct

import (
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"testing"
)

type TE struct {
	E1 string
	E2 int
}

type TD struct {
	D1 string
	D2 int
	D3 []TE
}

type TA struct {
	A string
	B bool
	C int
	D TD
}

func getSample() TA {
	return TA{
		A: "string A",
		B: true,
		C: 42,
		D: TD{
			D1: "string D1",
			D2: 52,
			D3: []TE{
				{
					E1: "string E1-0",
					E2: 62,
				},
				{
					E1: "string E1-1",
					E2: 72,
				},
			},
		},
	}
}

func TestEquality(t *testing.T) {

	a := getSample()
	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": false,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {B}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": false,
			"C": 43,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {B}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {C}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": false,
			"C": 43,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 60,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {B}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {C}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {D.D2}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": false,
			"C": 43,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 60,
				"D3": bh.ExpectedSlice{
					{
						"E1": "error",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 100,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {B}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {C}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {D.D2}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {D.D3[0].E1}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {D.D3[1].E2}")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestMissingFields(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Some fields are missing : {A}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Some fields of field {D} are missing : {D2}, {D3}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D2": 52,
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Some fields are missing : {A}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Some fields of field {D} are missing : {D1}, {D3}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"C": 42,
			"D": bh.ExpectedStruct{},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Some fields are missing : {A}, {B}")
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Some fields of field {D} are missing : {D1}, {D2}, {D3}")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestFieldNotFound(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
				"UnknownField": "Some value",
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Expected field not found : {D.UnknownField}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1":           "string E1-1",
						"E2":           72,
						"UnknownField": "Some value",
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Expected field not found : {D.D3[1].UnknownField}")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestSliceSizeDoesNotMatch(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "For field {D.D3}, The actual slice has a len of 2, while a size of 1 was expected")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestWithNonStructShouldFail(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, "test", bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(
			t,
			"Tried to check an expected struct against a non struct object, "+
				"consider fixing your test for field {}",
		)
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": bh.ExpectedStruct{"E": "Error"},
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(
			t,
			"Tried to check an expected struct against a non struct object, "+
				"consider fixing your test for field {A}",
		)
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": bh.ExpectedStruct{"E": "error"},
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(
			t,
			"Tried to check an expected struct against a non struct object, "+
				"consider fixing your test for field {D.D1}",
		)
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": bh.ExpectedStruct{"E": "error"},
						"E2": 62,
					}, {
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(
			t,
			"Tried to check an expected struct against a non struct object, "+
				"consider fixing your test for field {D.D3[0].E1}",
		)
		tInner.AssertAllErrorsWereAsserted(t)
	}

}

func TestWithNonSliceShouldFail(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": bh.ExpectedSlice{{"E": "error"}},
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "A slice was expected for field {D.D2}. We found : 52")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestIgnoreField(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": bh.IgnoredField{},
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": bh.IgnoredField{},
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": bh.IgnoredField{},
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": bh.IgnoredField{},
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": bh.IgnoredField{},
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "error",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": bh.IgnoredField{},
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {D.D3[0].E1}")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestNotEmptyField(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": bh.NotEmptyField{},
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}

		b := struct {
			A []string
		}{
			A: []string{},
		}

		AssertRecursive(tInner, b, bh.ExpectedStruct{
			"A": bh.NotEmptyField{},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Field {A} should not be empty, but was")
		tInner.AssertAllErrorsWereAsserted(t)
	}
}

func TestSameValueField(t *testing.T) {

	a := getSample()

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": int64(62),
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {D.D3[0].E2}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

	{
		tInner := &MockedT{}
		AssertRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": bh.SameValueField{Expected: int64(62)},
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}
}

func TestZeroValueField(t *testing.T) {

	b := struct {
		A int
	}{}

	c := struct {
		A int
	}{A: 12}

	{
		tInner := &MockedT{}

		AssertRecursive(tInner, b, bh.ExpectedStruct{
			"A": bh.ZeroValueField{}})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}

		AssertRecursive(tInner, c, bh.ExpectedStruct{
			"A": bh.ZeroValueField{}})
		// Test should have pass
		tInner.AssertFailed(t, false)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Field {A} should have zero value but has value 12")
		tInner.AssertAllErrorsWereAsserted(t)
	}

}

func TestRequire(t *testing.T) {

	a := getSample()
	{
		tInner := &MockedT{}
		RequireRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": true,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertPassed(t)
	}

	{
		tInner := &MockedT{}
		RequireRecursive(tInner, a, bh.ExpectedStruct{
			"A": "string A",
			"B": false,
			"C": 42,
			"D": bh.ExpectedStruct{
				"D1": "string D1",
				"D2": 52,
				"D3": bh.ExpectedSlice{
					{
						"E1": "string E1-0",
						"E2": 62,
					},
					{
						"E1": "string E1-1",
						"E2": 72,
					},
				},
			},
		})
		// Test should have pass
		tInner.AssertFailed(t, true)
		tInner.AssertRemainingErrorMessageAndDiscardIt(t, "Check of Equality for field {B}")
		tInner.AssertAllErrorsWereAsserted(t)
	}

}
