package main

import (
	"fmt"
	"net/http"
	//"context"
	"github.com/ollipelkonen/dmt-restfun/services"
	"github.com/ollipelkonen/dmt-restfun/repositories"
	"github.com/ollipelkonen/dmt-restfun/config"
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
	config := config.LoadConfig("settings.json")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)

	todoRepository := repositories.CreateRepository(connectionString)
	svc := services.CreateService(todoRepository)

	fmt.Println("____ running")
	/*funcHandler := httptransport.NewServer(
		services.MakeFunc1Endpoint(svc, todoRepository),
		httptransport.NopRequestDecoder,
		services.EncodeResponse,
	)*/

	r := mux.NewRouter()

	r.Handle("/todo", svc.CreateGetAllHandler()).Methods("GET")
	r.Handle("/todo/{id}", svc.CreateGetByIdHandler()).Methods("GET");
	r.Handle("/todo", svc.CreateInsertHandler()).Methods("POST")
	r.Handle("/todo/{id}", svc.CreateUpdateHandler()).Methods("PUT")

	r.Use(authMiddleware(config.Token));
	http.ListenAndServe(":"+config.Port, r)
}
