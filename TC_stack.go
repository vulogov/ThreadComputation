package ThreadComputation

import (
  "github.com/gammazero/deque"
  log "github.com/sirupsen/logrus"
)

func SingleValueFromStackAndAttr(l *TCExecListener, q *deque.Deque) *deque.Deque {
  if l.TC.Attrs.GLen() == 0 {
    if l.TC.Ready() {
      e := l.TC.Get()
      log.Debugf("Get from stack: %v", e)
      q.PushFront(e)
    }
  } else {
    aq := l.TC.Attrs.Q()
    e := aq.PopBack()
    q.PushFront(e)
  }
  return q
}

func AllValuesFromStackAndAttr(l *TCExecListener, q *deque.Deque) *deque.Deque {
  if l.TC.Attrs.GLen() == 0 {
    if l.TC.Ready() {
      for l.TC.Ready() {
        e := l.TC.Get()
        q.PushFront(e)
      }
    }
  } else {
    aq := l.TC.Attrs.Q()
    for aq.Len() > 0 {
      e := aq.PopBack()
      q.PushFront(e)
    }
  }
  return q
}

func SingleValueFromStackOrAttr(l *TCExecListener, q *deque.Deque) *deque.Deque {
  var res deque.Deque
  if q.Len() > 0 {
    return q
  } else {
    if l.TC.Attrs.GLen() == 0 {
      if l.TC.Ready() {
        e := l.TC.Get()
        res.PushFront(e)
        return &res
      }
    } else {
      aq := l.TC.Attrs.Q()
      e := aq.PopBack()
      res.PushFront(e)
      return &res
    }
  }
  return nil
}

func AllValuesFromStackOrAttr(l *TCExecListener, q *deque.Deque) *deque.Deque {
  var res deque.Deque
  if q.Len() > 0 {
    return q
  } else {
    if l.TC.Attrs.GLen() == 0 {
      if l.TC.Ready() {
        for l.TC.Ready() {
          e := l.TC.Get()
          res.PushFront(e)
        }
        return &res
      }
    } else {
      aq := l.TC.Attrs.Q()
      for aq.Len() > 0 {
        e := aq.PopBack()
        res.PushFront(e)
      }
      return &res
    }
  }
  return nil
}

func PushQueueToStack(l *TCExecListener, q *deque.Deque) {
  if q.Len() > 0 {
    for q.Len() > 0 {
      e := q.PopBack()
      l.TC.Set(e)
    }
  }
}

func (l *TCExecListener) IsTrue() bool {
  e := l.TC.Res.Q().Front()
  switch e.(type) {
  case bool:
    e := l.TC.Get()
    return e.(bool)
  }
  return false
}

func (l *TCExecListener) IsFalse() bool {
  e := l.TC.Res.Q().Front()
  switch e.(type) {
  case bool:
    e := l.TC.Get()
    return ! (e.(bool))
  }
  return false
}

func AddNewStack(l *TCExecListener, name string) {
  l.TC.AddNewStack(name)
}

func (tc *TCstate) AddNewStack(name string) {
  log.Debugf("Creating new stack: %v", name)
  if tc.ResN.Contains(name) {
    log.Debugf("Can not create new stack as it is already exists: %v", name)
    return
  }
  chcap, err := GetVariable("tc.Chancapacity")
  if err != nil {
    chcap = 4096
  }
  tc.ResN.Add(name)
  tc.ResNames.PushFront(name)
  tc.Res.Add()
  tc.StackList.Delete(name)
  tc.StackChan.Delete(name)
  tc.StackChan.Store(name, make(chan interface{}, chcap.(int)))
  tc.StackList.Store(name, tc.Res.Q())
}
