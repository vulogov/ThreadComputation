package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func (l *TCExecListener) ExecuteOperatorTyped(name string, x_type int, q *deque.Deque) *TCError {
  var fun TCOperatorCommand
  var x interface{}

  if q.Len() < 2 {
    return nil
  }
  x = l.AttrByType(Neural, q)
  if x == nil {
    return l.TC.MakeError(fmt.Sprintf("X type %v not found for ExecuteOperatorTyped", x_type))
  }
  ReturnFromFunction(l, name, x)
  for q.Len() > 0 {
    e := q.PopFront()
    if TCType(e) == x_type {
      //
      // Skip same typed arguments
      //
      continue
    }
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
