package server

import (
	. "gorest/handler"
	"fmt"
	"reflect"
)

type Router struct {
	handlerMap map[string]invoker
}

type invoker struct {
	method string
	handler Handler
}


func (this *Router) Register(url string, handler Handler, method string) {
	if (url == "") {
		fmt.Println("register url is nil")
		return
	}
	this.handlerMap = map[string]invoker{}
	this.handlerMap[url] = invoker{method, handler}
}

func (this invoker) Invoke() interface{} {
	reflectVal := reflect.ValueOf(this.handler)
	in := make([]reflect.Value, 0)
	val := reflectVal.MethodByName(this.method).Call(in)
	return val[0].Interface()
}