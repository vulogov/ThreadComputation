package ThreadComputation

import (
  "fmt"
  "strings"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)

type TCFun func(l *TCExecListener, q *deque.Deque) (interface{}, error)

type TCFunRef struct {
  Name        string
  Attrs      *deque.Deque
}

func (l *TCExecListener) EnterFun(c *parser.FunContext) {
  if l.TC.Errors() > 0 {
    return
  }
  if c.GetFname() == nil {
    l.TC.errmsg = "Can not get function name out of context"
    l.TC.errors += 1
    log.Errorf(l.TC.errmsg)
    return
  }
  func_name := c.GetFname().GetText()
  if _, ok := Vars.Load(func_name); ok {
    return
  }
  if _, ok := l.TC.Vars.Load(func_name); ok {
    return
  }
  if strings.HasPrefix(func_name, "`") {
    l.TC.InAttr += 1
    l.TC.InRef  += 1
    l.TC.Attrs.Add()
    return
  }
  if strings.HasPrefix(func_name, "?") {
    if l.TC.Ready() {
      e := l.TC.Get()
      switch e.(type) {
      case bool:
        func_name = strings.TrimPrefix(func_name, "?")
        if e.(bool) {
          l.TC.ToSkip = false
          l.TC.SkipFunction = func_name
        } else {
          l.TC.ToSkip = true
          l.TC.SkipFunction = func_name
        }
      default:
        l.TC.errmsg = "Boolean value is expected for conditional"
        l.TC.errors += 1
        log.Errorf(l.TC.errmsg)
        return
      }
    } else {
      l.TC.errmsg = "Stack is too shallow for conditional"
      l.TC.errors += 1
      log.Errorf(l.TC.errmsg)
      return
    }
  }
  _, ok := Functions.Load(func_name)
  if ok == false {
    _, ok = l.TC.Functions.Load(func_name)
  }
  if ok {
		l.TC.InAttr += 1
    l.TC.Attrs.Add()
    l.TC.FNStack.PushFront(func_name)
  } else {
    l.TC.errmsg = fmt.Sprintf("Function: %v not found", func_name)
    l.TC.errors += 1
    log.Errorf(l.TC.errmsg)
  }
}

func (l *TCExecListener) ExitFun(c *parser.FunContext) {
  var lfun interface{}
  if l.TC.InAttr > 0 {
    l.TC.InAttr -= 1
  }
  if l.TC.Errors() > 0 {
    return
  }
  func_name := c.GetFname().GetText()
  if strings.HasPrefix(func_name, "?") {
    func_name = strings.TrimPrefix(func_name, "?")
  }
  if l.TC.ToSkip {
    if func_name == l.TC.SkipFunction {
      l.TC.ToSkip = false
      return
    }
  }
  if gvdata, ok := Vars.Load(func_name); ok {
    l.TC.Res.Set(gvdata)
    return
  }
  if vdata, ok := l.TC.Vars.Load(func_name); ok {
    l.TC.Res.Set(vdata)
    return
  }
  if strings.HasPrefix(func_name, "`") {
    func_name := strings.TrimPrefix(func_name, "`")
    if l.TC.InRef > 0 {
      l.TC.InRef -= 1
    }
    r := &TCFunRef{
      Name: func_name,
      Attrs: l.TC.Attrs.Q(),
    }
    l.TC.Attrs.Del()
    if l.TC.Attrs.GLen() > 1 {
      l.TC.Attrs.Set(r)
    } else {
      l.TC.Res.Set(r)
    }
    return
  }
  lfun, ok := Functions.Load(func_name)
  if ok == false {
    lfun, ok = l.TC.Functions.Load(func_name)
  }
  if ok {
    fun := lfun.(TCFun)
    q   := l.TC.Attrs.Q()
    l.TC.Attrs.Del()
    res, err := fun(l, q)
    l.TC.FNStack.PopFront()
    if err != nil {
      l.TC.errmsg = err.Error()
      l.TC.errors += 1
      log.Errorf(l.TC.errmsg)
    } else {
      if res != nil {
        if l.TC.Attrs.GLen() > 1 {
          l.TC.Attrs.Set(res)
        } else {
          l.TC.Res.Set(res)
        }
      }
    }
  }
}

func (tc *TCstate) ExecFunction(l *TCExecListener, func_name string, q *deque.Deque) (interface {}, error) {
  if gvdata, ok := Vars.Load(func_name); ok {
    return gvdata, nil
  }
  if vdata, ok := tc.Vars.Load(func_name); ok {
    return vdata, nil
  }
  lfun, ok := Functions.Load(func_name)
  if ok == false {
    lfun, ok = tc.Functions.Load(func_name)
  }
  if ok {
    l.TC.FNStack.PushFront(func_name)
    fun := lfun.(TCFun)
    res, err := fun(l, q)
    l.TC.FNStack.PopFront()
    return res, err
  } else {
    l.TC.errmsg = fmt.Sprintf("Function: %v not found", func_name)
    l.TC.errors += 1
    return nil, errors.New(l.TC.errmsg)
  }
  return nil, nil
}

func SetFunction(name string, fun TCFun) {
  Functions.Delete(name)
  Functions.Store(name, fun)
}

func initStdlib() {
  initStdlibGenerics()
  initStdlibMath()
  initStdlibStack()
}
