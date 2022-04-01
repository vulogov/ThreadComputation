package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func TCRotateRightFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  log.Debugf("Rotate list right: q=%v", q.Len())
  err := l.ExecuteSingleArgumentFunction("rotateright", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func TCRotateLeftFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  log.Debugf("Rotate list left: q=%v", q.Len())
  err := l.ExecuteSingleArgumentFunction("rotateleft", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  SetFunction("->>", TCRotateRightFunction)
  SetFunction("right", TCRotateRightFunction)
  SetFunction("<<-", TCRotateLeftFunction)
  SetFunction("left", TCRotateLeftFunction)
}
