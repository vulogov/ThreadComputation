package ThreadComputation

import (

)

func tcMatrixMathInt(name string, x *TCMatrix, y int64) *TCMatrix {
  m,n := x.M.Dims()
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      switch name {
      case "+":
        x.M.Set(i, j, x.M.At(i,j) + float64(y))
      case "-":
        x.M.Set(i, j, x.M.At(i,j) - float64(y))
      case "*":
        x.M.Set(i, j, x.M.At(i,j) * float64(y))
      case "/":
        x.M.Set(i, j, x.M.At(i,j) / float64(y))
      }
    }
  }
  return x
}

func tcMatrixMathFloat(name string, x *TCMatrix, y float64) *TCMatrix {
  m,n := x.M.Dims()
  for i := 0; i < m; i++ {
    for j := 0; j < n; j++ {
      switch name {
      case "+":
        x.M.Set(i, j, x.M.At(i,j) + y)
      case "-":
        x.M.Set(i, j, x.M.At(i,j) - y)
      case "*":
        x.M.Set(i, j, x.M.At(i,j) * y)
      case "/":
        x.M.Set(i, j, x.M.At(i,j) / y)
      }
    }
  }
  return x
}

func tcMatrixMathMatrix(name string, x *TCMatrix, y *TCMatrix) *TCMatrix {
  m1,n1 := x.M.Dims()
  m2,n2 := y.M.Dims()
  if m1 != m2 || n1 != n2 {
    return nil
  }
  switch name {
  case "+":
    x.M.Add(x.M, y.M)
  case "-":
    x.M.Sub(x.M, y.M)
  case "*":
    x.M.MulElem(x.M, y.M)
  case "/":
    x.M.DivElem(x.M, y.M)
  }
  return x
}

func TCMatrixMath(name string, x interface{}, y interface{}) interface{} {
  switch x.(type) {
  case *TCMatrix:
    switch y.(type) {
    case int64:
      return tcMatrixMathInt(name, x.(*TCMatrix), y.(int64))
    case float64:
      return tcMatrixMathFloat(name, x.(*TCMatrix), y.(float64))
    case *TCMatrix:
      return tcMatrixMathMatrix(name, x.(*TCMatrix), y.(*TCMatrix))
    }
  }
  return nil
}



func init() {
  RegisterMathCallback(Matrix, Int, TCMatrixMath)
  RegisterMathCallback(Matrix, Float, TCMatrixMath)
  RegisterMathCallback(Matrix, Matrix, TCMatrixMath)
}
