package repositories

/*
	this package contains code to for Todo MySQL repository
*/

import (
	"fmt"
	"time"
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


func CreateRepository(connectString string) TodoRepositoryImpl {
	rep := TodoRepositoryImpl{}
	rep.init(connectString)
	return rep
}

