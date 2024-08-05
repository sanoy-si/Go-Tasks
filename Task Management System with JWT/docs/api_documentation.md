# Task-Manager-Api
The Task Manager API is a RESTful API that allows users to manage tasks. The API provides endpoints that can be used to perform the CRUD operations.


## End-point: Get Tasks
This end point returns a list of all tasks that has been added so far.

**Request**

No request body is required.

#### Response

The response will consist an array of JSON objects which represents task and will have the following keys:

- `id` (number): The unique identifier of the task.
    
- `title` (string): The title of the task.
    
- `description` (string): The description of the task.
    
- `due_date` (string): The due date of the task.
    
- `status` (string): The status of the task.



## Example 
### Method: GET
>```
>localhost:8080/tasks
>```
### Response: 200
```json
[
    {
        "id": "1",
        "title": "Task 1",
        "description": "First task",
        "due_date": "2024-08-02T09:34:58.4436435+03:00",
        "status": "Pending"
    },
    {
        "id": "2",
        "title": "Task 2",
        "description": "Second task",
        "due_date": "2024-08-03T09:34:58.4436435+03:00",
        "status": "In Progress"
    },
    {
        "id": "3",
        "title": "Task 3",
        "description": "Third task",
        "due_date": "2024-08-04T09:34:58.4451473+03:00",
        "status": "Completed"
    }
]
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Task
Get a specific task by passing its id.

#### Request

No request body is required.

#### Response

The response will contain a JSON object which represents a task object and will contain the following keys:

- `id` (number): The unique identifier of the task.
    
- `title` (string): The title of the task.
    
- `description` (string): The description of the task.
    
- `due_date` (string): The due date of the task.
    
- `status` (string): The status of the task.

## Example 
### Method: GET
>```
>localhost:8080/tasks/{id}
>```
### Response: 200
```json
{
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2024-08-02T09:34:58.4436435+03:00",
    "status": "Pending"
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
    

**Response**

The response will have a status code of 201 with a body that will contain a JSON object that represents the created task object with keys similar with the above.


## Example 
### Method: POST
>```
>localhost:8080/tasks
>```
### Body (**raw**)

```json

```

### Response: 201
```json
{
    "id": "4",
    "title": "Task 4",
    "description": "Fourth task",
    "due_date": "2024-08-02T09:34:58.4436435+03:00",
    "status": "Pending"
}
```

### Response: 201
```json
{
    "id": "4",
    "title": "Task 4",
    "description": "Fourth task",
    "due_date": "2024-08-02T09:34:58.4436435+03:00",
    "status": "Pending"
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
    

**Response**

The response will have a status code 200 with a body that will contain a JSON object that represents the updated task object with keys similar with the above and the updated values.


## Example 
### Method: PUT
>```
>localhost:8080/tasks/{id}
>```
### Body (**raw**)

```json

```

### Response: 200
```json
{
    "id": "1",
    "title": "Updated Task",
    "description": "Updated description",
    "due_date": "2024-08-20T15:00:00Z",
    "status": "Done"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Task
This endpoint in the Task Manager API is used to delete an existing task having an id of {id} which is passed as a path parameter.

## Example 
### Method: DELETE
>```
>localhost:8080/tasks/{id}
>```
### Response: 200
```json
{
    "message: ": "Task Deleted Successfully."
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________

#### This is a .md version of the postman documentation. To see the original postman documetation, [click-here](https://documenter.getpostman.com/view/37380246/2sA3rwLu56).
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
