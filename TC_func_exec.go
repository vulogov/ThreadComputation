package ThreadComputation

import (
  "fmt"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
)

func (l *TCExecListener) ExecFunction(func_name string, q *deque.Deque) (interface{}, error) {
  if l.TC.HasVariable(func_name)  {
    if q.Len() > 0 {
      e := q.PopFront()
      l.TC.SetVariable(func_name, e)
    }
    log.Debugf("ExecFunction: %v. Returning variable", func_name)
    data, _ := l.TC.GetVariable(func_name)
    return data, nil
  }
  if l.TC.IsUserFunction(func_name) {
    code, err := l.TC.GetUserFunCode(func_name)
    if err != nil {
      return nil, err
    }
    log.Debugf("ExecFunction: %v. User function", func_name)
    l.TC.Eval(code)
    if l.TC.Errors() > 0 {
      return nil, errors.New(l.TC.Error())
    }
    return nil, nil
  }
  if l.TC.HaveCommand(func_name) {
    fun := l.TC.GetCommand(func_name)
    log.Debugf("ExecFunction: %v. Executing command", func_name)
    res, err := fun(l, func_name, q)
    if err != nil {
      return nil, err
    }
    return res, nil
  }
  if l.TC.HaveFunction(func_name) {
    fun := l.TC.GetFunction(func_name)
    res, err := fun(l, func_name, q)
    if err != nil {
      return nil, err
    }
    return res, nil
  }
  if IsSimpleData(func_name) {
    e := GetSimpleData(func_name)
    e = ApplyToFunction(l, e, q)
    return e, nil
  }
  return nil, errors.New(fmt.Sprintf("ExecFunction: %v. Not found", func_name))
}
