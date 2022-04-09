package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCExportedTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case *TCBinary:
        ReturnFromFunction(l, "Binary", MakeExportedFromBinary(e.(*TCBinary)))
      case *TCValue:
        switch e.(*TCValue).Value.(type) {
        case *TCBinary:
          ReturnFromFunction(l, "Exported", MakeExportedFromBinary(e.(*TCValue).Value.(*TCBinary)))
        }
      }
    }
    return nil, nil
  }
  return MakeType(Exported), nil
}




func init() {
  SetCommand("Exported", TCExportedTypeFunction)
}
