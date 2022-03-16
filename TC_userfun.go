package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/google/uuid"
)

func (tc *TCstate) StartUserFun(name string) string {
  tc.UFNStack.PushFront(name)
  tc.UFStack.PushFront(fmt.Sprintf("/* function: %v */ ", name))
  return name
}

func (tc *TCstate) StartLambdaFun() string {
  return tc.StartUserFun(uuid.NewString())
}

func (tc *TCstate) EndUserFun() (string, error) {
  if tc.HaveUserFun() {
    fname := tc.UFNStack.PopFront().(string)
    code  := tc.UFStack.PopFront().(string)
    tc.UFunctions.Delete(fname)
    tc.UFunctions.Store(fname, code)
    return fname, nil
  }
  return "", errors.New("There is no user defined function in progress")
}

func (tc *TCstate) HaveUserFun() bool {
  if tc.UFNStack.Len() > 0 {
    return true
  }
  return false
}

func (tc *TCstate) GetUserFunCode(fname string) (string, error) {
  if code, ok := tc.UFunctions.Load(fname); ok {
    return code.(string), nil
  }
  return "", errors.New(fmt.Sprintf("User defined function %v not found", fname))
}

func (tc *TCstate) MakeUserFun(mod interface{}, name string) bool {
  if tc.HaveUserFun() {
    out := ""
    if mod != nil {
      out += fmt.Sprintf("%v", mod)
    }
    out += fmt.Sprintf("%v[ ", name)
    c := tc.UFStack.PopFront().(string)
    c += out
    tc.UFStack.PushFront(c)
    return true
  }
  return false
}

func (tc *TCstate) FinishUserFun() bool {
  if tc.HaveUserFun() {
    c := tc.UFStack.PopFront().(string)
    c += " ]"
    tc.UFStack.PushFront(c)
    return true
  }
  return false
}

func (tc *TCstate) AddToUserFun(data string) bool {
  if tc.HaveUserFun() {
    c := tc.UFStack.PopFront().(string)
    c += fmt.Sprintf("%v ", data)
    tc.UFStack.PushFront(c)
    return true
  }
  return false
}
