package ThreadComputation
import (
  "github.com/gammazero/deque"
)



func TCSetGlobalCtxFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    switch p := e.(type) {
    case *TCPair:
      switch key := p.X.(type) {
      case string:
        l.TC.Ctx.Delete(key)
        l.TC.Ctx.Store(key, p.Y)
      }
    }
  }
  return nil, nil
}

func init() {
  SetCommand("global", TCSetGlobalCtxFunction)
}
