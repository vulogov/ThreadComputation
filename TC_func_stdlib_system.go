package ThreadComputation

import (
  "strings"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)


func TCLoglevelFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    lvl, err := GetVariable("tc.Debuglevel")
    if err != nil {
      return nil, err
    }
    return lvl, nil
  }
  for q.Len() > 0 {
    e := q.PopFront()
    switch lvl := e.(type) {
    case string:
      lvl = strings.ToLower(lvl)
      switch lvl {
      case "trace":
        l.TC.IsDebug = true
        log.SetLevel(log.TraceLevel)
      case "debug":
        l.TC.IsDebug = true
        log.SetLevel(log.DebugLevel)
      case "info":
        l.TC.IsDebug = false
        log.SetLevel(log.InfoLevel)
      case "warning":
        l.TC.IsDebug = false
        log.SetLevel(log.WarnLevel)
      case "error":
        l.TC.IsDebug = false
        log.SetLevel(log.ErrorLevel)
      case "fatal":
        l.TC.IsDebug = false
        log.SetLevel(log.FatalLevel)
      default:
        l.TC.IsDebug = false
        log.SetLevel(log.InfoLevel)
      }
      log.Debugf("Loglevel set to: %v", lvl)
    }
  }
  return nil, nil
}

func TCTestEnableFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  l.TC.IsTest = true
  return nil, nil
}

func TCTestDisableFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  l.TC.IsTest = false
  return nil, nil
}

func TCObserveEnableFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  l.TC.IsObserve = true
  return nil, nil
}

func TCObserveDisableFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  l.TC.IsObserve = false
  return nil, nil
}

func BlockTesting(l *TCExecListener, name string, code string) interface{} {
  if l.TC.IsTest {
    log.Debugf("Executing testing[] section")
    l.TC.Eval(code)
  }
  return nil
}

func BlockMonitoring(l *TCExecListener, name string, code string) interface{} {
  if l.TC.IsTest {
    log.Debugf("Executing monitoring[] section")
    l.TC.Eval(code)
  }
  return nil
}

func TCNSNameFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.ResNames.Len() == 0 {
    return nil, l.TC.MakeError("Namespace context is too shallow for getting name")
  }
  return l.TC.ResNames.Front().(string), nil
}

func TCNSPreviousNameFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.ResNames.Len() == 0 {
    return nil, l.TC.MakeError("Namespace context is too shallow for getting name")
  }
  l.TC.StacksLeft(1)
  sname := l.TC.ResNames.Front().(string)
  l.TC.StacksRight(1)
  return sname, nil
}

func init() {
  RegisterBlockCallback("testing", BlockTesting)
  RegisterBlockCallback("monitoring", BlockMonitoring)
  SetCommand("system.Loglevel", TCLoglevelFunction)
  SetCommand("system.TestEnable", TCTestEnableFunction)
  SetCommand("system.TestDisable", TCTestDisableFunction)
  SetCommand("system.ObservabilityEnable", TCObserveEnableFunction)
  SetCommand("system.ObservabilityDisable", TCObserveDisableFunction)
  SetCommand("name", TCNSNameFunction)
  SetCommand("previous", TCNSPreviousNameFunction)
}
