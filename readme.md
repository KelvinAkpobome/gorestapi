## Golang Task 4


### Features:

## Steps:
1. Fire up MySql server on your machine.
2. Setup a database schema.
3. Line 14 in model/connect.model.go to should be replaced:
    * `db,err :=sql.Open("mysql", "\<username\>:\<password\>@tcp(localhost:\<port\>)/<database_name>")`


## Endpoints examples:
1. `curl http://localhost:3000/api/v1/books` For a GET request

2. `curl http://localhost:3000/api/v1/books/id` For a DELETE request

