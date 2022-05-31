package ThreadComputation

import (

)

func TCDataFrameCar(v interface{}) interface{} {
  switch v.(type) {
  case *TCData:
    v.(*TCData).D.Lock()
    df := v.(*TCData).D.Copy()
    v.(*TCData).D.Unlock()
    for n := 1; n < v.(*TCData).D.NRows(); n++ {
      df.Remove(n)
    }
    out := MakeData()
    out.D = df
    return out
  }
  return nil
}

func TCDataFrameCdr(v interface{}) interface{} {
  switch v.(type) {
  case *TCData:
    v.(*TCData).D.Lock()
    df := v.(*TCData).D.Copy()
    v.(*TCData).D.Unlock()
    df.Remove(0)
    out := MakeData()
    out.D = df
    return out
  }
  return nil
}

func init() {
  RegisterFunctionCallback("car", Data, TCDataFrameCar)
  RegisterFunctionCallback("cdr", Data, TCDataFrameCdr)
}
