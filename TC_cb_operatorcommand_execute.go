package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func (l *TCExecListener) ExecuteOperatorCommand(name string, q *deque.Deque) *TCError {
  var fun TCOperatorCommand
  var x interface{}

  if q.Len() < 2 {
    return l.TC.MakeError("Stack is too shallow for operator command")
  }
  rx := q.PopFront()
  switch rx.(type) {
  case *TCValue:
    x = rx.(*TCValue).Value
  default:
    x = rx
  }
  ReturnFromFunction(l, name, rx)
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *TCValue:
      fun = GetOperatorCmdCallback(name, x, e.(*TCValue).Value)
    default:
      fun = GetOperatorCmdCallback(name, x, e)
    }
    if fun == nil {
      return l.TC.MakeError(fmt.Sprintf("callback for %v(%T %T) opcmd not found", name, x, e))
    }
    fun(x, e)
    log.Debugf("operator cmd %v(%T %T)", name, x, e)
  }
  return nil
}
