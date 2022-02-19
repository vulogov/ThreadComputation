package ThreadComputation

import (
  "errors"
  "github.com/srfrog/dict"
  "github.com/deckarep/golang-set"
  "github.com/gammazero/deque"
)



func TCMergeFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if ! l.TC.Ready() && q.Len() == 0 {
    return nil, errors.New("merge function did not discover proper context")
  }
  if l.TC.Ready() {
    s := l.TC.Get()
    switch s.(type) {
    case mapset.Set:
      for q.Len() > 0 {
        e := q.PopFront()
        switch e.(type) {
        case string, int64, float64, bool:
          s.(mapset.Set).Add(e)
        case mapset.Set:
          s = s.(mapset.Set).Union(e.(mapset.Set))
        }
      }
      for l.TC.Res.Len() > 0 {
        e := l.TC.Get()
        switch e.(type) {
        case string, int64, float64, bool:
          s.(mapset.Set).Add(e)
        case mapset.Set:
          s = s.(mapset.Set).Union(e.(mapset.Set))
        }
      }
      return s, nil
    case *dict.Dict:
      for q.Len() > 0 {
        e := q.PopFront()
        switch e.(type) {
        case *dict.Dict:
          s = mergeDmaps(s.(*dict.Dict), e.(*dict.Dict))
        }
      }
      for l.TC.Res.Len() > 0 {
        e := l.TC.Get()
        switch e.(type) {
        case *dict.Dict:
          s = mergeDmaps(s.(*dict.Dict), e.(*dict.Dict))
        }
      }
      return s, nil
    }
    return nil, errors.New("data type that non-compartible with merge is on the stack")
  }
  return nil, nil
}

func TCUnMergeFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if ! l.TC.Ready() && q.Len() == 0 {
    return nil, errors.New("merge function did not discover proper context")
  }
  if l.TC.Ready() {
    s := l.TC.Get()
    switch s.(type) {
    case mapset.Set:
      for q.Len() > 0 {
        e := q.PopFront()
        switch e.(type) {
        case string, int64, float64, bool:
          s.(mapset.Set).Remove(e)
        case mapset.Set:
          s = s.(mapset.Set).Intersect(e.(mapset.Set))
        }
      }
      for l.TC.Res.Len() > 0 {
        e := l.TC.Get()
        switch e.(type) {
        case string, int64, float64, bool:
          s.(mapset.Set).Remove(e)
        case mapset.Set:
          s = s.(mapset.Set).Intersect(e.(mapset.Set))
        }
      }
      return s, nil
    }
    return nil, errors.New("data type that non-compartible with merge is on the stack")
  }
  return nil, nil
}

func initStdlibMerge() {
  SetFunction("merge", TCMergeFunction)
  SetFunction("+++", TCMergeFunction)
  SetFunction("---", TCUnMergeFunction)
  SetFunction("unmerge", TCUnMergeFunction)
}
