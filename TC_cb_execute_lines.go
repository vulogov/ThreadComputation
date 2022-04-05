package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)




func TCExecuteLines(l *TCExecListener, v interface{}, q *deque.Deque) interface{} {
  if TCType(v) != Lines {
    log.Debugf("lines-execute: %T is not a Lines", v)
    return nil
  }
  gfun := GetGenCallback(v)
  if gfun == nil {
    log.Debugf("lines-execute: Can not get generator for %T", v)
    return nil
  }
  iter := gfun(v)
  if iter == nil {
    log.Debugf("lines-execute: Can not get iterator for %T", v)
    return nil
  }
  for q.Len() > 0 {
    elem := q.PopFront()
    switch p := elem.(type) {
    case *TCPair:
      switch p.X.(type) {
      case string:
        key := p.X.(string)
        val := p.Y
        iter.Set(key, val)
      }
    }
  }
  return iter
}


func init() {
  RegisterExecuteCallback(Lines, TCExecuteLines)
}
