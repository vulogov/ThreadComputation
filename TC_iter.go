package ThreadComputation
import (

)

type TCIterator struct {
  Type       int
  Gen        interface{}
  Last       interface{}
}

func (tc *TCstate) Iterator(gen interface{}) *TCIterator {
  gfun := GetGenCallback(gen)
  if gfun == nil {
    tc.SetError("Can not ccreate iterator for %T", gen)
    return nil
  }
  res := gfun(gen)
  if res == nil {
    tc.SetError("Generator for %T returned nil", gen)
    return nil
  }
  return res
}

func (i *TCIterator) Next() interface{} {
  nfun := GetNextCallbackByType(i.Type)
  if nfun == nil {
    return nil
  }
  return nfun(i)
}

func (i *TCIterator) Prev() interface{} {
  nfun := GetPrevCallbackByType(i.Type)
  if nfun == nil {
    return nil
  }
  return nfun(i)
}
