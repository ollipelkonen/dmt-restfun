
-- 1

this application uses <b>go-kit</b>


go get github.com/go-kit/kit
go get github.com/go-kit/kit/endpoint
go get github.com/go-kit/kit/log@v0.11.0
go get github.com/go-sql-driver/mysql
go get "github.com/jmoiron/sqlx"
go get "github.com/gorilla/mux"


-- 2

To initialize default test database

<pre>
CREATE DATABASE restfun;
USE restfun;

CREATE TABLE todo (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT DEFAULT "",
  priority INT DEFAULT 1,
  duedate TIMESTAMP DEFAULT 0,
  completed TINYINT(1) DEFAULT 0,
  completiondate TIMESTAMP DEFAULT 0
);

CREATE USER 'restfun'@'localhost' IDENTIFIED BY 'restfun';
GRANT SELECT, INSERT, UPDATE, DELETE ON restfun.* TO 'restfun'@'localhost';
</pre>


-- 3

the application doesn't have any sanity checks for input values

curl -L -X POST -d '{"dueDate":"2020-01-01 01:11:11", "name":"task!", "description":"testing todo list"}' --header "Authorization: Bearer 1234"  localhost:8080/todo


curl -L -X POST -d '{"priority":3, "name":"task!", "description":"testing todo list"}' --header "Authorization: Bearer 1234"  localhost:8080/todo
