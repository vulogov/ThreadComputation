package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)



func TCPrintFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    fun := GetConverterCallback(e)
    if fun == nil {
      log.Debugf("print error: Can not get convert for %T", e)
      continue
    }
    switch e.(type) {
    case string:
      fmt.Printf("%v ", e.(string))
    default:
      fmt.Printf("%v ", fun(e, String))
    }
  }
  return nil, nil
}

func TCPrintlnFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  TCPrintFunction(l, name, q)
  fmt.Printf("\n")
  return nil, nil
}

func TCPrintfFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var args []interface{}

  if q.Len() > 0 {
    f := q.PopFront()
    switch f.(type) {
    case string:
      for q.Len() > 0 {
        v := q.PopFront()
        fun := GetConverterCallback(v)
        if fun == nil {
          continue
        }
        args = append(args, fun(v, String))
      }
      fmt.Printf(f.(string), args...)
    }
  }
  return nil, nil
}

func TCPrintflnFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  TCPrintfFunction(l, name, q)
  fmt.Printf("\n")
  return nil, nil
}

func TCnewlineFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      fun := GetConverterCallback(e)
      if fun == nil {
        log.Debugf("print error: Can not get convert for %T", e)
        continue
      }
      switch e.(type) {
      case string:
        fmt.Printf("%v\n", e.(string))
      default:
        fmt.Printf("%v\n", fun(e, String))
      }
    }
  } else {
    ReturnFromFunction(l, "newline", "\n")
  }
  return nil, nil
}


func init() {
  SetFunction("print", TCPrintFunction)
  SetFunction("println", TCPrintlnFunction)
  SetFunction("printf", TCPrintfFunction)
  SetFunction("printlnf", TCPrintflnFunction)
  SetCommand("newline", TCPrintflnFunction)
}
