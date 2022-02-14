package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/gammazero/deque"
  "github.com/lrita/cmap"
)

func toString(data interface{}) (string, error) {
  switch data.(type) {
  case int64, float64, string, bool:
    return fmt.Sprintf("%v", data), nil
  case *TCFunRef:
    out := fmt.Sprintf("`%v", data.(*TCFunRef).Name)
    if data.(*TCFunRef).Attrs.Len() > 0 {
      out += "[ "
      for x := 0; x < data.(*TCFunRef).Attrs.Len(); x++ {
        p, err := toString(data.(*TCFunRef).Attrs.At(x))
        if err != nil {
          return "", err
        }
        out += fmt.Sprintf("%v ", p)
      }
      out += "]"
    }
    return out, nil
  case nil:
    return "#NIL", nil
  }
  return "#ERROR", errors.New("Unknown data type in conversion to string")
}

func PrintFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      out, err := toString(e)
      if err != nil {
        return nil, err
      }
      fmt.Println(out)
    }
  } else {
    if l.TC.Ready() {
      data_out := l.TC.Get()
      out, err := toString(data_out)
      if err != nil {
        return nil, err
      }
      fmt.Println(out)
      return data_out, nil
    }
  }
  return nil, nil
}

func PrintAllFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      out, err := toString(e)
      if err != nil {
        return nil, err
      }
      fmt.Println(out)
    }

  } else {
    if l.TC.Ready() {
      for l.TC.Ready() {
        data_out := l.TC.Get()
        out, err := toString(data_out)
        if err != nil {
          return nil, err
        }
        fmt.Println(out)
      }
      return nil, nil
    }
  }
  return nil, nil
}

func ToStackFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, nil
  }
  for q.Len() > 0 {
    e:= q.PopFront()
    l.TC.Res.Set(e)
  }
  return nil, nil
}

func LenFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.Len(), nil
  } else {
    if l.TC.Ready() {
      return l.TC.Res.Len(), nil
    }
  }
  return 0, nil
}

func ClrFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    for l.TC.Ready() {
      l.TC.Get()
    }
  }
  return nil, nil
}

func DropFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    l.TC.Res.PopFront()
  }
  return nil, nil
}

func DupFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    e := l.TC.Get()
    l.TC.Res.Set(e)
    l.TC.Res.Set(e)
  }
  return nil, nil
}

func TheUltimateAnswerFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.PopFront(), nil
  }
  return 42, nil
}

func setLocalOrGlobalFunction(l *TCExecListener, q *deque.Deque, d *cmap.Cmap) error {
  if l.TC.Res.Len() != q.Len() {
    return errors.New("Number of requested variable set not matched with depth of the stack")
  }
  for q.Len() > 0 {
    n := q.PopFront()
    switch n.(type) {
    case string:
      data, err := l.TC.Res.Take()
      if err != nil {
        return err
      }
      d.Delete(n)
      d.Store(n, data)
    default:
      return errors.New("Variable name must be defined in string")
    }
  }
  return nil
}

func SetGlobalFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return nil, setLocalOrGlobalFunction(l, q, &Vars)
}

func SetLocalFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  return nil, setLocalOrGlobalFunction(l, q, &l.TC.Vars)
}

func ExecuteFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    e := l.TC.Get()
    switch e.(type) {
    case *TCFunRef:
      if q.Len() == 0 {
        return l.TC.ExecFunction(l, e.(*TCFunRef).Name, e.(*TCFunRef).Attrs)
      } else {
        return l.TC.ExecFunction(l, e.(*TCFunRef).Name, q)
      }
    default:
      return nil, errors.New("reference execution expects function reference in stack")
    }
  } else {
    return nil, errors.New("Stack is too shallow to execute function reference")
  }
}

func initStdlibGenerics() {
  SetFunction("print", PrintFunction)
  SetFunction("printAll", PrintAllFunction)
  SetFunction("stack", ToStackFunction)
  SetFunction("len", LenFunction)
  SetFunction("drop", DropFunction)
  SetFunction(",", DropFunction)
  SetFunction("dup", DupFunction)
  SetFunction("^", DupFunction)
  SetFunction("clr", ClrFunction)
  SetFunction("TheAnswer", TheUltimateAnswerFunction)
  SetFunction("local", SetLocalFunction)
  SetFunction("global", SetGlobalFunction)
  SetFunction("!", ExecuteFunction)
}
