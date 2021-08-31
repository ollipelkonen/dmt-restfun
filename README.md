
<h1> About </h1>

This application has been developed using go version go1.17 linux/amd64</br>

Used Timeformat is RFC 3339 (2021-01-26T20:10:59Z)<br/>


<h1>Database</h1>

To create default database<br/>

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

Don't worry, if table doesn't exist, the application tries to create it (database and user must exist)

<h1>Run</h1>

All configurations are in <i>settings.json</i>

There you can find required Bearer-token and port that the application listens to, and database settings</br>

</br>
you may have to install these manually<br/>

go get github.com/go-kit/kit<br/>
go get github.com/go-kit/kit/endpoint<br/>
go get github.com/go-kit/kit/log@v0.11.0<br/>
go get github.com/go-sql-driver/mysql<br/>
go get github.com/jmoiron/sqlx<br/>
go get github.com/gorilla/mux<br/>

</br>

Run:</br>
go run main.go</br>


<h1>Test</h1>

Use included Postman-json, or try with curl: <br/>

<pre>curl -X GET --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo</pre>
<pre>curl -X POST -d '{"priority":11, "name":"task!", "description":"testing todo list"}' --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo</pre>
<pre>curl -X PUT -d '{"dueDate":"0001-01-01T00:00:00Z", "name":"task!", "description":"updated"}' --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo/1</pre>
<pre>curl -X POST -d '{"name":"task!", "priority":3, "description":"Remember to test TODO-list", "duedate":"2022-01-20T16:00:00Z"}' --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo</pre>
<pre>curl -X GET --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo/1</pre>
<pre>curl -X DELETE --header "Authorization: Bearer fbf566e0bd1747409502db0f" localhost:8080/todo/1</pre>

etc
