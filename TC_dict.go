package ThreadComputation
import (
  "fmt"
  "github.com/srfrog/dict"
  "github.com/gammazero/deque"
)

type TCDict struct {
  D        *dict.Dict
}

func tcFillDict(d *dict.Dict, q *deque.Deque) *dict.Dict {
  if (q.Len() % 2) != 0 {
    return nil
  }
  for q.Len() > 0 {
    k := q.PopFront()
    v := q.PopFront()
    switch k.(type) {
    case string:
      d.Set(k.(string), v)
    default:
      return nil
    }
  }
  return d
}

func (tc *TCstate) Dict() *TCDict {
  res := new(TCDict)
  res.D = dict.New()
  return res
}

func (r *TCDict) String() string {
  var value string
  out := "dict[ "
  for item := range r.D.Items() {
    switch item.Value.(type) {
    case string:
      value = item.Value.(string)
    default:
      fun := GetConverterCallback(item.Value)
      if fun == nil {
        continue
      }
      r := fun(item.Value, String)
      if r == nil {
        continue
      }
      value = r.(string)
    }

    out += fmt.Sprintf("%v : '%v'", item.Key, value)
  }
  out += "]"
  return out
}

func TCDictFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := l.TC.Dict()
  res.D = tcFillDict(res.D, q)
  return res, nil
}

func TCDictValueFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  res := l.TC.Dict()
  res.D = tcFillDict(res.D, q)
  return MakeValue(res, 100.0, 0), nil
}

func init() {
  SetCommand("dict", TCDictFunction)
  SetCommand("Dict", TCDictValueFunction)
}
