package main


import "net/http"

func main() {
	server := NewHttpServer("test-server")
	server.Route("/", home)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/order", order)
	server.Start("8080")

}

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)

	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
	//panic("implement me")
}

func (s sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
	//panic("implement me")
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{Name: name}
}
