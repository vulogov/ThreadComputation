package ThreadComputation

import (
  "gonum.org/v1/gonum/floats"
)



func TCCompareNumbersNumbers(name string, x interface{}, y interface{}) interface{} {
  switch name {
  case "=":
    return floats.Equal(x.(*TCNumbers).N, y.(*TCNumbers).N)
  case ">":
    return x.(*TCNumbers).Len() > y.(*TCNumbers).Len()
  case ">=":
    return x.(*TCNumbers).Len() >= y.(*TCNumbers).Len()
  case "<":
    return x.(*TCNumbers).Len() < y.(*TCNumbers).Len()
  case "<=":
    return x.(*TCNumbers).Len() <= y.(*TCNumbers).Len()
  case "!=":
    return ! floats.Equal(x.(*TCNumbers).N, y.(*TCNumbers).N)
  }
  return nil
}

func init() {
  RegisterCompareCallback(Numbers, Numbers, TCCompareNumbersNumbers)
}
