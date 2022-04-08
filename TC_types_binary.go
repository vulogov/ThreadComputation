package ThreadComputation

import (
  "github.com/gammazero/deque"
)

func TCBinaryTypeFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case string:
        ReturnFromFunction(l, "Binary", MakeBinary(e.(string)))
      case *TCBinary:
        ReturnFromFunction(l, "Binary", MakeBinary(e.(*TCBinary).D))
      case *TCValue:
        switch e.(*TCValue).Value.(type) {
        case string:
          ReturnFromFunction(l, "Binary", MakeBinary(e.(*TCValue).Value.(string)))
        case *TCBinary:
          ReturnFromFunction(l, "Binary", MakeBinary(e.(*TCValue).Value.(*TCBinary).D))
        default:
          res := MakeBinary(e.(*TCValue).Value)
          if res != nil {
            ReturnFromFunction(l, "Binary", res)
          }
        }
      default:
        res := MakeBinary(e)
        if res != nil {
          ReturnFromFunction(l, "Binary", res)
        }
      }
    }
    return nil, nil
  }
  return MakeType(Binary), nil
}




func init() {
  SetCommand("Binary", TCBinaryTypeFunction)
}
