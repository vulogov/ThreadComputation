package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func TCStackBackSendFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  p := l.TC.GetContext("p")
  for q.Len() > 0 {
    e := q.PopFront()
    log.Debugf("\";\" Backsending the data: %T", e)
    if p == nil {
      ReturnFromFunction(l, ";", e)
    } else {
      switch p.(type) {
      case float64:
        ReturnFromFunction(l, ";", MakeValue(e, float32(p.(float64)), 0))
      case int64:
        ReturnFromFunction(l, ";", MakeValue(e, float32(p.(int64)), 0))
      case *TCValue:
        ReturnFromFunction(l, ";", MakeValue(e, p.(*TCValue).P, 0))
      default:
        ReturnFromFunction(l, ";", e)
      }
    }
  }
  return nil, nil
}

func init() {
  SetCommand(";", TCStackBackSendFunction)
}
