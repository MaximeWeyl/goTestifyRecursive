# goRecursiveAssert

Extending https://github.com/stretchr/testify.git with recursive capabilities.

## Why

### The problem

If you've into testing, you surely wrote a lot of such code :

```go
package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MyStruct struct {
	A string
	B string
}

func MyFunction() MyStruct {
	return MyStruct{A: "foo", B: "bar"}
}

func TestMyFunction(t *testing.T) {
    res := MyFunction()
    assert.Equal(t, "foo", res.A)
    assert.Equal(t, "bar", res.B)
}
```

While this is a perfectly working and useful test, here is one problem : if your struct 
*MyStruct* changes in a next release, you may forget to update your test *TestMyFunction*.
This means you may : Not check the value of a new field.

### Solution provided

Enters goRecursiveAssert :

```go
package main

import (
	rec "github.com/MaximeWeyl/goTestifyRecursive"
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"testing"
)

type MyStruct struct {
	A string
	B string
}

func MyFunction() MyStruct {
	return MyStruct{A: "foo", B: "bar"}
}

func TestMyFunction(t *testing.T) {
    res := MyFunction()
    rec.AssertRecursive(t, res, bh.ExpectedStruct{
    	"A": "foo",
    	"B": "bar",
    })
}
```

If, later, you add a new Field to your struct, the test will fail. This will alert you that
you probably want to update your test too.


## Error handling

When a field is missing in the expected or actual value, you will get a nice error
that you should understand easily :

```
        	Error:      	Fields are missing
        	Test:       	TestMyFunction
        	Messages:   	Some fields are missing : {C}
```

## But you said "recursive"

Yes and here is why. This new syntax allow you to write nested assertions, just like this :

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

func MyFunction() MyStruct {
	return MyStruct{
		A: "foo", 
		B: []MyStruct{{
                A: "bar1",
                B: nil,
		    },
			{
				A: "bar2",
				B: []MyStruct{},
			},
		},
	}
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

The errors will also try to show you were your error is :

```
        	Error:      	Field not found
        	Test:       	TestMyFunction
        	Messages:   	Expected field not found : {B[1].C}
```

This way of writing tests prevents you from forgetting to update tests that need to be.

## But I do not want to check for this new field

No problem, the idea of this lib is to let you know that you may want to check for a new field.
But, if you decided that this new field should not be tested in this test, you can tell the lib
to ignore it, you're the one that know your code best after all. Just use the *IgnoredField* struct.


```go
package main

import (
	rec "github.com/MaximeWeyl/goTestifyRecursive"
	bh "github.com/MaximeWeyl/goTestifyRecursive/behaviours"
	"testing"
)

type MyStruct struct {
	A string
	B string
	NewField int
}

func MyFunction() MyStruct {
	return MyStruct{A: "foo", B: "bar"}
}

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

You have seen several behaviours. The first one is known as the *default behaviour* :
It is what happen when you pass an argument that does not implement the special interface
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

## Require instead of Assert

Like in testify, you can stop your test immediately if the assertion fails.
Simply use the function *RequireRecursive* instead of *AssertRecursive*


## Contributing

Contributions are welcome, especially for writing new behaviours.

