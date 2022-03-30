package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

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
