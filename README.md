
-- 1

this application uses <b>go-kit</b>


you may have to install these manually<br/>

go get github.com/go-kit/kit<br/>
go get github.com/go-kit/kit/endpoint<br/>
go get github.com/go-kit/kit/log@v0.11.0<br/>
go get github.com/go-sql-driver/mysql<br/>
go get "github.com/jmoiron/sqlx"<br/>
go get "github.com/gorilla/mux"<br/>


All configurations are in <i>settings.json</i>

-- 2

To initialize default test database<br/>

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

Timeformat is RFC 3339 (2021-01-26T20:10:59Z)<br/>


-- 4

test<br/>

Use included Postman-json, or try with curl: <br/>

curl -X GET --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo<br/>
curl -X POST -d '{"priority":11, "name":"task!", "description":"testing todo list"}' --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo<br/>

curl -X PUT -d '{"dueDate":"0001-01-01T00:00:00Z", "name":"task!", "description":"updated"}' --header "Authorization: Bearer fbf566e0bd1747409502db0f"  localhost:8080/todo/15<br/>

curl -X POST -d '{"priority":3, "name":"task!", "description":"testing todo list"}' --header "Authorization: Bearer fbf566e0bd1747409502db0f"  localhost:8080/todo<br/>

curl -X POST -d '{ "name":"task!",  "priority":3, "description":"Remember to test TODO-list", "duedate":"2022-01-20T16:00:00Z"  }' --header "Authorization: Bearer fbf566e0bd1747409502db0f"  localhost:8080/todo<br/>

etc
