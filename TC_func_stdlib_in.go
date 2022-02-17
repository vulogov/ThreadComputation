package ThreadComputation

import (
  "errors"
  "github.com/deckarep/golang-set"
  "github.com/gammazero/deque"
)



func TCInFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var tq deque.Deque
  var found bool
  if ! l.TC.Ready() && q.Len() == 0 {
    return nil, errors.New("in function did not discover proper context")
  }
  for q.Len() > 0 {
    e := q.PopFront()
    found = false
    for x := 0; x < l.TC.Res.Len(); x++ {
      v := l.TC.Res.Q().At(x)
      switch v.(type) {
      case string, int64, float64, bool:
        if e == v {
          found = true
          break
        }
      case mapset.Set:
        if v.(mapset.Set).Contains(e) {
          found = true
          break
        }
      }
    }
    tq.PushFront(found)
  }
  for tq.Len() > 0 {
    l.TC.Res.Set(tq.PopFront())
  }
  return nil, nil
}


func initStdlibIn() {
  SetFunction("in", TCInFunction)
}
