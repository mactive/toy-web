package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")

	body, err := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, body)
	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		// 记住要返回，不然就还会执行后面的代码
		return
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "这是订单")
}


func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/user/create", createUser)
	http.HandleFunc("/order", order)
	http.ListenAndServe(":8080", nil)
}

//type Server interface {
//	Route(pattern string, handlerFunc http.HandlerFunc)
//	Start(address string) error
//}
//
//type sdkHttpServer struct {
//	Name string
//}
