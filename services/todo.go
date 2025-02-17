package services

import (
	"fmt"
	"context"
	"encoding/json"

	"net/http"

	"github.com/go-kit/kit/endpoint"

)


type TodoService interface {
	Func1(string) (string, error)
	Count(string) int
}


type TodoServiceImpl struct{}

func (TodoServiceImpl) Func1(s string) (string, error) {
	return "_" + s + "_", nil
}
func (TodoServiceImpl) Count(s string) int {
	return len(s)
}





type Func1Request struct {
	S string `json:"s"`
}

type Func1Response struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}


func DecodeFunc1Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request Func1Request
	fmt.Printf("%+v\n", r)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}


func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


func MakeFunc1Endpoint(svc TodoService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(Func1Request)
		v, err := svc.Func1(req.S)
		if err != nil {
			return Func1Response{v, err.Error()}, nil
		}
		return Func1Response{v, ""}, nil
	}
}



