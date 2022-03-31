package ThreadComputation

import (
  "gonum.org/v1/gonum/mat"
)



func TCConvertMatrixToList(x interface{}) *TCList {
  switch e := x.(type) {
  case *TCMatrix:
    res := new(TCList)
    m, _ := e.M.Dims()
    for i := 0; i < m; i++ {
      row := mat.Row(nil, i, e.M)
      r := MakeNumbers()
      r.N = row
      res.Add(r)
    }
    return res
  }
  return nil
}

func init() {
  RegisterToListCallback(Matrix, TCConvertMatrixToList)
}
