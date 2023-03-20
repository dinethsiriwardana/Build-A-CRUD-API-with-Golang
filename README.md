## Build a CRUD API with Golang
This project provides a simple example of building a CRUD (Create, Read, Update, Delete) API using Golang. We will be using Gorilla Mux as a router for our API.

### Prerequisites

-   Go (version 1.11 or higher)
-   Gorilla Mux package (`go get -u github.com/gorilla/mux`)
-   Postman (for testing API endpoints)

### Getting started

1.  Clone the repository using `git clone https://github.com/dinethsiriwardana/Build-A-CRUD-API-with-Golang.git`
2.  Navigate to the project directory and run `go run main.go`
3.  Open Postman and test the API endpoints listed below.

### API endpoints



| HTTP Method | Endpoint     | Description              |
|-------------|:--------------|--------------------------|
| GET         | /movies      | Get all movies           |
| GET         | /movies/{id} | Get a single movie by ID |
| POST        | /movies      | Add a new movie          |
| PUT         | /movies/{id} | Update an existing movie |
| DELETE      | /movies/{id} | Delete a movie by ID     |

### Data model

The API uses a simple `Movie` data model, which has the following fields:

| Field    | Type            |
|----------|-----------------|
| ID       | string          |
| Isbn     | string          |
| Title    | string          |
| Director | Director object |

The `Director` object has the following fields:

| Field     | Type   |
|-----------|--------|
| Firstname | string |
| Lastname  | string |


### Example requests and responses
 Use Postman for  R&R
 
![Screenshot 2023-03-20 at 9 51 47 PM](https://user-images.githubusercontent.com/91774218/226408926-346d5fa3-9102-46a1-a2c4-7a41f6995e46.png | width=100)




