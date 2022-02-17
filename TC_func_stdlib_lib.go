package ThreadComputation

import (
  "errors"
  "github.com/gammazero/deque"
)

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
    }
  }
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
  default:
    res = nil
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
