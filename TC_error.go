package ThreadComputation

import (
  "fmt"
  "errors"
  "time"
  "github.com/google/uuid"
  "github.com/lrita/cmap"
)

type TCError struct {
  Msg          string
  Attrs        cmap.Cmap
  Err          error
}

func MakeError(err interface{}, args ...interface{}) *TCError {
  if err == nil {
    return nil
  }
  res := new(TCError)
  switch err.(type) {
  case error:
    res.Err   = err.(error)
    res.Msg   = fmt.Sprintf("%v", err)
  case string:
    res.Err   = errors.New(fmt.Sprintf(err.(string), args...))
    res.Msg   = err.(string)
  default:
    return nil
  }
  res.Attrs.Store("id", uuid.NewString())
  res.Attrs.Store("stamp", time.Now())
  return res
}

func (tc *TCstate) MakeError(err interface{}, args ...interface{}) *TCError {
  res := MakeError(err, args...)
  if res == nil {
    return nil
  }
  tc.RegisterError(res)
  return res
}

func (tc *TCstate) RegisterError(err *TCError) {
  tc.errors += 1
  tc.errmsg = err.Msg
  tc.ErrStack.PushBack(err)
}

func (e *TCError) Error() string {
  stamp, _ := e.Attrs.Load("stamp")
  return fmt.Sprintf("Error at %v: %v", stamp.(time.Time).Format(time.RFC3339), e.Msg)
}

func (e *TCError) String() string {
  return e.Error()
}

func (e *TCError) Unix() int64 {
  stamp, _ := e.Attrs.Load("stamp")
  return stamp.(time.Time).Unix()
}
