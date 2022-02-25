# Functions, Functions everywhere !

 Concept of functions in BUND language is not that different from other functional languages, but it does have some specifics. In general, function is a sequence of BUND terms, taking arguments from argument list or the stack that can return value to the stack or to list of the arguments or other function. There are three types of functions in BUND:

 - Internal functions. Those functions implemented in GoLang and they are the part of your application, or BUND implementation. There are number of internal functions defined in ThreadComputation module, which is the implementation of the core of BUND language.
 - Named user functions. User-defined functions stored in the local context. Those functions are executed without creating a new stack.
 - Lambda functions, or anonymous functions. Those functions created in the local context and could be referred only by using "function reference".

## How to call a function ?

For internal and named user functions, function call is simple, you are specify a function name, with optional space separated arguments enclosed in square brackets. Example:

```
stack[ 1 2 3 ]
```
When you are place the items to the stack, the stack is LIFO, so when you place the items "1 2 3" on the stack, they are in reverse and on the stack you have "3 2 1". But in attributes, they are direct, so the arguments _[ 1 2 3 ]_ are ordered as they are: "1 2 3".

## How to define internal functions ?

First, you have to create a function in GoLang. Let's look at the most basic function, which will return the value to the BUND context.

```golang
func TheUltimateAnswerFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.PopFront(), nil
  }
  return 42, nil
}
```

First what we shall pay attention on is parameters of the functions:

- _l *TCExecListener_ is a reference on the BUND interpreter. All VM functions and variables are accessible through this reference. Here are the few examples:
  - _l.TC_ - the reference to the BUND VM;
  - _l.TC.Res_ - this is your "stack-of-stacks";
  - _l.TC.Attrs_ - "stack-of-stacks" holding functions attributes;
  - _l.TC.FNStack_ - functions calls stack.
- _q *deque.Deque_ - queue of the arguments passed to the function call

Next what we see, that the function is testing the size of the queue holding arguments and if any arguments is passed, it will take the top one by using _q.PopFront()_ and return it to the VM. Otherwise, the number _42_ is returned. The second return value is reference to the error. If error occured inside the function, you shall return that error and VM will stop execution and report that error.

Next, we must make this function part of the BUND internal functions:

```golang
SetFunction("TheAnswer", TheUltimateAnswerFunction)
```

Function _SetFunction()_ takes two parameters. First, the name of the function as it will be known in the BUND. Second, the reference to the function.

## User functions are simple.

User can define a function that can be used within an instance of BUND VM. User function defined as a sequence of BUND terms, associated with name. The syntax of function definition is a function name, prefixed with _@_, and BUND terms enclosed between _[_ and _]_. After you defined a function, you can call it in the same way as internal function as name of the function and optional parameters enclosed between _[_ and _]_. Example:

```
@answer[42] answer
```
In this example we are defining use3r function called _answer_ that will be executed within current VM context and stack and the only term in this function is a number 42. So, the outcome of this function execution will be the number 42 placed on top of the stack or arguments list.

## Lambda or anonymous functions

Lambda functions could be created using following syntax: *lambda\\ BUND terms \\*
