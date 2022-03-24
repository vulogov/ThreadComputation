package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func TCExecuteFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    efun := GetExecuteCallback(e)
    if efun != nil {
      log.Debugf("Executing for data type: %T", e)
      res := efun(l, e)
      if res != nil {
        ReturnFromFunction(l, "!", res)
      }
    }
  }
  return nil, nil
}

func init() {
  SetFunction("!", TCExecuteFunction)
}
