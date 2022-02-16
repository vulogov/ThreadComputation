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
    case string:
      l.TC.Eval(e.(string))
    case int64, float64, bool:
      return e, nil
    default:
      return nil, errors.New("reference execution expects function reference in stack")
    }
  } else {
    return nil, errors.New("Stack is too shallow to execute function reference")
  }
  return nil, nil
}

func ExecuteAllFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var tq deque.Deque
  if q.Len() == 0 {
    for l.TC.Ready() {
      e := l.TC.Get()
      switch e.(type) {
      case *TCFunRef:
        res, err := l.TC.ExecFunction(l, e.(*TCFunRef).Name, e.(*TCFunRef).Attrs)
        if err != nil {
          return nil, err
        }
        tq.PushFront(res)
      case string:
        l.TC.Eval(e.(string))
      case int64, float64, bool:
        tq.PushFront(e)
      default:
        return nil, errors.New(fmt.Sprintf("reference execution expects function reference in stack: %v", e))
      }
    }
  } else {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case *TCFunRef:
        res, err := l.TC.ExecFunction(l, e.(*TCFunRef).Name, e.(*TCFunRef).Attrs)
        if err != nil {
          return nil, err
        }
        tq.PushFront(res)
      case string:
        l.TC.Eval(e.(string))
      case int64, float64, bool:
        tq.PushFront(e)
      default:
        return nil, errors.New("reference execution expects function reference in arguments")
      }
    }
  }
  for tq.Len() > 0 {
    d := tq.PopFront()
    l.TC.Res.Set(d)
  }
  return nil, nil
}

func SetAttrsFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var r interface{}
  var tq deque.Deque
  var wq *deque.Deque
  if l.TC.Ready() {
    r = l.TC.Get()
  } else {
    return nil, errors.New("stack is too shallow for attr[] function")
  }
  switch r.(type) {
  case *TCFunRef:
    if q.Len() > 0 {
      wq = q
      for wq.Len() > 0 {
        e := wq.PopFront()
        tq.PushBack(e)
      }
    } else {
      wq = l.TC.Res.Q()
      for wq.Len() > 0 {
        e := wq.PopFront()
        tq.PushFront(e)
      }
    }
    r.(*TCFunRef).Attrs = &tq
    return r, nil
  default:
    return nil, errors.New(fmt.Sprintf("reference execution expects function reference in stack: %v", r))
  }
  return nil, nil
}

func UseFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      fn := q.PopFront()
      switch fn.(type) {
      case string:
        data, err := readVfsFile(fn.(string))
        if err != nil {
          return nil, err
        }
        l.TC.Eval(data)
      default:
        break
      }
    }
    return nil, nil
  } else {
    if l.TC.Ready() {
      fn := l.TC.Get()
      switch fn.(type) {
      case string:
        data, err := readVfsFile(fn.(string))
        if err != nil {
          return nil, err
        }
        l.TC.Eval(data)
        return nil, nil
      }
    }
  }
  return nil, errors.New("use function did not discover proper context")
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
  SetFunction("!*", ExecuteAllFunction)
  SetFunction("attr", SetAttrsFunction)
  SetFunction("use", UseFunction)
}
