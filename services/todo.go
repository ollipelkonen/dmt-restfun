package services

import (
	"io"
	"fmt"
	"context"
	"encoding/json"

	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/ollipelkonen/dmt-restfun/repositories"
	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

)


type TodoService interface {
	//GetAllEndpoint()
	//GetAllHandler() *httptransport.Server
	CreateGetAllHandler() *httptransport.Server
	CreateGetByIdHandler() *httptransport.Server
	CreateInsertHandler() *httptransport.Server
	GetAll() ([]repositories.Todo, error)
	GetById(id string) (repositories.Todo, error)
	Post(map[string]string) (repositories.Todo, error)
	Func1(string) (string, error)
	Count(string) int
}

type TodoServiceImpl struct {
	todoRepository repositories.TodoRepositoryImpl
}

type JsonMapInterface struct {
	id string
	data map[string]string
}

func (s TodoServiceImpl) CreateGetAllHandler() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			v, err := s.GetAll()
			return v, err
		},
		httptransport.NopRequestDecoder,
		EncodeResponse,
	)
	return handler
}

func (s TodoServiceImpl) CreateGetByIdHandler() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			v, err := s.GetById(request.(PathIdRequest).Id)
			return v, err
		},
		DecodePathId,
		EncodeResponse,
	)
	return handler
}

func (s TodoServiceImpl) CreateInsertHandler() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			fmt.Println("__ insert ", request)
			v, err := s.Post( request.(JsonMapInterface).data );
			return v, err
		},
		DecodeRequest,
		EncodeResponse,
	)
	return handler
}


func (s TodoServiceImpl) GetAll() ([]repositories.Todo, error) {
	d, _ := s.todoRepository.GetAll()
	return d, nil
}

func (s TodoServiceImpl) Func1(string) (string, error) {
	fmt.Println("___ 1")
	d, _ := s.todoRepository.GetAll()
	fmt.Println("___ d: ", d)
	j, _ := json.Marshal(d)
	return string(j), nil;
}

func (TodoServiceImpl) Count(s string) int {
	return len(s)
}

func (s TodoServiceImpl) GetById(id string) (repositories.Todo, error) {
	d, _ := s.todoRepository.GetById(id)
	return d, nil
}

func (s TodoServiceImpl) Post(data map[string]string) (repositories.Todo, error) {
	//TOO:
	s.todoRepository.Insert( data )
	return repositories.Todo{}, nil
}

type EmptyRequest struct {
}

type PathIdRequest struct {
	Id string
}

type Func1Request struct {
	S string `json:"s"`
}

type Func1Response struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}


func DecodePathId(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	return PathIdRequest{params["id"]}, nil
}

// decode post body from json to map[string]string, get {id} from path if applicable
func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var v JsonMapInterface
	params := mux.Vars(r)
	v.id = params["id"]
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &v.data); err != nil {
		return nil, err
	}
	return v, nil
}

func DecodeFunc1Request(_ context.Context, r *http.Request) (interface{}, error) {
	var request Func1Request
	fmt.Printf("____ decode %+v\n", r)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Printf("__ decode error %+v\n", err)
		return nil, err
	}
	fmt.Printf("__ decoded %+v\n", request)
	return request, nil
}


func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}


func MakeFunc1Endpoint(svc TodoService, todoRepository repositories.TodoRepositoryImpl) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		//todoRepository.GetAll()
		fmt.Println("??")
		v, err := svc.GetAll()
		return v, err
		/*req := request.(Func1Request)
		v, err := svc.Func1(req.S)
		if err != nil {
			return Func1Response{v, err.Error()}, nil
		}
		return Func1Response{v, ""}, nil*/
		//return svc.GetAll(), nil
	}
}

func CreateService(/*rs routing.Service,*/ todoRepository repositories.TodoRepositoryImpl) TodoService {
	impl := &TodoServiceImpl {
		/*routingService: rs,*/
		todoRepository,
	 }
	 return impl
}



