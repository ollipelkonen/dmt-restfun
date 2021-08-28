package main

import (
	"fmt"
	"net/http"
	"github.com/ollipelkonen/dmt-restfun/services"
	httptransport "github.com/go-kit/kit/transport/http"
)


func main() {
	fmt.Println("test")
	svc := services.TodoServiceImpl{}
	//services.Service()

	funcHandler := httptransport.NewServer(
		services.MakeFunc1Endpoint(svc),
		services.DecodeFunc1Request,
		services.EncodeResponse,
	)

	http.Handle("/todo/", funcHandler)

	http.ListenAndServe(":8080", nil)
}
