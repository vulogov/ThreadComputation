package ThreadComputation

import (
  "fmt"
  "math"
  "gonum.org/v1/gonum/mat"
  "github.com/gammazero/deque"
)

type TCMatrix struct {
  P            float64
  M           *mat.Dense
}

func (tc *TCstate) Matrix(q *deque.Deque) *TCMatrix {

  d := make([]float64, 0)
  if q.Len() == 0 {
    res := new(TCMatrix)
    res.P = 0.0
    res.M = new(mat.Dense)
    return res
  }
  x := q.Len()
  y := 0
  for i := 0; i < q.Len(); i++ {
    e := q.At(i)
    switch e.(type) {
    case *TCNumbers:
      if y < e.(*TCNumbers).Len() {
        y = e.(*TCNumbers).Len()
      }
    case *TCList:
      return tc.Matrix(&e.(*TCList).Q)
    default:
      tc.SetError("Can not crreate matrix[] due to invalid source data")
      return nil
    }
  }
  for i := 0; i < q.Len(); i++ {
    e := q.At(i).(*TCNumbers)
    for j := 0; j < e.Len(); j++ {
      d = append(d, e.N[j])
    }
    if y > e.Len() {
      for j := 0; j < (y-e.Len()); j++ {
        d = append(d, math.NaN())
      }
    }
  }
  res := new(TCMatrix)
  res.P = 0.0
  res.M = mat.NewDense(x, y, d)
  return res
}

func (m *TCMatrix) String() string {
  m1, n1 := m.M.Dims()
  if m1 == 0 && n1 == 0 {
    return fmt.Sprintf("matrix[ 0 x 0 ]")
  }
  out := mat.Formatted(m.M, mat.Prefix(""), mat.Squeeze())
  return fmt.Sprintf("%v", out)
}

func (m *TCMatrix) RotateLeft() *TCMatrix {
  d := make([]float64, 0)
  m1, n1 := m.M.Dims()
  for i := n1-1; i >= 0; i-- {
    v := m.M.ColView(i)
    for j := 0; j < v.Len(); j++ {
      e := v.AtVec(j)
      d = append(d, e)
    }
  }
  res := new(TCMatrix)
  res.P = 0.0
  res.M = mat.NewDense(n1, m1, d)
  return res
}

func (m *TCMatrix) RotateRight() *TCMatrix {
  d := make([]float64, 0)
  m1, n1 := m.M.Dims()
  for i := 0; i < n1; i++ {
    v := m.M.ColView(i)
    for j := v.Len()-1; j >= 0; j-- {
      e := v.AtVec(j)
      d = append(d, e)
    }
  }
  res := new(TCMatrix)
  res.P = 0.0
  res.M = mat.NewDense(n1, m1, d)
  return res
}

func TCMatrixFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  m := l.TC.Matrix(q)
  return m, nil
}


func init() {
  SetCommand("matrix", TCMatrixFunction)
}
