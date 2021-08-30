
-- 1

this application uses <b>go-kit</b>


go get github.com/go-kit/kit
go get github.com/go-kit/kit/endpoint
go get github.com/go-kit/kit/log@v0.11.0
go get github.com/go-sql-driver/mysql
go get  "github.com/jmoiron/sqlx"


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


