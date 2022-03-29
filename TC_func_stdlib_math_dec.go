package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCDecInt(v interface{}) interface{} {
  return v.(int64) - 1
}

func TCDecFloat(v interface{}) interface{} {
  return v.(float64) - 1.0
}

func TCDecMatrix(v interface{}) interface{} {
  return tcMatrixMathFloat("-", v.(*TCMatrix), 1.0)
}

func TCDecNumbers(v interface{}) interface{} {
  return tcNumbersMathFloat("-", v.(*TCNumbers), 1.0)
}

func TCDecFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("dec", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  RegisterFunctionCallback("dec", Int, TCDecInt)
  RegisterFunctionCallback("dec", Float, TCDecFloat)
  RegisterFunctionCallback("dec", Matrix, TCDecMatrix)
  RegisterFunctionCallback("dec", Numbers, TCDecNumbers)
  SetFunction("--", TCDecFunction)
}
