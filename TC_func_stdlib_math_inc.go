package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCIncInt(v interface{}) interface{} {
  return v.(int64) + 1
}

func TCIncFloat(v interface{}) interface{} {
  return v.(float64) + 1.0
}

func TCIncMatrix(v interface{}) interface{} {
  return tcMatrixMathFloat("+", v.(*TCMatrix), 1.0)
}

func TCIncNumbers(v interface{}) interface{} {
  return tcNumbersMathFloat("+", v.(*TCNumbers), 1.0)
}

func TCIncFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("inc", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}


func init() {
  RegisterFunctionCallback("inc", Int, TCIncInt)
  RegisterFunctionCallback("inc", Float, TCIncFloat)
  RegisterFunctionCallback("inc", Matrix, TCIncMatrix)
  RegisterFunctionCallback("inc", Numbers, TCIncNumbers)
  SetFunction("++", TCIncFunction)
}
