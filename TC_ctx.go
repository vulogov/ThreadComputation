package ThreadComputation
import (
  "fmt"
  "github.com/google/uuid"
  "github.com/lrita/cmap"
  log "github.com/sirupsen/logrus"
)

type TCCtx struct {
  Id       string
  C        cmap.Cmap
}

func MakeContext() *TCCtx {
  res := new(TCCtx)
  res.Id = uuid.NewString()
  return res
}

func (c *TCCtx) String() string {
  return fmt.Sprintf("local[ %v element(s)]", c.C.Count())
}

func IsContext(x interface{}) bool {
  switch x.(type) {
  case *TCCtx:
    return true
  }
  return false
}

func (tc *TCstate) SetContext(name string, data interface{}) {
  if tc.CtxStack.Len() == 0 {
    tc.Ctx.Delete(name)
    tc.Ctx.Store(name, data)
    log.Debugf("global context %v->%v", data, name)
  } else {
    e := tc.CtxStack.Front().(*TCCtx)
    e.C.Delete(name)
    e.C.Store(name, data)
    log.Debugf("local context %v %v->%v", e.Id, data, name)
  }
}

func (tc *TCstate) GetContext(name string) interface {} {
  if tc.CtxStack.Len() == 0 {
    if data, ok := tc.Ctx.Load(name); ok {
      log.Debugf("global context %v->%v", data, name)
      return data
    }
  } else {
    e := tc.CtxStack.Front().(*TCCtx)
    if data, ok := e.C.Load(name); ok {
      log.Debugf("local context %v %v->%v", e.Id, data, name)
      return data
    }
  }
  return nil
}

func (tc *TCstate) HaveContext(name string) bool {
  if tc.GetContext(name) == nil {
    return false
  }
  return true
}

func (tc *TCstate) AddContext(c interface{})  {
  if c == nil {
    ctx := MakeContext()
    log.Debugf("local context %v added", ctx.Id)
    tc.CtxStack.PushFront(ctx)
  } else {
    switch c.(type) {
    case *TCCtx:
      tc.CtxStack.PushFront(c.(*TCCtx))
    }
  }
}

func (tc *TCstate) DelContext()  {
  if tc.CtxStack.Len() > 0 {
    ctx := tc.CtxStack.PopFront()
    log.Debugf("local context %v deleted", ctx.(*TCCtx).Id)
    tc.CtxStack.PushFront(ctx)
  }
}
