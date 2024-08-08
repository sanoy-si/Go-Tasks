# Project: Task-Manager-Api
The Task Manager API is a secured RESTful API that allows users to manage tasks. The API provides secure endpoints that can be used to perform the CRUD operations on tasks. It also supports registration and logging in by using JWT authentication.

## End-point: Get Tasks
This end point returns a list of all tasks that has been added so far.

**Request**  
No request body is required.

**Headers**

- `Authorization`: The authorization token, which should be a valid JWT token with a role of either "user" or "admin".
    

**Response**  
The response will consist of an array of JSON objects which represent tasks and will have the following keys:

- `id` (number): The unique identifier of the task.
    
- `title` (string): The title of the task.
    
- `description` (string): The description of the task.
    
- `due_date` (string): The due date of the task.
    
- `status` (string): The status of the task.
### Method: GET
>```
>localhost:8080/tasks
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer {valid token}|


### Response: 200
```json
[
    {
        "id": "1",
        "title": "Task 1",
        "description": "Task1 description",
        "due_date": "0001-01-01T00:00:00Z",
        "status": "In Progress"
    }
]
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Task
Get a specific task by passing its id.

**Request**  
No request body is required.

**Headers**

- `Authorization`: The authorization token, which should be a valid JWT token with a role of either "user" or "admin".
    

#### Response

The response will contain a JSON object which represents a task object and will contain the following keys:

- `id` (number): The unique identifier of the task.
    
- `title` (string): The title of the task.
    
- `description` (string): The description of the task.
    
- `due_date` (string): The due date of the task.
    
- `status` (string): The status of the task.
### Method: GET
>```
>localhost:8080/tasks/{id}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer {valid token}|


### Response: 200
```json
{
    "id": "1",
    "title": "Task 1",
    "description": "Task1 description",
    "due_date": "0001-01-01T00:00:00Z",
    "status": "In Progress"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create Task
This endpoint in the Task Manager API is used to create a new taskEndFragment

**Request Body**

The request body should contain a JSON object with the following keys:

- `id` (number): The unique identifier for the task.
    
- `title` (string): The title of the task.
    
- `description` (string): A description of the task.
    
- `due_date` (string): The due date for the task.
    
- `status` (string): The status of the task.
    

**Headers**

- `Authorization`: The authorization token, which should be a valid JWT token with a role of "admin".
    

**Response**

The response will have a status code of 201 with a body that will contain a JSON object that represents the created task object with keys similar with the above.
### Method: POST
>```
>localhost:8080/tasks
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer {valid token}|


### Body (**raw**)

```json

```

### Response: 201
```json
{
    "id": "1",
    "title": "Task 1",
    "description": "Task1 description",
    "due_date": "0001-01-01T00:00:00Z",
    "status": "In Progress"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Update Task
This endpoint Updates an existing task having an id od {id} which is passed as path parameter.

**Request Body**

The request body should contain a JSON object with the following keys:

- `id` (number): The unique identifier for the task.
    
- `title` (string): The title of the task.
    
- `description` (string): A description of the task.
    
- `due_date` (string): The due date for the task.
    
- `status` (string): The status of the task.
    

**Headers**

- `Authorization`: The authorization token, which should be a valid JWT token with a role of "admin".
    

**Response**

The response will have a status code 200 with a body that will contain a JSON object that represents the updated task object with keys similar with the above and the updated values.
### Method: PUT
>```
>localhost:8080/tasks/{id}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer {valid token}|


### Body (**raw**)

```json

```

### Response: 200
```json
{
    "id": "1",
    "title": "Updated title",
    "description": "Updated description",
    "due_date": "0001-01-01T00:00:00Z",
    "status": "Done"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Task
This endpoint in the Task Manager API is used to delete an existing task having an id of {id} which is passed as a path parameter.

**Request**  
No request body is required.

**Headers**

- `Authorization`: The authorization token, which should be a valid JWT token with a role of either "admin".
    

#### Response

The response will be a status code of 200 with a success message.
### Method: DELETE
>```
>localhost:8080/tasks/{id}
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer {valid token}|


### Response: 200
```json
{
    "message: ": "Task Deleted Successfully."
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Register
This end point can be used for registration of users.

**Request**  
The request body should be a JSON object with the following fields:

- `username` (string, required): The username for the new user.
    
- `email` (string, required): The email address for the new user.
    
- `password` (string, required): The password for the new user.
    
- `first_name` (string, required): The first name for the new user.
    
- `last_name` (string, required): The last name for the new user.
    

**Headers**  
No additional headers are required.

**Response**  
The response will be the id of the created user object with a status code of 200.
### Method: POST
>```
>localhost:8080/register
>```
### Body (**raw**)

```json

```

### Response: 201
```json
{
    "insertedID": {
        "InsertedID": "66b47032110ed9fb575574bf"
    }
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Login
This end point can be used for logging in of users.

**Request**  
The request body should be a JSON object with the following fields:

- `username` (string, required): The username for the user.
    
- `password` (string, required): The password for the user.
    

**Headers**  
No additional headers are required.

**Response**  
The response will be a JWT token with a status code of 200.
### Method: POST
>```
>localhost:8080/login
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzX2F0IjoxNzIzMTg3NzY5LCJpc19hZG1pbiI6ZmFsc2UsInVzZXJuYW1lIjoidXNlcl91c2VybmFtZSJ9.rjvVF4GX6EhTw8a3oKnqB-SEBHpC5fHNTV5rJo_TQYE"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Promote User
This end point can be used by admins to promote a non-admin user having a username of {username} to an admin.

**Request**  
No request body is required.

**Headers**

- `Authorization`: The authorization token, which should be a valid JWT token with a role of "admin".
    

#### Response

The response will be a status code of 200 with a success message.
### Method: POST
>```
>localhost:8080/users/{username}/promote
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|bearer {valid jwt token}|


### Response: 200
```json
{
    "message": "success"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

#### This is a .md version of the postman documentation. To see the original postman documetation, [click-here](https://documenter.getpostman.com/view/37380246/2sA3rwLu56).
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
