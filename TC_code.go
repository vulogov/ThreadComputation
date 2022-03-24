package ThreadComputation

import (
  "fmt"
  "errors"
  log "github.com/sirupsen/logrus"
  "github.com/google/uuid"
  "github.com/gammazero/deque"
)

type TCCode struct {
  Name       string
  Code       deque.Deque
}

func (tc *TCstate) NewCode() *TCCode {
  res := new(TCCode)
  res.Name = uuid.NewString()
  res.New()
  return res
}

func (tc *TCstate) ExecuteCode(c *TCCode) error {
  return c.Execute(tc)
}

func (c *TCCode) New() {
  c.Code.PushBack("")
}

func (c *TCCode) String() string {
  return fmt.Sprintf("code[ %v segment(s)]", c.Code.Len())
}

func (c *TCCode) Add(v interface{}) {
  switch v.(type) {
  case string:
    code := c.Code.PopFront().(string)
    code += fmt.Sprintf("\"%v\" ", v.(string))
    c.Code.PushFront(code)
  case *deque.Deque:
    for v.(*deque.Deque).Len() > 0 {
      e := v.(*deque.Deque).PopFront()
      c.Add(e)
    }
  default:
    cfun := GetConverterCallback(v)
    if cfun != nil {
      code := c.Code.PopFront().(string)
      out := cfun(v, String)
      if out == nil {
        code += fmt.Sprintf("%v ", v)
      } else {
        code += fmt.Sprintf("%v ", out)
      }
      c.Code.PushFront(code)
    } else {
      code := c.Code.PopFront().(string)
      code += fmt.Sprintf("%v ", v)
      c.Code.PushFront(code)
    }
  }
}

func (c *TCCode) AddCall(name string, q *deque.Deque) {
  c.AddRaw(fmt.Sprintf("%v[ ", name))
  c.Add(q)
  c.AddRaw(fmt.Sprintf(" ]"))
}

func (c *TCCode) AddRaw(data string) {
  code := c.Code.PopFront().(string)
  code += fmt.Sprintf("%v ", data)
  c.Code.PushFront(code)
}

func (c *TCCode) Execute(tc *TCstate) error {
  for x := 0; x < c.Code.Len(); x++ {
    line := c.Code.At(x)
    switch line.(type) {
    case string:
      tc.Eval(line.(string))
    }
    if tc.Errors() > 0 {
      return errors.New(tc.Error())
    }
  }
  return nil
}

func BlockCode(l *TCExecListener, name string, code string) interface{} {
  log.Debugf("code %v: %v", name, code)
  res := l.TC.NewCode()
  res.AddRaw(code)
  ReturnFromFunction(l, "code", res)
  return nil
}

func init() {
  RegisterBlockCallback("code", BlockCode)
}
