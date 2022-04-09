package ThreadComputation

import (
  "fmt"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

type TCRouterFun func(*TCExecListener, string, *deque.Deque) (bool, interface{}, error)

var FunctionRouterQ deque.Deque

func TCFunctionRouter(l *TCExecListener, func_name string, q *deque.Deque) (interface{}, error) {
  log.Debugf("%v functions regusterted in function router", FunctionRouterQ.Len())
  for x := 0; x < FunctionRouterQ.Len(); x++ {
    fun := FunctionRouterQ.At(x).(TCRouterFun)
    fname := GetFunctionName(fun)
    log.Debugf("Calling function-router function %v()", fname)
    served, res, err := fun(l,func_name, q)
    if err != nil {
      return nil, err
    }
    if res != nil {
      ReturnFromFunction(l, fname, res)
    }
    log.Debugf("%v() served status: %v", fname, served)
    if served {
      return nil, nil
    }
  }
  return nil, l.TC.MakeError(fmt.Sprintf("%v is not recognized as function", func_name))
}

func TCAddRouterFunction(fun TCRouterFun) {
  FunctionRouterQ.PushBack(fun)
}
