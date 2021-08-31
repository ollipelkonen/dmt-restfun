package services

import (
	"io"
	_ "fmt"
	"context"
	"encoding/json"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/ollipelkonen/dmt-restfun/repositories"
	"github.com/gorilla/mux"

)


type TodoService interface {
	CreateGetAllEndpoint() *httptransport.Server
	CreateGetByIdEndpoint() *httptransport.Server
	CreateInsertEndpoint() *httptransport.Server
	CreateUpdateEndpoint() *httptransport.Server
	CreateDeleteEndpoint() *httptransport.Server
}

type TodoServiceImpl struct {
	todoRepository repositories.TodoRepositoryImpl
}


type JsonMapInterface struct {
	id string
	data map[string]interface{}
}

type EmptyRequest struct {
}

type PathIdRequest struct {
	Id string
}



func (s TodoServiceImpl) CreateGetAllEndpoint() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			return s.todoRepository.GetAll()
		},
		httptransport.NopRequestDecoder,
		EncodeResponse,
	)
	return handler
}

func (s TodoServiceImpl) CreateGetByIdEndpoint() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			return s.todoRepository.GetById(request.(PathIdRequest).Id)
		},
		DecodePathId,
		EncodeResponse,
	)
	return handler
}

func (s TodoServiceImpl) CreateInsertEndpoint() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			return s.todoRepository.Insert( request.(JsonMapInterface).data );
		},
		DecodeRequest,
		EncodeResponse,
	)
	return handler
}

func (s TodoServiceImpl) CreateUpdateEndpoint() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			return s.todoRepository.Update( request.(JsonMapInterface).id, request.(JsonMapInterface).data )
		},
		DecodeRequest,
		EncodeResponse,
	)
	return handler
}

func (s TodoServiceImpl) CreateDeleteEndpoint() *httptransport.Server {
	handler := httptransport.NewServer(
		func(_ context.Context, request interface{}) (interface{}, error) {
			return s.todoRepository.DeleteById( request.(PathIdRequest).Id )
		},
		DecodePathId,
		EncodeResponse,
	)
	return handler
}




// get {id} from path if applicable
func DecodePathId(_ context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	return PathIdRequest{params["id"]}, nil
}

// decode post body from json to map[string]interface{}, get {id} from path if applicable
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


func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func CreateService(/*rs routing.Service,*/ todoRepository repositories.TodoRepositoryImpl) TodoService {
	impl := &TodoServiceImpl {
		/*routingService: rs,*/
		todoRepository,
	 }
	 return impl
}



