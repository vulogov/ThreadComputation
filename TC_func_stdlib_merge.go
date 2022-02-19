package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/srfrog/dict"
  "github.com/deckarep/golang-set"
  "github.com/Jeffail/gabs/v2"
  "github.com/gammazero/deque"
)



func TCMergeFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if ! l.TC.Ready() && q.Len() == 0 {
    return nil, errors.New("merge function did not discover proper context")
  }
  if l.TC.Ready() {
    sr := l.TC.Get()
    switch s := sr.(type) {
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
          s = mergeDmaps(s, e.(*dict.Dict))
        }
      }
      for l.TC.Res.Len() > 0 {
        e := l.TC.Get()
        switch e.(type) {
        case *dict.Dict:
          s = mergeDmaps(s, e.(*dict.Dict))
        }
      }
      return s, nil
    case *gabs.Container:
      for q.Len() > 0 {
        e := q.PopFront()
        switch e.(type) {
        case *gabs.Container:
          s.Merge(e.(*gabs.Container))
        case *dict.Dict:
          s = tcMergeDmapToJSON(s, e.(*dict.Dict))
        case string:
          je, err := gabs.ParseJSON([]byte(e.(string)))
          if err != nil {
            return nil, err
          }
          s.Merge(je)
        }
        return s, nil
      }
      for l.TC.Res.Len() > 0 {
        e := l.TC.Get()
        switch e.(type) {
        case *gabs.Container:
          s.Merge(e.(*gabs.Container))
        case *dict.Dict:
          s = tcMergeDmapToJSON(s, e.(*dict.Dict))
        case string:
          je, err := gabs.ParseJSON([]byte(e.(string)))
          if err != nil {
            return nil, err
          }
          s.Merge(je)
        }
      }
      return s, nil
    }
    return nil, errors.New(fmt.Sprintf("data type that non-compartible with merge is on the stack: %T", sr))
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
