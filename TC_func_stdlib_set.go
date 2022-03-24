package ThreadComputation

import (
  "github.com/deckarep/golang-set"
  "github.com/gammazero/deque"
)


func TCSetFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := mapset.NewSet()
  for q.Len() > 0 {
    e := q.PopFront()
    res.Add(e)
  }
  return res, nil
}

func init() {
  SetCommand("set", TCSetFunction)
}
