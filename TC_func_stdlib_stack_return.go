package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func TCStackReturnFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 1 {
    log.Debug("Return last stack value to the stack on left")
    e := q.PopFront()
    l.TC.StacksLeft(1)
    l.TC.Set(e)
    l.TC.StacksRight(1)
    return e, nil
  } else {
    log.Debug("Return last stack value to the list of stacks")
    e := q.PopFront()
    sn := l.TC.ResNames.Front().(string)
    for q.Len() > 0 {
      n := q.PopFront()
      switch n.(type) {
      case string:
        err := l.TC.PositionStack(n.(string))
        if err != nil {
          return nil, l.TC.MakeError(err)
        }
        l.TC.Set(e)
      case int64:
        l.TC.StacksLeft(int(n.(int64)))
        l.TC.Set(e)
        l.TC.PositionStack(sn)
      }
    }
    l.TC.PositionStack(sn)
    return e, nil
  }
  return nil, nil
}

func init() {
  SetFunction("_", TCStackReturnFunction)
}
