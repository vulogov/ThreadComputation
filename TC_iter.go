package ThreadComputation
import (
  "fmt"
  "github.com/lrita/cmap"
)

type TCIterator struct {
  Type       int
  Gen        interface{}
  Last       interface{}
  Attrs      cmap.Cmap
}

func (tc *TCstate) Iterator(gen interface{}) *TCIterator {
  gfun := GetGenCallback(gen)
  if gfun == nil {
    tc.SetError("Can not create iterator for %T", gen)
    return nil
  }
  res := gfun(gen)
  if res == nil {
    tc.SetError("Generator for %T returned nil", gen)
    return nil
  }
  return res
}

func (i *TCIterator) String() interface{} {
  return fmt.Sprintf("iterator[ type=(%v)]", TypeToStr(i.Type))
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

func (i *TCIterator) Set(name string, val interface{}) interface{} {
  i.Attrs.Delete(name)
  i.Attrs.Store(name, val)
  return val
}

func (i *TCIterator) Get(name string) interface{} {
  if val, ok := i.Attrs.Load(name); ok {
    return val
  }
  return nil
}
