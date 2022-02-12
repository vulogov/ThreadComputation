package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
)

func PrintFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      fmt.Println(e)
    }

  } else {
    if l.TC.Ready() {
      data_out := l.TC.Get()
      fmt.Println(data_out)
      return data_out, nil
    }
  }
  return nil, nil
}

func PrintAllFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      fmt.Println(e)
    }

  } else {
    if l.TC.Ready() {
      for l.TC.Ready() {
        data_out := l.TC.Get()
        fmt.Println(data_out)
      }
      return nil, nil
    }
  }
  return nil, nil
}

func ToStackFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, nil
  }
  for q.Len() > 0 {
    e:= q.PopFront()
    l.TC.Res.Set(e)
  }
  return nil, nil
}

func LenFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.Len(), nil
  } else {
    if l.TC.Ready() {
      return l.TC.Res.Len(), nil
    }
  }
  return 0, nil
}

func ClrFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    for l.TC.Ready() {
      l.TC.Get()
    }
  }
  return nil, nil
}

func DropFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    l.TC.Res.PopFront()
  }
  return nil, nil
}

func DupFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    e := l.TC.Get()
    l.TC.Res.Set(e)
    l.TC.Res.Set(e)
  }
  return nil, nil
}

func TheUltimateAnswerFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.PopFront(), nil
  }
  return 42, nil
}

func initStdlibGenerics() {
  SetFunction("print", PrintFunction)
  SetFunction("printAll", PrintAllFunction)
  SetFunction("stack", ToStackFunction)
  SetFunction("len", LenFunction)
  SetFunction("drop", DropFunction)
  SetFunction(",", DropFunction)
  SetFunction("dup", DupFunction)
  SetFunction("^", DupFunction)
  SetFunction("clr", ClrFunction)
  SetFunction("TheAnswer", TheUltimateAnswerFunction)
}
