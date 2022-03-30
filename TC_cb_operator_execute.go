package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func (l *TCExecListener) ExecuteOperator(name string, q *deque.Deque) *TCError {
  var fun TCGenericOperator
  var x interface{}
  var p, cp float32

  if q.Len() < 2 {
    return l.TC.MakeError("Stack is too shallow for operator")
  }
  rx := q.PopFront()
  cp = 100.0
  is_value := false
  switch rx.(type) {
  case *TCValue:
    is_value = true
    x = rx.(*TCValue).Value
    p = rx.(*TCValue).P
  default:
    x = rx
  }
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *TCValue:
      cp = (p + e.(*TCValue).P)/2
      fun = GetOperatorCallback(name, x, e.(*TCValue).Value)
    default:
      fun = GetOperatorCallback(name, x, e)
    }
    if fun == nil {
      return l.TC.MakeError(fmt.Sprintf("callback for %v(%T) not found", name, e))
    }
    res := fun(x, e)
    if res == nil {
      log.Debugf("callback for %v(%T %T) returned nil", name, x, e)
      continue
    }
    log.Debugf("operator %v(%T %T) = %v", name, x, e, res)
    switch e.(type) {
    case *TCValue:
      ReturnFromFunction(l, name, MakeValue(res, float32(cp), e.(*TCValue).TTL))
    default:
      if is_value {
        ReturnFromFunction(l, name, MakeValue(res, rx.(*TCValue).P, rx.(*TCValue).TTL))
      } else {
        ReturnFromFunction(l, name, res)
      }
    }
  }
  return nil
}
