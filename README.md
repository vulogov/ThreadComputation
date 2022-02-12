# ThreadComputation

ThreadComputation is a GoLang module, that been created with hope to simplify creation of the DSL languages based on thread computation. What is thread computing ? Essentially, it is an idea, that been explored in various flavors of Lisp, including Clojure, called threading macros. The idea behind threading macros is simple, imagine, that we are calling functions, passing return of one function as parameter to another, in a way that we do have a conveyer or thread function call. Example:

```
c( b( a() ) )
```
In this example, function a is called, passing return as parameter of function b, which then called and passed it's return as a parameter to function c which is returning the result of computation. While, the idea of that "thread" or "conveyer" call is quite simple, the code is not quite readable and you do need to unfold the sequence of calls in reverse order in order to understand of what is calling what. Wouldn't it be simple, if we just have a syntax similar to this:

```
a() b() c()
```

Where a() is called and passing return value to b() and b() is passing return value to c() and return value of c() will be an outcome of computation. This syntax will be much easier to comprehend and therefore more error-proof. Lisp-like languages offering some solution for that problem but I've decided to step forward in different direction.

## What is exactly ThreadComputation module ?

Imagine, that you are in despeate need of embeddable, simple and extendable DSL language for your GoLang application. Something you can use as your application query language or expression evaluation. Don't have to be a Turing complete, but easy to understand by your users, powerful and easily extendable. And ThreadComputation module is trying to fill that niche. Let me show you the code.

```
"Hello world !" print
```

This snippet of the code will print "Hello world!" on standard output. But wait, I've mentioned threading macros and this is looks like another implementaion of FORT language. Well, this is because ThreadComputation DLS doesn't making a difference, between data and a code. Just like a Lisp. And "execution" of the string, which is "Hello world!" will change the VM context. And of course, _print_ is a function, which is if not given any paramters will try to get data from VM context. And postfix notation is supported too. Example:

```
print["Hello world!"]
```

here, we are passing parameters to the function as an argument. Arguments are inclosed between _"["_ and _"]"_ and space separated. Example:

```
printAll["Hello world!" "and sorry, Dave"]
```

## The Stack

Classic threading macros are passing the values as a first argument of the next function, but ThreadComputation is doing this through the reference to a context, and context stores the data in the stack, the FIFO structure.

## Show me the Go code

Let's do some RPN computations:

```(language=Go)
import "github.com/vulogov/ThreadComputation"

...

tc := ThreadComputation.Init()
tc.Eval("2 2 + 6 -")
fmt.Println(tc.Get())
```

- First, we are importing the the module.
- Next, by calling _ThreadComputation.Init()_, we are creating VM instance.
- _tc.Eval()_ will evaluate the code.
- _tc.Get()_ will return the last return value of the last called function

There are more GoLang module functions, but those are most common ones.

How to extend the module ?

## How to extend ThreadComputation ?

ThreadComputation module by itself is not very useful. You have to add your own functions to it. Here is how. First, you define your function:

```(language=Go)
func TheUltimateAnswerFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.PopFront(), nil
  }
  return 42, nil
}
```

- This function takes two arguments: first is a reference to a evaluation context, where attribute _l.TC_ is a reference to execution context. Second, is a reference to optional parameters, as a queue.
- This function will return value and error. If everything fine, error will be nil. The value will be pushed to the front of the context stack. Result of execution of the previous function will be always stored in the front of the context stack.
- We are checking, if we are passing anything as an arguments to this function call, then we are returning the first argument after we popping this argument from the arguments queue. VM will place this value to the top of the stack.
- And if no arguments passed, we are returning 42.

Next, you shall registed this function in the list of global functions:

```(language=Go)
ThreadComputation.SetFunction("TheAnswer", TheUltimateAnswerFunction)
```

- First parameter is a name of the function as it will be known in DSL.
- Second, is a reference to your function implementation.

After that, you can use that function

```
TheAnswer
```

This call will place 42 into a top of the context stack, but if you pass an argument to the call, like that:

```
TheAnswer[3.14]
```

then value of the Pi will be returned to the stack.

##ThreadComputation language

What is the ThreadComputation language syntax ? It could be defined as combination of three types of terms:

- Data definition
- Variables declaration
- Functions calling.

### Data definition

ThreadComputation is dynamically-typed language and there are four data types currently defined

#### Boolean data type

Described as ether #TRUE or #FALSE in DSL. Represented by _bool_ data type on the GoLang level.

#### Integer data type

Any combination of the digits will be converted to the integer number, represented by Int64 on the GoLang level.

#### Float data type

Numbers".(dot)"Numbers will be converted to the Float value, represented by Float64 on GlLang level.

#### String data type

"\"(quote)"CHARACTERS"\"(quote)" will be converted to the string value, represented by the String in GoLang.

### Variables declaration

### Functions calling

There are number of the functions that already
