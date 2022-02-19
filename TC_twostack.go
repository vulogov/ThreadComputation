package ThreadComputation

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/gammazero/deque"
	"github.com/google/uuid"
	"github.com/lrita/cmap"
)

type TwoStack struct {
	R      deque.Deque
	C      cmap.Cmap
	tmp    cmap.Cmap
	ID     string
	Status bool
	Mode   bool
}

const minCap = 1024

//
// Initialize new 2-dimentional stack structure
//
func InitTS() *TwoStack {
	ts := &TwoStack{
		Status: true,
		Mode:   true,
		ID:     uuid.New().String(),
	}
	return ts
}

func (ts *TwoStack) Left() {
	ts.R.Rotate(-1)
}

func (ts *TwoStack) CLeft() {
	q := ts.Q()
	q.Rotate(-1)
}

func (ts *TwoStack) Right() {
	ts.R.Rotate(1)
}

func (ts *TwoStack) CRight() {
	q := ts.Q()
	q.Rotate(1)
}

func (ts *TwoStack) Normal() {
	ts.Mode = true
}

func (ts *TwoStack) Reverse() {
	ts.Mode = false
}


func (ts *TwoStack) PopFront() interface{} {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New(0, minCap))
		return nil
	}
	return ts.R.Front().(*deque.Deque).PopFront()
}

func (ts *TwoStack) Front() interface{} {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New(0, minCap))
	}
	if ts.R.Front().(*deque.Deque).Len() > 0 {
		return ts.R.Front().(*deque.Deque).Front()
	} else {
		return nil
	}
}

func (ts *TwoStack) Set(data interface{}) {
	if ts.R.Len() == 0 {
		ts.R.PushBack(deque.New(0, minCap))
	}
	if ts.Mode == true {
		ts.R.Back().(*deque.Deque).PushBack(data)
	} else {
		ts.R.Front().(*deque.Deque).PushFront(data)
	}
}

func (ts *TwoStack) Get() (interface{}, error) {
	log.Debugf("Getting from twostack len=%v, mode=%v, len=%v glen=%v", ts.R.Len(), ts.Mode, ts.Len(), ts.GLen())
	if ts.R.Len() == 0 {
		return nil, fmt.Errorf("Stack is to shallow for .Get() operation")
	}
	if ts.Mode == true {
		if ts.R.Back().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Back().(*deque.Deque).Back(), nil
	} else {
		if ts.R.Front().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Front().(*deque.Deque).Front(), nil
	}
}

func (ts *TwoStack) Take() (interface{}, error) {
	log.Debugf("Taking from twostack len=%v, mode=%v", ts.R.Len(), ts.Mode)
	if ts.R.Len() == 0 {
		return nil, fmt.Errorf("Stack is to shallow for .Take() operation")
	}
	if ts.Mode == true {
		if ts.R.Back().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Back().(*deque.Deque).PopBack(), nil
	} else {
		if ts.R.Front().(*deque.Deque).Len() == 0 {
			return nil, fmt.Errorf("Stack(Cell) is to shallow for .Take() operation")
		}
		return ts.R.Front().(*deque.Deque).PopFront(), nil
	}
}

func (ts *TwoStack) Add() {
	log.Debugf("Adding new twostack mode=%v glen=%v", ts.Mode, ts.GLen())
	if ts.Mode == true {
		ts.R.PushBack(deque.New(0, minCap))
	} else {
		ts.R.PushFront(deque.New(0, minCap))
	}
}

func (ts *TwoStack) addQ(q *deque.Deque) {
	log.Debugf("Adding custom twostack mode=%v", ts.Mode)
	if ts.Mode == true {
		ts.R.PushBack(q)
	} else {
		ts.R.PushFront(q)
	}
}

func (ts *TwoStack) Del() {
	log.Debugf("Delete twostack mode=%v glen=%v", ts.Mode, ts.GLen())
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			ts.R.PopBack()
		} else {
			ts.R.PopFront()
		}
	}
}

func (ts *TwoStack) GLen() int {
	return ts.R.Len()
}

func (ts *TwoStack) Len() int {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			return ts.R.Back().(*deque.Deque).Len()
		} else {
			return ts.R.Front().(*deque.Deque).Len()
		}
	}
	return 0
}

func (ts *TwoStack) Q() *deque.Deque {
	if ts.R.Len() > 0 {
		if ts.Mode == true {
			return ts.R.Back().(*deque.Deque)
		} else {
			return ts.R.Front().(*deque.Deque)
		}
	}
	return nil
}
