# ThreadComputation

ThreadComputation is a GoLang module, that been created with hope to simplify creation of the DSL languages based on thread computation. What is thread computing ? Essentially, it is an idea, that been explored in various flavors of Lisp, including Clojure, called threading macros. The idea behind threading macros is simple, imagine, that we are calling functions, passing return of one function as parameter to another, in a way that we do have a conveyer or thread function call. Example:

```
c( b( a() ) )
```
In this example, function a is called, passing return as parameter of function b, which then called and passed it's return as a parameter to function c which is returning the result of computation.
