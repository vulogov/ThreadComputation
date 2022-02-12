package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/gammazero/deque"
)

func NewStackFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  l.TC.Res.Add()
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      l.TC.Res.Set(e)
    }
  }
  return nil, nil
}

func ReturnToStackFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Res.GLen() < 2 {
    return nil, nil
  }
  if q.Len() == 0 {
    if l.TC.Res.Len() == 0 {
      return nil, nil
    }
    e, err := l.TC.Res.Take()
    if err != nil {
      return nil, err
    }
    l.TC.Res.Left()
    l.TC.Res.Set(e)
    l.TC.Res.Right()
  } else {
    l.TC.Res.Left()
    for q.Len() > 0 {
      e := q.PopFront()
      l.TC.Res.Set(e)
    }
    l.TC.Res.Right()
  }
  return nil, nil
}

func StackRotationFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var x []int
  if q.Len() == 0 {
    x = append(x, 1)
  } else {
    for q.Len() > 0 {
      ex := q.PopFront()
      switch ex.(type) {
      case int64:
        x = append(x, int(ex.(int64)))
      case float64:
        x = append(x, int(ex.(float64)))
      }
    }
  }
  for i := 0; i < len(x); i++ {
    for j := 0; j < x[i]; j++ {
      switch l.TC.CurrentFunctionName() {
      case "<-":
        l.TC.Res.Left()
      case "->":
        l.TC.Res.Right()
      case ">>":
        l.TC.Res.CRight()
      case "<<":
        l.TC.Res.CLeft()
      default:
        return nil, errors.New(fmt.Sprintf("Unknown stack operator: %v", l.TC.CurrentFunctionName()))
      }
    }
  }
  return nil, nil
}

func initStdlibStack() {
  SetFunction("|", NewStackFunction)
  SetFunction("_", ReturnToStackFunction)
  SetFunction("<-", StackRotationFunction)
  SetFunction("->", StackRotationFunction)
  SetFunction(">>", StackRotationFunction)
  SetFunction("<<", StackRotationFunction)
}
