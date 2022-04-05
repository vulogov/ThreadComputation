package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCIOTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    res := MakeIO(l, q)
    return res, nil
  }
  return MakeType(IO), nil
}

func init() {
  SetCommand("IO", TCIOTypeFunction)
}
