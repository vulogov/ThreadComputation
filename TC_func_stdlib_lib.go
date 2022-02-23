package ThreadComputation

import (
  "errors"
  "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
)

func getFloatArray(l *TCExecListener, q *deque.Deque) []float64 {
  res := make([]float64, 0)
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case int64:
        res = append(res, float64(e.(int64)))
      case float64:
        res = append(res, e.(float64))
      }
    }
  } else if l.TC.Ready() {
    for l.TC.Ready() {
      e := l.TC.Get()
      switch e.(type) {
      case int64:
        res = append(res, float64(e.(int64)))
      case float64:
        res = append(res, e.(float64))
      }
    }
  }
  return res
}

func getTwoValues(l *TCExecListener, q *deque.Deque) (interface{}, interface{}, error) {
  if q.Len() >= 2 {
    e1 := q.PopFront()
    e2 := q.PopFront()
    return e1, e2, nil
  } else if q.Len() == 1 {
    e1 := q.PopFront()
    if l.TC.Ready() {
      e2 := l.TC.Get()
      return e1, e2, nil
    }
  } else if l.TC.Res.Len() >= 2 {
    e1 := l.TC.Get()
    e2 := l.TC.Get()
    return e1, e2, nil
  }
  return nil, nil, errors.New("Can not get two values from known context")
}

func collectAllData(l *TCExecListener, q *deque.Deque) interface{} {
  var res interface{}
  var e   interface{}
  var t   int
  if q.Len() > 0 {
    e = q.Front()
  } else {
    e = l.TC.Res.Front()
  }
  switch e.(type) {
  case bool:
    res = make([]bool, 0)
    t = 0
  case int64, float64:
    res = make([]float64, 0)
    t = 1
  case string:
    res = make([]string, 0)
    t = 2
  case *mat.Dense:
    res = make([]*mat.Dense, 0)
    t = 3
  default:
    return nil
  }
  if q.Len() > 0 {
    for q.Len() > 0 {
      v := q.PopFront()
      switch v.(type) {
      case bool:
        if t == 0 {
          res = append(res.([]bool), v.(bool))
        }
      case int64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(int64)))
        }
      case float64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(float64)))
        }
      case string:
        if t == 2 {
          res = append(res.([]string), v.(string))
        }
      case *mat.Dense:
        if t == 3 {
          res = append(res.([]*mat.Dense), v.(*mat.Dense))
        }
      }
    }
  }
  for l.TC.Res.Len() > 0 {
    v := l.TC.Get()
    switch v.(type) {
    case bool:
      if t == 0 {
        res = append(res.([]bool), v.(bool))
      }
    case int64:
      if t == 1 {
        res = append(res.([]float64), float64(v.(int64)))
      }
    case float64:
      if t == 1 {
        res = append(res.([]float64), float64(v.(float64)))
      }
    case string:
      if t == 2 {
        res = append(res.([]string), v.(string))
      }
    case *mat.Dense:
      if t == 3 {
        res = append(res.([]*mat.Dense), v.(*mat.Dense))
      }
    }
  }
  return res
}

func appendDataToArray(v interface{}, res []interface{}) []interface{} {
  res = append(res, v)
  return res
}

func collectData(l *TCExecListener, q *deque.Deque) interface{} {
  var res interface{}
  var e   interface{}
  var t   int
  if q.Len() > 0 {
    e = q.Front()
  } else {
    e = l.TC.Res.Front()
  }
  switch e.(type) {
  case bool:
    res = make([]bool, 0)
    t = 0
  case int64, float64:
    res = make([]float64, 0)
    t = 1
  case string:
    res = make([]string, 0)
    t = 2
  case *mat.Dense:
    res = make([]*mat.Dense, 0)
    t = 3
  default:
    res = nil
  }
  if q.Len() > 0 {
    if l.TC.Ready() {
      v := l.TC.Get()
      switch v.(type) {
      case bool:
        if t == 0 {
          res = append(res.([]bool), v.(bool))
        }
      case int64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(int64)))
        }
      case float64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(float64)))
        }
      case string:
        if t == 2 {
          res = append(res.([]string), v.(string))
        }
      case *mat.Dense:
        if t == 3 {
          res = append(res.([]*mat.Dense), v.(*mat.Dense))
        }
      }
    }
    for q.Len() > 0 {
      v := q.PopFront()
      switch v.(type) {
      case bool:
        if t == 0 {
          res = append(res.([]bool), v.(bool))
        }
      case int64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(int64)))
        }
      case float64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(float64)))
        }
      case string:
        if t == 2 {
          res = append(res.([]string), v.(string))
        }
      case *mat.Dense:
        if t == 3 {
          res = append(res.([]*mat.Dense), v.(*mat.Dense))
        }
      }
    }
  } else {
    for l.TC.Res.Len() > 0 {
      v := l.TC.Get()
      switch v.(type) {
      case bool:
        if t == 0 {
          res = append(res.([]bool), v.(bool))
        }
      case int64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(int64)))
        }
      case float64:
        if t == 1 {
          res = append(res.([]float64), float64(v.(float64)))
        }
      case string:
        if t == 2 {
          res = append(res.([]string), v.(string))
        }
      case *mat.Dense:
        if t == 3 {
          res = append(res.([]*mat.Dense), v.(*mat.Dense))
        }
      }
    }
  }
  return res
}

func applyAggFunToFloatArray(l *TCExecListener, q *deque.Deque, fun func([]float64) float64) (interface{}, error) {
  var acc float64
  data := collectData(l, q)
  if l.TC.Errors() > 0 {
    return nil, errors.New(l.TC.Error())
  }
  acc = fun(data.([]float64))
  return acc, nil
}

func applyAggFunToFloatArrayWithWeight(l *TCExecListener, q *deque.Deque, fun func([]float64, []float64) float64) (interface{}, error) {
  var acc float64
  data := collectData(l, q)
  if l.TC.Errors() > 0 {
    return nil, errors.New(l.TC.Error())
  }
  acc = fun(data.([]float64), nil)
  return acc, nil
}

func applyFunToFloats(l *TCExecListener, q *deque.Deque, fun func(float64) float64) (interface{}, error) {
  data := collectData(l, q).([]float64)
  if l.TC.Errors() > 0 {
    return nil, errors.New(l.TC.Error())
  }
  for x := 0; x < len(data); x++ {
    e := fun(data[x])
    l.TC.Res.Set(e)
  }
  return nil, nil
}
