package main


import "net/http"

func main() {

}

type Server interface {
	// Route
	Route(pattern string, handlerFunc http.HandlerFunc)
	
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	panic("implement me")
}

func (s sdkHttpServer) Start(address string) error {
	panic("implement me")
}
