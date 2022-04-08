package ThreadComputation

import (
  "fmt"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func TCImportFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var res interface{}

  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *TCExported:
      b := e.(*TCExported).Export()
      if b == nil {
        return nil, errors.New("import: Can not import TCExport")
      }
      fun := GetConverterCallback(b)
      if b == nil {
        return nil, errors.New("import: Can not load converted for TCBinary")
      }
      switch b.Type {
      case String:
        res = b.Raw()
      default:
        res = fun(b, b.Type)
      }
      log.Debugf("import: converting to %v=%T", TypeToStr(b.Type), res)
      if res == nil {
        return nil, errors.New(fmt.Sprintf("import: Can not convert %v", TypeToStr(b.Type)))
      }
      log.Debugf("import: %T", res)
      ReturnFromFunction(l, "import", res)
    default:
      return nil, errors.New(fmt.Sprintf("import: unacceptable type for import %T", e))
    }
  }
  return nil, nil
}

func init() {
  SetFunction("import", TCImportFunction)
}
