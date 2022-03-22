package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/gammazero/deque"
  "github.com/google/uuid"
)



func TCLenFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() > 0 {
    return q.Len(), nil
  } else {
    if l.TC.Ready() {
      return l.TC.Res.Len(), nil
    }
  }
  return 0, nil
}

func TCClrFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.Ready() {
    for l.TC.Ready() {
      l.TC.Get()
    }
    for q.Len() > 0 {
      e := q.PopFront()
      l.TC.Set(e)
    }
  } else if l.TC.HaveAttrs() {
    aq := l.TC.Attrs.Q()
    for aq.Len() > 0 {
      aq.PopFront()
    }
    for q.Len() > 0 {
      e := q.PopFront()
      aq.PushBack(e)
    }
  }
  return nil, nil
}

func TCDupFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.HaveAttrs() {
    aq := l.TC.Attrs.Q()
    for q.Len() > 0 {
      e := q.PopFront()
      fun := GetCloneCallback(e)
      if fun == nil {
        continue
      }
      e1 := fun(e)
      if e1 != nil {
        aq.PushFront(e)
        aq.PushFront(e1)
      }
    }
  } else {
    for q.Len() > 0 {
      e := q.PopFront()
      fun := GetCloneCallback(e)
      if fun == nil {
        continue
      }
      e1 := fun(e)
      if e1 != nil {
        l.TC.Set(e)
        l.TC.Set(e1)
      }
    }
  }
  return nil, nil
}

func TCDropFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if l.TC.HaveAttrs() {
    l.TC.Attrs.Q().PopFront()
  } else if l.TC.Ready() {
    l.TC.Get()
  }
  return nil, nil
}

func ToStackFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return nil, nil
  }
  for q.Len() > 0 {
    e:= q.PopFront()
    log.Debugf("Taking from atts[%v]: %v", q.Len(), e)
    ReturnFromFunction(l, name, e)
  }
  return nil, nil
}

func TCNewStackFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  bname := uuid.NewString()
  AddNewStack(l, bname)
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopFront()
      l.TC.Res.Set(e)
    }
  }
  return nil, nil
}

func TCSStackFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  if q.Len() == 0 {
    return l.TC.ResNames.Front().(string), nil
  }
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case string:
      err := l.TC.PositionStack(e.(string))
      if err != nil {
        return nil, err
      }
    }
  }
  return nil, nil
}

func TCsStackFunction(l *TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  for q.Len() > 0 {
    e := q.PopFront()
    switch e.(type) {
    case string:
      err := l.TC.PositionStack(e.(string))
      if err != nil {
        return nil, err
      }
    }
  }
  return nil, nil
}

func init() {
  SetCommand("len", TCLenFunction)
  SetCommand("stack", ToStackFunction)
  SetCommand("clr", TCClrFunction)
  SetFunction("dup", TCDupFunction)
  SetFunction("^", TCDupFunction)
  SetCommand("drop", TCDropFunction)
  SetCommand(",", TCDropFunction)
  SetCommand("|", TCNewStackFunction)
  SetCommand("S", TCSStackFunction)
  SetFunction("s", TCsStackFunction)
}
