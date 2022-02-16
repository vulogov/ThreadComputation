package ThreadComputation

import (
  "fmt"
  "errors"
  "github.com/deckarep/golang-set"
  "github.com/gammazero/deque"
)

func CompareFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var e1 interface{}
  var e2 interface{}

  if l.TC.Ready() && q.Len() == 0 && l.TC.Res.Len() >= 2 {
    e1 = l.TC.Get()
    e2 = l.TC.Get()
  } else if l.TC.Ready() && q.Len() == 1 && l.TC.Res.Len() >= 1 {
    e1 = l.TC.Get()
    e2 = q.PopFront()
  } else if q.Len() == 2 {
    e1 = q.PopFront()
    e2 = q.PopFront()
  } else {
    return nil, errors.New("Inconclusive context for comparing")
  }

  fun_name := l.TC.FNStack.Front().(string)
  if fun_name == "=" {
    switch e1.(type) {
    case int64:
      switch e2.(type) {
      case int64:
        return e1.(int64) == e2.(int64), nil
      }
    case float64:
      switch e2.(type) {
      case float64:
        return e1.(float64) == e2.(float64), nil
      }
    case string:
      switch e2.(type) {
      case string:
        return e1.(string) == e2.(string), nil
      }
    case bool:
      switch e2.(type) {
      case bool:
        return e1.(bool) == e2.(bool), nil
      }
    case mapset.Set:
      switch e2.(type) {
      case mapset.Set:
        return e1.(mapset.Set).Equal(e2.(mapset.Set)), nil
      }
    }
  } else if fun_name == "<>" {
    switch e1.(type) {
    case int64:
      switch e2.(type) {
      case int64:
        return e1.(int64) != e2.(int64), nil
      }
    case float64:
      switch e2.(type) {
      case float64:
        return e1.(float64) != e2.(float64), nil
      }
    case string:
      switch e2.(type) {
      case string:
        return e1.(string) != e2.(string), nil
      }
    case bool:
      switch e2.(type) {
      case bool:
        return e1.(bool) != e2.(bool), nil
      }
    case mapset.Set:
      switch e2.(type) {
      case mapset.Set:
        return ! e1.(mapset.Set).Equal(e2.(mapset.Set)), nil
      }
    }
  } else if fun_name == "<" {
    switch e1.(type) {
    case int64:
      switch e2.(type) {
      case int64:
        return e1.(int64) < e2.(int64), nil
      }
    case float64:
      switch e2.(type) {
      case float64:
        return e1.(float64) < e2.(float64), nil
      }
    case string:
      switch e2.(type) {
      case string:
        return e1.(string) < e2.(string), nil
      }
    }
  } else if fun_name == ">" {
    switch e1.(type) {
    case int64:
      switch e2.(type) {
      case int64:
        return e1.(int64) > e2.(int64), nil
      }
    case float64:
      switch e2.(type) {
      case float64:
        return e1.(float64) > e2.(float64), nil
      }
    case string:
      switch e2.(type) {
      case string:
        return e1.(string) > e2.(string), nil
      }
    }
  } else if fun_name == ">=" {
    switch e1.(type) {
    case int64:
      switch e2.(type) {
      case int64:
        return e1.(int64) >= e2.(int64), nil
      }
    case float64:
      switch e2.(type) {
      case float64:
        return e1.(float64) >= e2.(float64), nil
      }
    case string:
      switch e2.(type) {
      case string:
        return e1.(string) >= e2.(string), nil
      }
    }
  } else if fun_name == "<=" {
    switch e1.(type) {
    case int64:
      switch e2.(type) {
      case int64:
        return e1.(int64) <= e2.(int64), nil
      }
    case float64:
      switch e2.(type) {
      case float64:
        return e1.(float64) <= e2.(float64), nil
      }
    case string:
      switch e2.(type) {
      case string:
        return e1.(string) <= e2.(string), nil
      }
    }
  } else {
    return nil, errors.New(fmt.Sprintf("Unknown compare operator: %v", fun_name))
  }
  return nil, errors.New(fmt.Sprintf("Unknown context for compare operator"))
}

func LogicFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {
  var e1 interface{}
  var e2 interface{}

  if l.TC.Ready() && q.Len() == 0 && l.TC.Res.Len() >= 2 {
    e1 = l.TC.Get()
    e2 = l.TC.Get()
  } else if l.TC.Ready() && q.Len() == 1 && l.TC.Res.Len() >= 1 {
    e1 = l.TC.Get()
    e2 = q.PopFront()
  } else if q.Len() == 2 {
    e1 = q.PopFront()
    e2 = q.PopFront()
  } else {
    return nil, errors.New("Inconclusive context for logic operator")
  }

  fun_name := l.TC.FNStack.Front().(string)

  if fun_name == "and" || fun_name == "&&" || fun_name == "∧" {
    switch e1.(type) {
    case bool:
      switch e2.(type) {
      case bool:
        return e1.(bool) && e2.(bool), nil
      }
    }
  } else if fun_name == "or" || fun_name == "||" || fun_name == "∨" {
    switch e1.(type) {
    case bool:
      switch e2.(type) {
      case bool:
        return e1.(bool) || e2.(bool), nil
      }
    }
  }
  return nil, errors.New("Unknown context for logic operator")
}

func LogicNotFunction(l *TCExecListener, q *deque.Deque) (interface{}, error) {

  if q.Len() == 0 && l.TC.Ready() {
    e := l.TC.Get()
    switch e.(type) {
    case bool:
      return ! e.(bool), nil
    }
  } else if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      switch e.(type) {
      case bool:
        l.TC.Res.Set(! e.(bool))
      default:
        return nil, errors.New("Inconclusive context for logic NOT operator")
      }
    }
    return nil, nil
  }
  return nil, errors.New("Unknown context for logic NOT operator")
}

func initStdlibCmp() {
  SetFunction("=", CompareFunction)
  SetFunction("<>", CompareFunction)
  SetFunction(">", CompareFunction)
  SetFunction("<", CompareFunction)
  SetFunction("<=", CompareFunction)
  SetFunction(">=", CompareFunction)
  SetFunction("and", LogicFunction)
  SetFunction("or", LogicFunction)
  SetFunction("&&", LogicFunction)
  SetFunction("||", LogicFunction)
  SetFunction("not", LogicNotFunction)
  SetFunction("¬", LogicNotFunction)

}
