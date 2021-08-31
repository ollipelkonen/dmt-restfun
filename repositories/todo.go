package repositories

/*
	this package contains code to for Todo MySQL repository
*/

import (
	"fmt"
	"time"
	"strings"
	"strconv"
	"reflect"
	_ "github.com/go-sql-driver/mysql"
	_ "encoding/json"
	"github.com/jmoiron/sqlx"
)

// database model
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
	Insert(map[string]interface{}) (string, error)
	Update(id string, data map[string]interface{}) (string, error)
}

type TodoRepositoryImpl struct {
	db	*sqlx.DB
}

// create database connection
func (repo *TodoRepositoryImpl) init(connectString string) {
	fmt.Println("connecting " + connectString);
	repo.db = sqlx.MustConnect("mysql", connectString)
	//defer repo.db.Close()
}


// helper function: try to populate struct Todo with values in map
func MapToTodo(dict map[string]interface{}) (Todo) {
	todo := &Todo{}
	t := reflect.ValueOf(todo).Elem()
	for k, v := range dict {
		fname := strings.Title(k)
		f := t.FieldByName(fname)
		if f.CanSet() {
			val := reflect.ValueOf(v)
			f.Set(val.Convert(f.Type()))
		} else {
		}
	}
	return *todo;
}

// helper function: return keys in two arrays for sql query, one with names and one with :name
func parseMapForQuery(data map[string]interface{}) ([]string, []string) {
	keys := []string{};
	x := []string{}
	for k, _ := range data {
		keys = append(keys, k)
		x = append(x, ":"+k)
	}
	return keys, x
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
	return "ok", err
}

func (repo *TodoRepositoryImpl) Insert(data map[string]interface{}) (string, error) {
	keys, x := parseMapForQuery(data);
	query := "INSERT INTO todo (" + strings.Join(keys,",") + ") VALUES (" + strings.Join(x,",") + ")";
	values := MapToTodo(data)
	fmt.Println("___ query: ", query, values)
	_, err := repo.db.NamedExec(query, values);
	return "ok", err
}

func (repo *TodoRepositoryImpl) Update(id string, data map[string]interface{}) (string, error) {
	// create list of "name=:name, desc=:desc..." for qyery
	keys, x := parseMapForQuery(data);
	vals := []string{};
	for k,v := range keys {
		vals = append( vals, v + "=" + x[k] );
	}
	query := "UPDATE todo SET " + strings.Join(vals,",") + " WHERE id=:id";
	values := MapToTodo(data)
	values.Id, _ = strconv.Atoi(id)
	_, err := repo.db.NamedExec(query, values);
	return "ok", err
}


func CreateRepository(connectString string) TodoRepositoryImpl {
	rep := TodoRepositoryImpl{}
	rep.init(connectString)
	return rep
}

