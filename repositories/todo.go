package repositories

import (
	"fmt"
	"time"
	_ "github.com/go-sql-driver/mysql"
	//"encoding/json"
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
	//db	*sql.DB
	db	*sqlx.DB
}

func (repo *TodoRepositoryImpl) init(connectString string) {
	fmt.Println("connecting " + connectString);
	repo.db = sqlx.MustConnect("mysql", connectString + "?parseTime=true")
	//defer db.Close()
}

func (repo *TodoRepositoryImpl) GetAll() ([]Todo, error) {
	result := []Todo{}
	err := repo.db.Select(&result, "SELECT * FROM todo")
	if err != nil {
		fmt.Printf("!! error: %+v\n", err)
	}
	/*fmt.Printf("%+v\n", result)
	j, _ := json.Marshal(result)
  fmt.Println(string(j))*/
	return result, err
}

func (repo *TodoRepositoryImpl) GetById(id string) (Todo, error) {
	todo := Todo{}
	err := repo.db.Get(todo, "SELECT * FROM todo WHERE Id=?", id)
	return todo, err
}


func CreateRepository(connectString string) TodoRepositoryImpl {
	rep := TodoRepositoryImpl{}
	rep.init(connectString)
	return rep
}

