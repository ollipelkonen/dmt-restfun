package main

import (
	"fmt"
	"net/http"
	"github.com/ollipelkonen/dmt-restfun/services"
	"github.com/ollipelkonen/dmt-restfun/repositories"
	"github.com/ollipelkonen/dmt-restfun/config"
	httptransport "github.com/go-kit/kit/transport/http"
)


func main() {
	svc := services.TodoServiceImpl{}

	config := config.LoadConfig("settings.json")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)

	rep := repositories.CreateRepository(connectionString)

	rep.GetAll()

	fmt.Println("running")
	funcHandler := httptransport.NewServer(
		services.MakeFunc1Endpoint(svc),
		services.DecodeFunc1Request,
		services.EncodeResponse,
	)

	http.Handle("/todo/", funcHandler)

	http.ListenAndServe(":"+config.Port, nil)
}
