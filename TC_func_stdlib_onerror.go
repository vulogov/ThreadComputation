package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)

func TCNErrorsFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return l.TC.TrueErrors(), nil
}

func TCErrorsHandleFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  log.Debug("Errors are handled by user")
  l.TC.HandleErr = true
  return nil, nil
}

func TCErrorsNotHandleFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  log.Debug("Errors are handled by system")
  l.TC.HandleErr = false
  return nil, nil
}


func TCClearErrorsFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  l.TC.ClearErrors()
  return nil, nil
}

func TCIfErrorFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.TrueErrors() > 0 {
    return true, nil
  }
  return false, nil
}

func TCLastErrorFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return l.TC.Error(), nil
}

func TCThrowErrorFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  msg := ""
  for q.Len() > 0 {
    e := q.PopFront()
    cfun := GetConverterCallback(e)
    if cfun != nil {
      res := cfun(e, String)
      if res == nil {
        res = fmt.Sprintf("%v", e)
      }
      msg += fmt.Sprintf("%v ", res)
    }
  }
  log.Debugf("User-generated error: %v", msg)
  return nil, errors.New(msg)
}

func TCIfLastErrorFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.TrueErrors() > 0 {
    return l.TC.Error(), nil
  }
  return nil, nil
}

func BlockOnError(l *TCExecListener, name string, code string) interface{} {
  if l.TC.TrueErrors() > 0 {
    log.Debugf("onError catches error condition: %v", l.TC.Error())
    l.TC.Eval(code)
  }
  return nil
}

func BlockNotOnError(l *TCExecListener, name string, code string) interface{} {
  if l.TC.TrueErrors() == 0 {
    log.Debugf("notErrors catches good condition")
    l.TC.Eval(code)
  }
  return nil
}

func BlockOnDebug(l *TCExecListener, name string, code string) interface{} {
  if l.TC.IsDebug {
    log.Debugf("Executing debug[] section")
    l.TC.Eval(code)
  }
  return nil
}

func init() {
  RegisterBlockCallback("onerror", BlockOnError)
  RegisterBlockCallback("noterror", BlockNotOnError)
  RegisterBlockCallback("debug", BlockOnDebug)
  SetCommand("errors", TCNErrorsFunction)
  SetCommand("throw", TCThrowErrorFunction)
  SetCommand("errors.Last", TCLastErrorFunction)
  SetCommand("errors.Message", TCIfLastErrorFunction)
  SetCommand("errors.Check", TCIfErrorFunction)
  SetCommand("errors.Clear", TCClearErrorsFunction)
  SetCommand("errors.Handle", TCErrorsHandleFunction)
  SetCommand("errors.NotHandle", TCErrorsNotHandleFunction)
}
