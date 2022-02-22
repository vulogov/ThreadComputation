package ThreadComputation

import (
  "fmt"
  "errors"
  "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
)




func TCNewMatFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var re1 int
  var re2 int
  e1, e2, err := getTwoValues(l, q)
  if err != nil {
    return nil, err
  }
  switch e2.(type) {
  case int64:
    re2 = int(e2.(int64))
  case float64:
    re2 = int(e2.(float64))
  default:
    return nil, errors.New(fmt.Sprintf("invalid parameter for M[] e2=%T", e2))
  }
  switch e1.(type) {
  case int64:
    re1 = int(e1.(int64))
  case float64:
    re1 = int(e1.(float64))
  default:
    return nil, errors.New(fmt.Sprintf("invalid parameter for M[] e1=%T", e2))
  }
  return mat.NewDense(re1, re2, nil), nil
}

func TCNewEmptyMatFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var r mat.Dense
  res, err := TCNewMatFunction(l, q)
  if err != nil {
    return &r, nil
  }
  return res, nil
}

func initStdlibMat() {
  SetFunction("matrix", TCNewMatFunction)
  SetFunction("M", TCNewMatFunction)
  SetFunction("m", TCNewEmptyMatFunction)
}
