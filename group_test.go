package ggroup

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func MyP() {
	log.Println("my p")
}

type Goods struct {
}

func (gs *Goods) GetInfo(id int) string {
	log.Printf("goods.getInfo.id %d", id)
	return "goodsInfo"
}

func (gs *Goods) GetList(page, pageSize int) string {
	log.Printf("goods.GetList page: %d pageSize: %d", page, pageSize)
	return "GetList"
}

func TestGroup(t *testing.T) {
	gg := NewGroup(5)
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			log.Println("hello")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			time.Sleep(time.Second)
			log.Println("hello1")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: MyP,
	})
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			log.Println("hello2")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			time.Sleep(time.Second)
			log.Println("hello3")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			log.Println("hello4")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			time.Sleep(time.Second)
			log.Println("hello5")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: func() {
			log.Println("hello6")
		},
	})
	gg.AddFunc(&FuncEntity{
		Obj: func(name string) {
			log.Println("hello ", name)
		},
		Param: []reflect.Value{reflect.ValueOf("jack")},
	})
	gg.AddFunc(&FuncEntity{
		Obj:   &Goods{},
		Name:  "GetInfo",
		Param: []reflect.Value{reflect.ValueOf(1)},
	})
	gg.AddFunc(&FuncEntity{
		Obj:   &Goods{},
		Name:  "GetList",
		Param: []reflect.Value{reflect.ValueOf(1), reflect.ValueOf(10)},
	})
	gg.Run()
}
