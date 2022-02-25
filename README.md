[![Go](https://github.com/vulogov/ThreadComputation/actions/workflows/go.yml/badge.svg)](https://github.com/vulogov/ThreadComputation/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/vulogov/ThreadComputation.svg)](https://pkg.go.dev/github.com/vulogov/ThreadComputation)

# ThreadComputation

ThreadComputation is a programmatic module implementing core interpreter and virtual machine for programming language BUND. Being a GoLang module, ThreadComputation can be embedded in your project, giving you ability to add Domain Specific Language to your application.

## What is the BUND language ?

BUND language is interpreted, dynamically typed, functional and stack-based language, implemented in Go, embeddable a extendable with GoLang. BUND build around two-dimentional stack engine acting like a storage for the functions call modeled with idea of threading macros in mind. What is the main features of the BUND:

- Dynamic typing, with support for boolean, string, float, integer, set, map, matrix data types;
- Transparent (or striving to be transparent) approach to operators that work seamlessly with all supported types;
- URI-based access to external data;
- Lambda and named functions;
- Advanced operations with stack;
- Named stacks supported by two-dimentional stack VM;
- Very easy to write new functions in Go or in BUND;
- On-the-fly function redefine;
- Native support for map and json keys query;
- Prefix and postfix notations;
- and more ...

## Show me 'Hello World!' now !

The famous HelloWorld program will look like this:

```
print['Hello World!']
```
In this example, we are calling function print and pass the string as a function argument. Alternatively, you can execute that code:

```
'Hello world!' print
```
And in this example, you do not pass the string as argument and function print will take data from the stack. Or you can do something like that:

```
+['Hello ' 'world!'] `print attr !
```
Here, first, you are creating string using concatenation, then create function reference, then dynamically assign arguments to the that reference than execute that reference. As you see, possibilitis of how you can greet the world are endless.
 

## How to use ThreadComputation module

ThreadComputation module is hosted on GitHub, fully tested with automatic Actions call. You are welcome to [fork and contribute](https://github.com/vulogov/ThreadComputation) new functions and features of the BUND.

### Installation

```
go get github.com/vulogov/ThreadComputation
```

or you can check out the module and run

```
make pre; make
```

If you are planning to change BUND syntax, you must have ANTLR4 for Go installed.

```
make rebuild
```

will rebuild ANTLR4 code.

### Use from inside Go code

```golang
import "github.com/vulogov/ThreadComputation"
```

After you imported the module, you have to create TC instance

```golang
tc := ThreadComputation.Init()
```

This function call will create a BUND VM and initialize all structures for VM. Initial stack will be created for you.

```golang
tc = tc.Eval("BUND code goes here")
```

This call evaluates and executed a BUND code in created VM instance.

```golang
if tc.Errors() != 0 {
  log.Fatalf(tc.Error())
}
```

- Function tc.Errors() of BUND VM will return a number of errors from lexer, parser and run-time;
- Function tc.Error() of BUND VM will return last error message.

```golang
if tc.Ready() {
  res := tc.Get()
}
```

This function will test if stack have any value to return and return that value as an interface{}.


## Key concepts of the BUND language

- Basics of the [BUND stack engine](Documentation/BUND_stack.md)
- Work with [data](Documentation/BUND_data.md) in BUND
- BUND [functions](Documentation/BUND_fun.md) are the principal elements of the language

## To the source !

Information about [ThreadComputation package](https://pkg.go.dev/github.com/vulogov/ThreadComputation) is available on [https://pkg.go.dev](https://pkg.go.dev/github.com/vulogov/ThreadComputation)

Source code available on GitHub: [https://github.com/vulogov/ThreadComputation](https://github.com/vulogov/ThreadComputation)

## Get support

[Here, you can open ticket](https://github.com/vulogov/ThreadComputation/issues).
