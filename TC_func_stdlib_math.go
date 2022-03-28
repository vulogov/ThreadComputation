package ThreadComputation

import (
  "fmt"
  "errors"
  "math"
  "github.com/gammazero/deque"
)

func TCSimpleMathFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var acc interface{}
  if q.Len() == 0 {
    return math.NaN(), nil
  }
  acc = q.PopFront()
  if q.Len() == 0 {
    return acc, nil
  }
  for q.Len() > 0 {
    e := q.PopFront()
    fun := GetMathCallback(acc, e)
    if fun == nil {
      return nil, errors.New(fmt.Sprintf("Operator non-existing for %T %v %T", acc, name, e))
    }
    res := fun(name, acc, e)
    switch res.(type) {
    case *TCError:
      l.TC.RegisterError(res.(*TCError))
      return nil, res.(*TCError).Err
    }
    if res == nil {
      return nil, errors.New(fmt.Sprintf("%T %v %T returns nil", acc, name, e))
    }
    acc = res
  }
  return acc, nil
}


func init() {
  SetFunction("+", TCSimpleMathFunction)
  SetFunction("*", TCSimpleMathFunction)
  SetFunction("/", TCSimpleMathFunction)
  SetFunction("-", TCSimpleMathFunction)
}
