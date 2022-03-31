package ThreadComputation

import (
  "os"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

type TCExitCbFun func(*TCExecListener)

func (l *TCExecListener) ExitRequest() {
  log.Debugf("Exit requesting. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
  l.TC.IsExitReq = true
  l.TC.ExReq <- true
}

func (l *TCExecListener) ExitRequested() bool {
  if l.TC.IsExitReq {
    l.TC.IsExitReq = true
    l.TC.ExReq <- true
    return true
  }
  if len(l.TC.ExReq) > 0 {
    e := <- l.TC.ExReq
    if e == true {
      log.Debugf("Exit requested. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
      l.TC.IsExitReq = true
      l.TC.ExReq <- true
      return true
    }
  }
  return false
}

func (l *TCExecListener) RunExitCallbacks() {
  for cb := range l.TC.ExitCb.Items() {
    log.Debugf("Running exit callback: %v", cb.Key)
    cb.Value.(TCExitCbFun)(l)
  }
}

func (l *TCExecListener) RegisterExitCallback(name string, fun TCExitCbFun) {
  l.TC.RegisterExitCallback(name, fun)
}

func (tc *TCstate) RegisterExitCallback(name string, fun TCExitCbFun) {
  log.Debugf("Reistering exit callback: %v", name)
  tc.ExitCb.Del(name)
  tc.ExitCb.Set(name, fun)
}

func InitExitCallbacks(tc *TCstate) {
  log.Debugf("Initialize exit callbacks")
  tc.RegisterExitCallback("closePool", TCExitCbPoolClose)
}

func (tc *TCstate) ExitRequested() bool {
  if tc.IsExitReq {
    return true
  }
  if len(tc.ExReq) > 0 {
    e := <- tc.ExReq
    if e == true {
      tc.IsExitReq = true
      tc.ExReq <- true
      return true
    }
  }
  return false
}

func TCExitFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  log.Debugf("Exit requested from %v", l.TC.ID)
  l.ExitRequest()
  if l.ExitRequested() {
    log.Debugf("Waiting for threads to terminate. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
    l.TC.Wg.Wait()
    l.RunExitCallbacks()
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case string:
        log.Info(e.(string))
      case bool:
        if e.(bool) {
          log.Debugf("OS-level exit has been requested with return code 0 (zero)")
          os.Exit(0)
        }
      case int64:
        code := int(e.(int64))
        log.Debugf("OS-level exit been requested. Exit code=%v", code)
        os.Exit(code)
      }
    }
    return true, nil
  }
  return false, nil
}

func TCExitRequestedFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.ExitRequested() {
    log.Debugf("Exit been requested. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
    return true, nil
  }
  log.Debugf("Exit not been requested. Capacity=%v, Active=%v", l.TC.Pool.PoolSize(), l.TC.Pool.ActiveWorkers())
  return false, nil
}



func init() {
  SetCommand("exit", TCExitFunction)
  SetCommand("isExit", TCExitRequestedFunction)
}
