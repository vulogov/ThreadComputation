package ThreadComputation

import (
  "fmt"
  "errors"
  "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
)


func getMatrixFromStack(l *TCExecListener, q *deque.Deque) (*mat.Dense, error) {
  if l.TC.Ready() {
    e := l.TC.Get()
    switch e.(type) {
    case *mat.Dense:
      return e.(*mat.Dense), nil
    }
  }
  return nil, errors.New("There is no matrix value on the stack")
}

func TCMatSetFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var x interface{}
  var y interface{}
  var v interface{}

  e, err := getMatrixFromStack(l, q)
  if err != nil {
    return nil, err
  }
  if q.Len() == 3 {
    x = q.PopFront()
    y = q.PopFront()
    v = q.PopFront()
  } else {
    return nil, errors.New("matrix.Set[] must have three arguments: x, y, value")
  }
  switch x.(type) {
  case int64:
    switch y.(type) {
    case int64:
      switch v.(type) {
      case int64:
        e.Set(int(x.(int64)), int(y.(int64)), float64(v.(int64)))
        return e, nil
      case float64:
        e.Set(int(x.(int64)), int(y.(int64)), v.(float64))
        return e, nil
      }
    }
  }
  return nil, errors.New("matrix.Set[] can not determine a proper context")
}

func TCMatGetFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var x interface{}
  var y interface{}

  e, err := getMatrixFromStack(l, q)
  if err != nil {
    return nil, err
  }
  if q.Len() == 2 {
    x = q.PopFront()
    y = q.PopFront()
  } else {
    return nil, errors.New("matrix.Get[] must have two arguments: x, y, value")
  }
  switch x.(type) {
  case int64:
    switch y.(type) {
    case int64:
      ReturnToStackFunction(l, "matrix.Get", e)
      return e.At(int(x.(int64)),int(y.(int64))), nil
    }
  }
  return nil, errors.New("matrix.Get[] can not determine a proper context")
}

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
  SetFunction("matrix.Set", TCMatSetFunction)
  SetFunction("matrix.Get", TCMatGetFunction)
  SetFunction("M", TCNewMatFunction)
  SetFunction("m", TCNewEmptyMatFunction)
}
