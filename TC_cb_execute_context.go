package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)




func TCExecuteContext(l *TCExecListener, v interface{}, q *deque.Deque) interface{} {
  if TCType(v) != Context {
    log.Debugf("context-execute: %T is not a Context", v)
    return nil
  }
  c := v.(*TCCtx)
  c = TCFillCtx(c, q)
  l.TC.AddContext(c)
  return nil
}


func init() {
  RegisterExecuteCallback(Context, TCExecuteContext)
}
