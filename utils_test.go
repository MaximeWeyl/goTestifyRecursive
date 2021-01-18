package testifyRecursive

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

type MockedT struct {
	errors []string

	assertionsBegan   bool
	errorsMapAsserted []bool
	alreadyAsserted   int

	failedWithFailNow bool
}

func (mot *MockedT) FailNow() {
	mot.failedWithFailNow = true
}

func (mot *MockedT) Errorf(format string, args ...interface{}) {
	if mot.assertionsBegan {
		panic("Please don't run tests after assertions have begun")
	}

	mot.errors = append(mot.errors, fmt.Sprintf(format, args...))
}

func (mot *MockedT) AssertFailed(t *testing.T, withFailNow bool) bool {
	mot.initAssertionsIfNeeded()

	if len(mot.errors) == 0 {
		return assert.Fail(t, "Did not fail", "We expected this test to have fail")
	}

	if withFailNow && !mot.failedWithFailNow {
		return assert.Fail(t, "Did fail but without FailNow", "We expected this test to have fail with FailNow")
	}

	if !withFailNow && mot.failedWithFailNow {
		return assert.Fail(t, "Did fail but with FailNow", "We expected this test to have fail without FailNow")
	}

	return true
}

func (mot *MockedT) AssertPassed(t *testing.T) bool {
	mot.initAssertionsIfNeeded()

	if len(mot.errors) != 0 {
		return assert.Fail(t, "Did not pass", "We expected this test to have passed")
	}

	if mot.failedWithFailNow {
		return assert.Fail(t, "FailNow was called", "We expected FailNow not to be called")
	}

	return true
}

func (mot *MockedT) AssertAllErrorsWereAsserted(t *testing.T) bool {
	if !mot.assertionsBegan {
		return assert.Fail(t, "No assertions happened")
	}

	if mot.alreadyAsserted != len(mot.errors) {
		var remainingErrors []string
		for i, asserted := range mot.errorsMapAsserted {
			if asserted {
				continue
			}
			errorSplit := strings.Split(mot.errors[i], "\n")
			remainingErrors = append(
				remainingErrors,
				"\n"+errorSplit[len(errorSplit)-2]+"\n",
			)
		}

		return assert.Failf(
			t,
			"Errors remaining",
			"There are %d errors left to be asserted (on a total of %d) :\n%s",
			len(mot.errors)-mot.alreadyAsserted,
			len(mot.errors),
			remainingErrors,
		)
	}

	return true
}

func (mot *MockedT) AssertRemainingErrorMessageAndDiscardIt(t *testing.T, expectedMessage string) bool {
	mot.initAssertionsIfNeeded()

	for i, err := range mot.errors {
		if mot.errorsMapAsserted[i] {
			continue
		}

		var re = regexp.MustCompile(fmt.Sprintf(`(?m)\n\s+Messages:\s+%s\n`, regexp.QuoteMeta(expectedMessage)))
		found := re.MatchString(err)

		if found {
			mot.errorsMapAsserted[i] = true
			mot.alreadyAsserted++
			return true
		}
	}

	return assert.Fail(
		t,
		"Not found",
		"Message '%s' was not found in any of the %d remaining errors (over %d total errors)",
		expectedMessage,
		len(mot.errors)-mot.alreadyAsserted,
		len(mot.errors),
	)
}

func (mot *MockedT) initAssertionsIfNeeded() {
	if !mot.assertionsBegan {
		mot.errorsMapAsserted = make([]bool, len(mot.errors))
		mot.assertionsBegan = true
	}
}
