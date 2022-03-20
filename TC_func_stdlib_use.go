package ThreadComputation

import (
  "github.com/gammazero/deque"
)


func TCUseFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    e := l.TC.Get()
    switch e.(type) {
    case string:
      data, err := readVfsFile(e.(string))
      if err != nil {
        return nil, err
      }
      l.TC.Eval(data)
    default:
      l.TC.Set(e)
    }
  }
  if q.Len() > 0 {
    f := q.PopFront()
    switch f.(type) {
    case string:
      data, err := readVfsFile(f.(string))
      if err != nil {
        return nil, err
      }
      l.TC.Eval(data)
    }
  }
  return nil, nil
}

func init() {
  SetCommand("use", TCUseFunction)
}
