package repositories

/*
	this package contains code to for Todo MySQL repository
*/

import (
	"fmt"
	"time"
	"strings"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
	_ "encoding/json"
	"github.com/jmoiron/sqlx"
)

type Todo struct {
  Id int
  Name string
  Description string
  Priority int
	DueDate time.Time
  Completed int
  CompletionDate time.Time
}


type TodoRepository interface {
	init(string)
	GetAll() ([]Todo, error)
	GetById(id string) (Todo, error)
	DeleteById(id string) (string, error)
	Insert(map[string]string) (int, error)
	Update(id, data Todo) (string, error)
}

type TodoRepositoryImpl struct {
	db	*sqlx.DB
}

func (repo *TodoRepositoryImpl) init(connectString string) {
	fmt.Println("connecting " + connectString);
	repo.db = sqlx.MustConnect("mysql", connectString)
	//defer db.Close()
}

func (repo *TodoRepositoryImpl) GetAll() ([]Todo, error) {
	result := []Todo{}
	err := repo.db.Select(&result, "SELECT * FROM todo")
	if err != nil {
		fmt.Printf("!! error: %+v\n", err)
	}
	return result, err
}

func (repo *TodoRepositoryImpl) GetById(id string) (Todo, error) {
	todo := Todo{}
	err := repo.db.Get(&todo, "SELECT * FROM todo WHERE id = ?", id)
	fmt.Println("Get By ID " + id, todo)
	return todo, err
}

func (repo *TodoRepositoryImpl) DeleteById(id string) (string, error) {
	_, err := repo.db.Query("DELETE FROM todo WHERE id = ?", id)
	return "", err
}


// try to populate struct Todo with values in map
func MapToTodo(dict map[string]string) (Todo) {
	todo := &Todo{}
	t := reflect.ValueOf(todo).Elem()
	for k, v := range dict {
		fname := strings.Title(k)
		f := t.FieldByName(fname)
		if f.CanSet() && f.Type() == reflect.TypeOf("string") {
			newValue := reflect.ValueOf(v)
			t.FieldByName(fname).Set( newValue )
		}
	}
	return *todo;
}


func (repo *TodoRepositoryImpl) Insert(data map[string]string) (int, error) {
	fields := []string{ "Name", "Description", "Priority", "DueDate", "Complated", "CompletionDate" }
	_ = fields
	keys := []string{};
	values := []string{};
	x := []string{}
	for k, v := range data {
		keys = append(keys, k)
		values = append(values, v)
		x = append(x, ":"+k)
	}
	query := "INSERT INTO todo (" + strings.Join(keys,",") + ") VALUES (" + strings.Join(x,",") + ")";
	fmt.Println("___ query: ", query)
	_, err := repo.db.NamedExec(query, MapToTodo(data));
	return 0, err
/*	k, err := repo.db.NamedExec(`INSERT INTO person (Name, Description, Priority, DueDate, Completed, CompletionDate)
		VALUES (:Name, :Description, :Priority, :DueDate, :Completed, :CompletionDate)`, data)
	//return d.LastInsertId(), err
	fmt.Println("__ insert? ", k, err)*/
	//return 1, err
//	return 1, nil
}

func (repo *TodoRepositoryImpl) Update(id, data Todo) (string, error) {
	//d, err := repo.db.NamedExec(`UPDATE person SET Name=, Description, Priority, DueDate, Completed, CompletionDate)
	//	VALUES (:Name, :Description, :Priority, :DueDate, :Completed, :CompletionDate)`, personStructs)
	//return d, err
	//TODO: find out which parameters exist
	fmt.Println("____ update ", data);
	return "ok", nil
}


func CreateRepository(connectString string) TodoRepositoryImpl {
	rep := TodoRepositoryImpl{}
	rep.init(connectString)
	return rep
}

