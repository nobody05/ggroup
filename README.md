# ggroup
## concurrent execution function

### Install

```go
go get github.com/nobody05/ggroup
```

### Example

```go

type Goods struct {
}

func (gs *Goods) GetInfo(id int) string {
	log.Printf("goods.getInfo.id %d", id)
	return "goodsInfo"
}


gg := NewGroup(5)
// add func 
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
    Obj: func(name string) {
        log.Println("hello ", name)
    },
    Param: []reflect.Value{reflect.ValueOf("jack")},
})
// add Goods struct method
gg.AddFunc(&FuncEntity{
    Obj:   &Goods{},
    Name:  "GetInfo",
    Param: []reflect.Value{reflect.ValueOf(1)},
})
gg.Run()

```