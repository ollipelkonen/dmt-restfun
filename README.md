
-- 1

this application uses <b>go-kit</b>


go get github.com/go-kit/kit
go get github.com/go-kit/kit/endpoint
go get github.com/go-kit/kit/log@v0.11.0
go get github.com/go-sql-driver/mysql

-- 2

To initialize default test database

<pre>
CREATE DATABASE restfun;
USE restfun;

CREATE TABLE restfun (
  Id INT AUTO_INCREMENT PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Description TEXT DEFAULT "",
  Priority INT DEFAULT 1,
  DueDate TIMESTAMP DEFAULT 0,
  Completed TINYINT(1) DEFAULT 0,
  CompletionDate TIMESTAMP DEFAULT 0
);

CREATE USER 'restfun'@'localhost' IDENTIFIED BY 'restfun';
GRANT SELECT, INSERT, UPDATE, DELETE ON restfun.* TO 'restfun'@'localhost';
</pre>


