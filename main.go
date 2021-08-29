package main

import (
	"fmt"
	"net/http"
	"github.com/ollipelkonen/dmt-restfun/services"
	"github.com/ollipelkonen/dmt-restfun/repositories"
	httptransport "github.com/go-kit/kit/transport/http"
)


func main() {
	svc := services.TodoServiceImpl{}

	fmt.Println("database connection")
	rep := repositories.CreateRepository("restfun:restfun@/restfun")
	rep.GetAll()

	fmt.Println("server")
	funcHandler := httptransport.NewServer(
		services.MakeFunc1Endpoint(svc),
		services.DecodeFunc1Request,
		services.EncodeResponse,
	)

	http.Handle("/todo/", funcHandler)

	http.ListenAndServe(":8080", nil)
}
