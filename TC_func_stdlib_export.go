package ThreadComputation

import (
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func TCExportFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *TCExported:
      return e, nil
    case *TCBinary:
      res :=  MakeExportedFromBinary(e.(*TCBinary))
      if res != nil {
        return res, nil
      } else {
        log.Debugf("export: Error exporting TCBinary")
      }
    default:
      b := MakeBinary(e)
      if b != nil {
        res :=  MakeExportedFromBinary(b)
        if res != nil {
          return res, nil
        } else {
          return nil, errors.New("export: Error exporting TCBinary")
        }
      } else {
        return nil, errors.New("export: Error creating TCBinary")
      }
    }
  }
  return nil, nil
}

func init() {
  SetFunction("export", TCExportFunction)
}
