package ThreadComputation

import (
  "github.com/deckarep/golang-set"
)


func setSetFunc(x mapset.Set, name string, y interface{}) mapset.Set {
  x.Add(y)
  return x
}

func setUnSetFunc(x mapset.Set, name string, y interface{}) mapset.Set {
  x.Remove(y)
  return x
}

func TCSetMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case mapset.Set:
    switch y.(type) {
    case int64, float64, bool, string:
      switch name {
      case "+":
        return setSetFunc(x.(mapset.Set), name, y)
      case "-":
        return setUnSetFunc(x.(mapset.Set), name, y)
      }
    case mapset.Set:
      switch name {
      case "+":
        return x.(mapset.Set).Union(y.(mapset.Set))
      case "-":
        return x.(mapset.Set).Intersect(y.(mapset.Set))
      }
    }
  }
  return nil
}

func init() {
  RegisterMathCallback(Set, Int, TCSetMath)
  RegisterMathCallback(Set, Bool, TCSetMath)
  RegisterMathCallback(Set, Float, TCSetMath)
  RegisterMathCallback(Set, String, TCSetMath)
  RegisterMathCallback(Set, Set, TCSetMath)
}
