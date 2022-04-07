package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func TCStackBackSendFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    log.Debugf("\";\" Backsending the data: %T", e)
    ReturnFromFunction(l, ";", e)
  }
  return nil, nil
}

func init() {
  SetFunction(";", TCStackBackSendFunction)
}
