package repositories

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


type TodoRepository interface {
	init(string)
	GetAll() (*sql.Rows, error)
	GetById(id string) (*sql.Rows, error)
}

type TodoRepositoryImpl struct {
	db	*sql.DB
}

func (repo *TodoRepositoryImpl) init(connectString string) {
	fmt.Println(connectString);

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	repo.db = db
	//defer db.Close()
}

func (repo *TodoRepositoryImpl) GetAll() (*sql.Rows, error) {
	result, err := repo.db.Query("SELECT * FROM restfun;")
	fmt.Printf("%+v\n", result)
	return result, err
}

func (repo *TodoRepositoryImpl) GetById(id string) (*sql.Rows, error) {
	return repo.db.Query("SELECT * FROM restfun WHERE Id=?;", id)
}


func CreateRepository(connectString string) TodoRepositoryImpl {
	rep := TodoRepositoryImpl{}
	rep.init(connectString)
	return rep
}

