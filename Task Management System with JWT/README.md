# Task-Manager-Api
The Task Manager API is a secured RESTful API that allows users to manage tasks. The API provides secure endpoints that can be used to perform the CRUD operations on tasks. It also supports registration and logging in by using JWT authentication.


# Requirements
### Go Programming Language:
- The service is written in Go, so you'll need to have Go installed on your system. 
The recommended version is Go 1.18 or higher.

- You can download the latest version of Go from the official website: https://golang.org/dl/

### MongoDB Database
- The service uses MongoDB as the underlying database to store task data. You'll need to have a MongoDB server running, either locally or on a remote server.

- You can download and install MongoDB from the official website: https://www.mongodb.com/try/download/community
- Setup the connection by populating the DATABASE_URI field in the .env.template file

### Dependencies
The service uses the following Go dependencies:

- github.com/dgrijalva/jwt-go v3.2.0+incompatible
- github.com/gin-gonic/gin v1.10.0
- github.com/go-playground/validator/v10 v10.20.0
- go.mongodb.org/mongo-driver v1.16.0
- golang.org/x/crypto v0.23.0

- You can install these dependencies by running the following command in your project directory:

```sh
go mod tidy
```
# Running the API
- Clone the repository
```sh 
    git clone https://github.com/sanoy-si/Go-Tasks.git
```
- Navigate to the project directory
```sh
    cd Go-Tasks/"Task Management System with MongoDb"
```
- Install dependencies
```sh
 go mod tidy  
```
- Run the main.go file.
```sh
go run .
```

This will start the server at localhost:8080.
## End Points
- The end points are explained in docs/api_documentaion.
