package ThreadComputation

import (
  "fmt"
  "strings"
  "errors"
  "github.com/srfrog/dict"
  "github.com/Jeffail/gabs/v2"
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
  log.Debugf("open call: %v", func_name)
  if l.TC.AddToUserFun(fmt.Sprintf("%v[", func_name)) == true {
    return
  }
  if _, ok := Vars.Load(func_name); ok {
    log.Debugf("global variable: %v", func_name)
    return
  }
  if _, ok := l.TC.Vars.Load(func_name); ok {
    log.Debugf("local variable: %v", func_name)
    return
  }
  if strings.HasPrefix(func_name, "`") {
    log.Debugf("making reference: %v", func_name)
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
          log.Debugf("? call will not result skip: %v", func_name)
          l.TC.ToSkip = false
          l.TC.SkipFunction = func_name
        } else {
          log.Debugf("? call will result skip: %v", func_name)
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
  if l.TC.Ready() {
    switch l.TC.Res.Q().Front().(type) {
    case *dict.Dict:
      if l.TC.Res.Q().Front().(*dict.Dict).Key(func_name) {
        log.Debugf("will return key from dmap in stack: %v", func_name)
        return
      }
    case *gabs.Container:
      if l.TC.Res.Q().Front().(*gabs.Container).Exists() {
        log.Debugf("will return key from json in stack: %v", func_name)
        return
      }
    }
  }
  if _, ok := l.TC.UserFun.Load(func_name); ok {
    log.Debugf("begin defining user function: %v", func_name)
    l.TC.InAttr += 1
    l.TC.Attrs.Add()
    l.TC.FNStack.PushFront(func_name)
    return
  }
  _, ok := Functions.Load(func_name)
  if ok == false {
    _, ok = l.TC.Functions.Load(func_name)
  }
  if ok {
    log.Debugf("will call ether local or global function: %v", func_name)
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
  if l.TC.AddToUserFun("]") == true {
    return
  }
  func_name := c.GetFname().GetText()
  if strings.HasPrefix(func_name, "?") {
    func_name = strings.TrimPrefix(func_name, "?")
  }
  if l.TC.ToSkip {
    if func_name == l.TC.SkipFunction {
      log.Debugf("finishing skipping function call: %v", func_name)
      l.TC.ToSkip = false
      return
    }
  }
  if gvdata, ok := Vars.Load(func_name); ok {
    log.Debugf("returning global value to stack: %v", func_name)
    l.TC.Res.Set(gvdata)
    return
  }
  if vdata, ok := l.TC.Vars.Load(func_name); ok {
    log.Debugf("returning local value to stack: %v", func_name)
    l.TC.Res.Set(vdata)
    return
  }
  if l.TC.Ready() {
    switch l.TC.Res.Q().Front().(type) {
    case *dict.Dict:
      dmdata := l.TC.Res.Q().Front().(*dict.Dict).Get(func_name)
      if dmdata != nil {
        if l.TC.Attrs.GLen() > 1 {
          l.TC.Attrs.Set(dmdata)
        } else {
          l.TC.Res.Set(dmdata)
        }
        return
      }
    case *gabs.Container:
      if l.TC.Res.Q().Front().(*gabs.Container).Exists() {
        jdata := l.TC.Res.Q().Front().(*gabs.Container).Path(func_name).Data()
        ReturnFromFunction(l, func_name, jdata)
        return
      }
    }
  }
  if code, ok := l.TC.UserFun.Load(func_name); ok {
    q   := l.TC.Attrs.Q()
    l.TC.Attrs.Del()
    for q.Len() > 0 {
      e := q.PopFront()
      l.TC.Res.Set(e)
    }
    l.TC.Eval(code.(string))
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
    log.Debugf("calling actual function: %v", func_name)
    res, err := fun(l, q)
    log.Debugf("function %v return res=%T, err=%v", func_name, res, err)
    l.TC.FNStack.PopFront()
    if err != nil {
      l.TC.errmsg = err.Error()
      l.TC.errors += 1
      log.Errorf(l.TC.errmsg)
    } else {
      if res != nil {
        log.Debugf("we got the value: fun=%v, glen=%v", func_name, l.TC.Attrs.GLen())
        if l.TC.Attrs.GLen() >= 1 {
          log.Debugf("returning function value to attribute: %v", func_name)
          l.TC.Attrs.Set(res)
        } else {
          log.Debugf("returning function value to stack: %v", func_name)
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
  if code, ok := l.TC.UserFun.Load(func_name); ok {
    l.TC.FNStack.PushFront(func_name)
    for q.Len() > 0 {
      e := q.PopFront()
      l.TC.Res.Set(e)
    }
    l.TC.Eval(code.(string))
    l.TC.FNStack.PopFront()
    return nil, nil
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

func ReturnFromFunction(l *TCExecListener, func_name string, data interface{}) {
  log.Debugf("return from fun=%v() res=%T", func_name, data)
  if l.TC.Attrs.GLen() >= 1 {
    log.Debugf("returning function value to attribute: %v", func_name)
    l.TC.Attrs.Set(data)
  } else {
    log.Debugf("returning function value to stack: %v", func_name)
    l.TC.Res.Set(data)
  }
}

func SetFunction(name string, fun TCFun) {
  Functions.Delete(name)
  Functions.Store(name, fun)
}

func initStdlib() {
  initStdlibGenerics()
  initStdlibMath()
  initStdlibStack()
  initStdlibCmp()
  initStdlibVfs()
  initStdlibString()
  initStdlibTime()
  initStdlibSet()
  initStdlibMerge()
  initStdlibIn()
  initStdlibDmap()
  initStdlibJson()
}
