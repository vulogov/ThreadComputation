package ThreadComputation

import (
  "fmt"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)

func TCCompareFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, l.TC.MakeError("Nothing to compare with")
  }
  if q.Len() == 1 {
    return true, nil
  }
  x := q.PopFront()
  for q.Len() > 0 {
    y := q.PopFront()
    fun := GetCompareCallback(x, y)
    if fun == nil {
      return nil, l.TC.MakeError(fmt.Sprintf("Error getting comparator for %T %v %T", x, name, y))
    }
    log.Debugf("comparing: %T %v %T", x, name, y)
    res := fun(name, x, y)
    if res == nil {
      return nil, l.TC.MakeError(" %T %v %T return nil", x, name, y)
    }
    ReturnFromFunction(l, name, res)
  }
  return nil, nil
}

func init() {
  SetFunction("=", TCCompareFunction)
  SetFunction("!=", TCCompareFunction)
  SetFunction("<", TCCompareFunction)
  SetFunction(">", TCCompareFunction)
  SetFunction("<=", TCCompareFunction)
  SetFunction("=>", TCCompareFunction)
}
