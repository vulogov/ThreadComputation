package ThreadComputation

import (
  "errors"
  "github.com/deckarep/golang-set"
  "github.com/gammazero/deque"
)



func TCSetFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  s := mapset.NewSet()
  if q.Len() > 0 {
    for q.Len() > 0 {
      data := q.PopFront()
      switch data.(type) {
      case string, int64, float64, bool:
        s.Add(data)
      }
    }
    return s, nil
  } else {
    for l.TC.Ready() {
      data := l.TC.Get()
      switch data.(type) {
      case string, int64, float64, bool:
        s.Add(data)
      }
    }
    return s, nil
  }
  return nil, errors.New("read function did not discover proper context")
}

func TCUnSetFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      s := q.PopFront()
      switch s.(type) {
      case mapset.Set:
        iter := s.(mapset.Set).Iterator()
        for e := range iter.C {
          l.TC.Res.Set(e)
        }
      }
    }
  } else {
    s := l.TC.Get()
    switch s.(type) {
    case mapset.Set:
      iter := s.(mapset.Set).Iterator()
      for e := range iter.C {
        l.TC.Res.Set(e)
      }
    }
  }
  return nil, nil
}

func initStdlibSet() {
  SetFunction("set", TCSetFunction)
  SetFunction("unset", TCUnSetFunction)
}
