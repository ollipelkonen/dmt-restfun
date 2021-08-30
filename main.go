package main

import (
	"fmt"
	"net/http"
	//"context"
	"github.com/ollipelkonen/dmt-restfun/services"
	"github.com/ollipelkonen/dmt-restfun/repositories"
	"github.com/ollipelkonen/dmt-restfun/config"
	//"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)


func authMiddleware(token string) mux.MiddlewareFunc {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(r.Header["Authorization"])>0 {
				if r.Header["Authorization"][0] == ("Bearer " + token) {
					next.ServeHTTP(w, r)
				}
				//TODO: else show error
			}
		})
	}
}


func main() {
	svc := services.TodoServiceImpl{}

	config := config.LoadConfig("settings.json")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)

	todoRepository := repositories.CreateRepository(connectionString)

	fmt.Println("running")
	funcHandler := httptransport.NewServer(
		(services.MakeFunc1Endpoint(svc, todoRepository)),
		services.DecodeFunc1Request,
		services.EncodeResponse,
	)

	r := mux.NewRouter()
	r.Handle("/todo", funcHandler).Methods("GET")
	r.Use(authMiddleware(config.Token));

	http.ListenAndServe(":"+config.Port, r)
}
