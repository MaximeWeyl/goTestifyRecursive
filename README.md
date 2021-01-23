# goTestifyRecursive


Extending https://github.com/stretchr/testify.git with recursive capabilities.

Tested for go 1.15, 1.14, 1.13.

[![Build Status](https://travis-ci.org/MaximeWeyl/goTestifyRecursive.svg?branch=master)](https://travis-ci.org/MaximeWeyl/goTestifyRecursive)
[![codecov](https://codecov.io/gh/MaximeWeyl/goTestifyRecursive/branch/master/graph/badge.svg?token=UFOL6XICXV)](https://codecov.io/gh/MaximeWeyl/goTestifyRecursive)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/MaximeWeyl/goTestifyRecursive)](https://pkg.go.dev/github.com/MaximeWeyl/goTestifyRecursive)


- [goTestifyRecursive](#gotestifyrecursive)
  * [Why](#why)
    + [The problem](#the-problem)
    + [Solution provided](#solution-provided)
  * [But you said "recursive"](#but-you-said-recursive)
  * [But I do not want to check for this new field](#but-i-do-not-want-to-check-for-this-new-field)
  * [Behaviours](#behaviours)
  * [Writing your own behaviours](#writing-your-own-behaviours)
  * [Using any testify function as behaviour](#using-any-testify-function-as-behaviour)
  * [Require instead of Assert](#require-instead-of-assert)
  * [Contributing](#contributing)

## Why

### The problem

If you test your GO code (and you should), you're probably
writting a lot of assert functions in order to test
every fields of some struct.


```go
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyFunction(t *testing.T) {
    res := MyFunction()
    
    assert.Equal(t, "foo", res.A)
    assert.Equal(t, "bar", res.B)
    // ^ And potentially many more asserts here
}
```

While this is a perfectly useful test, here is one problem : if your struct 
*MyStruct* changes in a next release, you may forget to update your test *TestMyFunction*. 
When you're adding a new field, forgetting to test its value
in your existing tests won't prevent them to pass.

### Solution provided

Enters goTestifyRecursive, which introduces the *AssertRecursive*
function. Here is how the test above is rewritten :

```go
package main

import (
	rec "github.com/MaximeWeyl/goTestifyRecursive"
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"testing"
)

func TestMyFunction(t *testing.T) {
    res := MyFunction()
    rec.AssertRecursive(t, res, bh.ExpectedStruct{
    	"A": "foo",
    	"B": "bar",
    })
}
```

If, later, you add a new Field to your struct, the test will fail,
letting you know which field is missing :

```
Error:      	Fields are missing
Test:       	TestMyFunction
Messages:   	Some fields are missing : {C}
```

## But you said "recursive"

Yes and here is why : This new syntax allow you
to write nested assertions, just like this :

```go
package main

import (
	rec "github.com/MaximeWeyl/goTestifyRecursive"
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"testing"
)

type MyStruct struct {
	A string
	B []MyStruct
}

func TestMyFunction(t *testing.T) {
    res := MyFunction()
    rec.AssertRecursive(t, res, bh.ExpectedStruct{
    	"A": "foo",
    	"B": bh.ExpectedSlice{
    		bh.ExpectedStruct{"A":"bar1", "B": []MyStruct(nil)},
    		bh.ExpectedStruct{"A":"bar2", "B": bh.ExpectedSlice{}},
        },
    })
}
```

The errors will also try to show you were your error is,
taking into account nested slices and structs :

```
Error:      	Field not found
Test:       	TestMyFunction
Messages:   	Expected field not found : {B[1].C}
```

This way of writing tests prevents you from forgetting to update tests that need to be.

## But I do not want to check for this new field

No problem, the idea of this lib is to let you know that you may want to check for a new field.
But if you decided that this new field should not be tested in this test, you can tell the lib
to ignore it : you're the one that know your code best after all. Just use the *IgnoredField* struct.


```go
package main

import (
	rec "github.com/MaximeWeyl/goTestifyRecursive"
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"testing"
)

func TestMyFunction(t *testing.T) {
    res := MyFunction()
    rec.AssertRecursive(t, res, bh.ExpectedStruct{
    	"A": "foo",
    	"B": "bar",
    	"NewField": bh.IgnoredField{},
    })
}
```



## Behaviours

In the previous sections
we have seen several behaviours. 
The first one is known as the *default behaviour* :
It is what happen when you pass an argument that 
does not implement the special interface
*FieldBehaviour*.

In this case, the assertion checks that the field actual value is **Equal** to the expected
value, with the help of testify's *Equal* function.

If you use an expected value that implements the *FieldBehaviour* interface, you will get
customized behaviours, like with *ExpectedStruct* or *ExpectedSlice* or *IgnoredField*.
This lib let you use several other behaviours, in the 
package *github.com/MaximeWeyl/goTestifyRecursive/behaviours*.


## Writing your own behaviours

Writing a new behaviour is as simple as implementing the *FieldBehaviour* interface.

```go
type FieldBehaviour interface {
CheckField(t assert.TestingT, actualValueInterface interface{}, parentFieldName string) bool
}
```

For instance, here is the code to write a behaviour that asserts that the field is not
empty, using the original testify lib (this example is officially in this lib's behaviours).


```go
package behaviours

import (
	"github.com/stretchr/testify/assert"
)

type NotEmptyField struct {
}

func (n NotEmptyField) CheckField(t assert.TestingT, actualValueInterface interface{}, fieldName string) bool {
	return assert.NotEmptyf(t, actualValueInterface, "Field {%s} should not be empty, but was", fieldName)
}

```

## Using any testify function as behaviour

You can use any function from testify package (or any other function
as long as it uses the same conventions).

This function must begin with a "t" parameter and return a bool
parameter (like any functions in the testify's assert package)

For this, you need to use the *Func* behaviour factory,
which takes the function (for instance, assert.Equal) 
as first parameter.

Any following parameter will be passed to the function.
The first "t" parameter should never be passed, the one from
*AssertRecursive* or *RequireRecursive* will be used.

In order to pass the currently tested value, use the 
custom parameter X. In order to pass the name of the 
current field, use the custom parameter F (this is useful
for printing custom error messages).

```go
package main

import (
	rec "github.com/MaximeWeyl/goTestifyRecursive"
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	a := struct {
		A int
	}{
		A: 12,
	}

	expected := 12
	rec.AssertRecursive(t, a, bh.ExpectedStruct{
		"A": bh.Func(assert.Equal, expected, bh.X, "Error for field %s : expected %d but got %d", bh.F, expected, bh.X),
	})
}

```


## Require instead of Assert

Like in testify, you can stop your test immediately if the assertion fails.
Simply use the function *RequireRecursive* instead of *AssertRecursive*


## Contributing

Contributions are welcome, especially for writing new behaviours.

