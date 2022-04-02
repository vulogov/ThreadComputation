package ThreadComputation

import (
  "gonum.org/v1/gonum/mat"
  log "github.com/sirupsen/logrus"
)

func tcGetNFromMatrix(res *TCIterator, i int64) interface{} {
  m, _ := res.Gen.(*TCMatrix).M.Dims()
  if i < 0 || i >= int64(m) {
    res.Last = nil
    return nil
  }
  log.Debugf("matrix-iterator: getting value at %v", i)
  e  := mat.Row(nil, int(i), res.Gen.(*TCMatrix).M)
  if e == nil {
    res.Last = nil
    return nil
  }
  en := MakeNumbers()
  en.N = e
  res.Last = en
  return en
}

func MatrixGenerator(v interface{}) *TCIterator {
  switch e := v.(type) {
  case *TCMatrix:
    log.Debug("iterator: create iterator for Matrix")
    res := new(TCIterator)
    res.Type = Matrix
    res.Gen  = e
    res.Last = nil
    res.Set("current", int64(0))
    return res
  }
  return nil
}

func MatrixNext(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  out := tcGetNFromMatrix(res, i)
  i = i + 1
  res.Set("current", i)
  return out
}

func MatrixPrev(res *TCIterator) interface{} {
  i := res.Get("current").(int64)
  out := tcGetNFromMatrix(res, i)
  i = i - 1
  res.Set("current", i)
  return out
}


func init() {
  RegisterGenCallback(Matrix, MatrixGenerator)
  RegisterNextCallback(Matrix, MatrixNext)
  RegisterPrevCallback(Matrix, MatrixPrev)
}
