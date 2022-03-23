package ThreadComputation

import (
  "github.com/gammazero/deque"
)


func TCListFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var res TCList
  for q.Len() > 0 {
    e := q.PopFront()
    res.Q.PushBack(e)
  }
  return &res, nil
}

func init() {
  SetCommand("list", TCListFunction)
}
