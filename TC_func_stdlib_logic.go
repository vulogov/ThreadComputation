package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
)

func TCNotFunction(x interface{}) interface{} {
  switch x.(type) {
  case bool:
    return ! x.(bool)
  }
  return nil
}

func TCNotValueFunction(x interface{}) interface{} {
  switch x.(type) {
  case *TCValue:
    switch x.(*TCValue).Value.(type) {
    case bool:
      res := TCNotFunction(x.(*TCValue).Value.(bool))
      if res != nil {
        return MakeValue(res.(bool), x.(*TCValue).P, 0)
      }
    }
  }
  return nil
}

func TCSimpleLogicFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 1 {
    return true, nil
  }
  x := q.PopFront()
  for q.Len() > 0 {
    y := q.PopFront()
    fun := GetCompareCallback(x, y)
    if fun == nil {
      return nil, l.TC.MakeError(fmt.Sprintf("Can not get comparator for %T %v %T", x, name, y))
    }
    res := fun(name, x, y)
    if res == nil {
      return nil, l.TC.MakeError(fmt.Sprintf("%T %v %T returned nil", x, name, y))
    }
    ReturnFromFunction(l, name, res)
  }
  return nil, nil
}

func TCNotLogicFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    fun := GetFunctionCallback("not", e)
    if fun == nil {
      return nil, l.TC.MakeError("Invalid type for not[] function")
    }
    res := fun(e)
    if res == nil {
      return nil, l.TC.MakeError("not[] function returned nil")
    }
    ReturnFromFunction(l, "not", res)
  }
  return nil, nil
}

func init() {
  RegisterFunctionCallback("not", Bool, TCNotFunction)
  RegisterFunctionCallback("not", Value, TCNotValueFunction)
  SetFunction("and", TCSimpleLogicFunction)
  SetFunction("or", TCSimpleLogicFunction)
  SetFunction("not", TCNotLogicFunction)
}
