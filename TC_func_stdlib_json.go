package ThreadComputation

import (
  "github.com/srfrog/dict"
  "github.com/Jeffail/gabs/v2"
  "github.com/gammazero/deque"
)




func TCNewJsonFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  res := gabs.New()
  if l.TC.Ready() {
    e := l.TC.Res.Q().Front()
    switch e.(type) {
    case *dict.Dict:
      e := l.TC.Get()
      for item := range e.(*dict.Dict).Items() {
        res.Set(item.Value, item.Key.(string))
      }
    case *gabs.Container:
      e := l.TC.Get()
      res.Merge(e.(*gabs.Container))
    case string:
      e := l.TC.Get()
      je, err := gabs.ParseJSON([]byte(e.(string)))
      if err != nil {
        return nil, err
      }
      res.Merge(je)
    }
  }
  if q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case *dict.Dict:
      for item := range e.(*dict.Dict).Items() {
        res.Set(item.Value, item.Key.(string))
      }
    case *gabs.Container:
      res.Merge(e.(*gabs.Container))
    case string:
      je, err := gabs.ParseJSON([]byte(e.(string)))
      if err != nil {
        return nil, err
      }
      res.Merge(je)
    }
  }
  return res, nil
}

func initStdlibJson() {
  SetFunction("json", TCNewJsonFunction)
}
