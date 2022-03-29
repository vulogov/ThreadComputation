package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func TCListFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := new(TCList)
  for q.Len() > 0 {
    e := q.PopFront()
    log.Debugf("Pushing %T to list", e)
    res.Q.PushBack(e)
  }
  log.Debugf("%v elements has been added to the list", res.Q.Len())
  return res, nil
}

func init() {
  SetCommand("list", TCListFunction)
}
