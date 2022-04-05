package ThreadComputation

import (
  "fmt"
  "path"
  "github.com/srfrog/dict"
  "github.com/gammazero/deque"
)

var TCIOReadRouters = dict.New()
var TCIOWriteRouters = dict.New()

type TCIOReadFun func(*TCIO, string) *TCBinary
type TCIOWriteFun func(*TCIO, string, interface{}) interface{}

func TCIORegisterReadRouter(patt string, fun TCIOReadFun) {
  TCIOReadRouters.Del(patt)
  TCIOReadRouters.Set(patt, fun)
}

func TCIOGetReadRouter(name string) TCIOReadFun {
  for item := range(TCIOReadRouters).Items() {
    res, _ := path.Match(item.Key.(string), name)
    if res {
      return item.Value.(TCIOReadFun)
    }
  }
  return nil
}

func TCIORegisterWriteRouter(patt string, fun TCIOWriteFun) {
  TCIOWriteRouters.Del(patt)
  TCIOWriteRouters.Set(patt, fun)
}

func TCIOGetWriteRouter(name string) TCIOWriteFun {
  for item := range(TCIOWriteRouters).Items() {
    res, _ := path.Match(item.Key.(string), name)
    if res {
      return item.Value.(TCIOWriteFun)
    }
  }
  return nil
}

func tcReadFromIO(io *TCIO) {
  for item := range io.D.Items() {
    path := item.Value
    fun := TCIOGetReadRouter(path.(string))
    if fun != nil {
      res := fun(io, path.(string))
      if res != nil {
        ReturnFromFunction(io.L, "Read", MakeBinary(res))
      }
    } else {
      res, err := ReadFile(path.(string))
      if err != nil {
        io.L.SetError(fmt.Sprintf("IO error: %v", err))
      } else {
        ReturnFromFunction(io.L, "Read", MakeBinary(res))
      }
    }
  }
}

func TCReadFromIO(v interface{}) interface{} {
  switch io := v.(type) {
  case *TCIO:
    tcReadFromIO(io)
  }
  return nil
}

func TCIOReadFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteSingleArgumentFunction("ioget", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func tcWriteToIO(io *TCIO, d interface{}) {
  for item := range io.D.Items() {
    path := item.Value
    fun := TCIOGetWriteRouter(path.(string))
    if fun != nil {
      res := fun(io, path.(string), d)
      if res != nil {
        ReturnFromFunction(io.L, "Write", res)
      }
    } else {
      switch v := d.(type) {
      case string:
        err := writeVfsFile(path.(string), []byte(v))
        if err != nil {
          io.L.SetError(fmt.Sprintf("IO error: %v", err))
          return
        }
      case *TCBinary:
        err := writeVfsFile(path.(string), v.D)
        if err != nil {
          io.L.SetError(fmt.Sprintf("IO error: %v", err))
          return
        }
      }
    }
  }
}

func TCWriteToIO(x interface{}, y interface{}) {
  switch d := x.(type) {
  case *TCIO:
    tcWriteToIO(d, y)
  }
}

func TCIOWriteFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  err := l.ExecuteOperatorCommand("ioset", q)
  if err != nil {
    return nil, err
  }
  return nil, nil
}

func init() {
  SetFunction("Read", TCIOReadFunction)
  SetFunction("Write", TCIOWriteFunction)
  RegisterFunctionCallback("ioget", IO, TCReadFromIO)
  RegisterOperatorCmdCallback("ioset", IO, String, TCWriteToIO)
  RegisterOperatorCmdCallback("ioset", IO, Binary, TCWriteToIO)
}
