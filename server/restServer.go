package server

import (
	"fmt"
	"net/http"
	"gorest/model"
	"encoding/json"
	"gorest/handler"
)

type RestServer struct {
	port string
	router Router
}

func (this *RestServer) SetPort(port string)  (*RestServer){
	this.port = port
	return this
}

func (this *RestServer) AddRouter(url string, handler handler.Handler, method string)  *RestServer{
	this.router.Register(url, handler, method)
	return this
}

func (this *RestServer) Start()  {
	http.HandleFunc("/", this.dispatcher)
	http.ListenAndServe(":"+this.port, nil)
	fmt.Println("rest server start, port:", this.port)
}


func (this *RestServer) dispatcher(w http.ResponseWriter, req *http.Request)  {
	header := w.Header()
	header.Add("Content-Type", "application/json;charset=utf-8")

	result := this.exec(req)
	template := model.HttpResponseTemplate{Code : 200, Msg : "ok", Data: result}

	enc := json.NewEncoder(w)
	enc.Encode(template)
}

func (this *RestServer) exec(req *http.Request) interface{} {
	// get method from router to invoke
	return this.router.handlerMap[req.URL.String()].Invoke()
}
