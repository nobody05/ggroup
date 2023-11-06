package ggroup

import (
	"reflect"
	"sync"
)

type Group struct {
	Concurrent     int
	concurrentChan chan int
	wg             sync.WaitGroup

	funcEntities []*FuncEntity
	mu           sync.Mutex
}

type FuncEntity struct {
	Obj   any             // func or struct with some method
	Name  string          // method name
	Param []reflect.Value // method params
}

//NewGroup
//concurrent Number of concurrent executions(goroutine)
func NewGroup(concurrent int) *Group {
	return &Group{
		funcEntities:   make([]*FuncEntity, 0, 30),
		concurrentChan: make(chan int, concurrent),
	}
}

func (g *Group) AddFunc(fun *FuncEntity) {
	g.mu.Lock()
	defer func() {
		g.mu.Unlock()
	}()

	g.funcEntities = append(g.funcEntities, fun)
}

func (g *Group) Run() {
	g.mu.Lock()
	defer func() {
		g.mu.Unlock()
		g.clearEntities()
	}()
	for _, funEn := range g.funcEntities {
		f := funEn.Obj
		p := funEn.Param
		n := funEn.Name

		g.concurrentChan <- 1
		g.wg.Add(1)
		go func() {
			defer func() {
				g.wg.Done()
				<-g.concurrentChan
			}()
			f := reflect.ValueOf(f)
			if !f.IsValid() {
				return
			}
			switch f.Kind() {
			case reflect.Func:
				f.Call(p)
				return
			case reflect.Struct:
			case reflect.Pointer:
				if len(n) == 0 {
					return
				}
				m := f.MethodByName(n)
				if !m.IsValid() {
					return
				}
				m.Call(p)
			}
		}()
	}
	g.wg.Wait()
}

func (g *Group) clearEntities() {
	g.mu.Lock()
	defer func() {
		g.mu.Unlock()
	}()
	g.funcEntities = g.funcEntities[:0]
}
