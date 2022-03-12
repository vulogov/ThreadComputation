package ThreadComputation

import (
  "github.com/gammazero/deque"
)



func TCLenFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.Len(), nil
  } else {
    if l.TC.Ready() {
      return l.TC.Res.Len(), nil
    }
  }
  return 0, nil
}

func ToStackFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, nil
  }
  for q.Len() > 0 {
    e:= q.PopFront()
    ReturnFromFunction(l, name, e)
  }
  return nil, nil
}

func init() {
  SetCommand("len", TCLenFunction)
  SetCommand("stack", ToStackFunction)
}
