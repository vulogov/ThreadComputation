package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func (l *TCExecListener) ExecuteSingleArgumentFunction(name string, q *deque.Deque) *TCError {
  var fun TCGenericFunction

  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *TCValue:
      fun = GetFunctionCallback(name, e.(*TCValue).Value)
    default:
      fun = GetFunctionCallback(name, e)
    }
    if fun == nil {
      return l.TC.MakeError(fmt.Sprintf("callback for %v(%T) not found", name, e))
    }
    res := fun(e)
    if res == nil {
      //
      // Yes, we do not care if execute function returned nothing
      //
      // return l.TC.MakeError(fmt.Sprintf("callback for %v(%T) returned nil", name, e))
      return nil
    }
    log.Debugf("function %v(%T) = %v", name, e, res)
    switch e.(type) {
    case *TCValue:
      ReturnFromFunction(l, name, MakeValue(res, e.(*TCValue).P, e.(*TCValue).TTL))
    default:
      ReturnFromFunction(l, name, res)
    }
  }
  return nil
}
