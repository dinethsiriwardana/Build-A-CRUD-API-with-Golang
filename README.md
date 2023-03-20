# Build a CRUD API with Golang


This project provides a simple example of building a CRUD (Create, Read, Update, Delete) API using Golang based on a tutorial by `freeCodeCamp.org`. 
We will be using Gorilla Mux as a router for our API.



[![IMAGE ALT TEXT HERE](https://user-images.githubusercontent.com/91774218/226343585-b8ac5425-d720-4a7b-9888-64b713980296.jpg)](https://youtu.be/jFfo23yIWac?t=1235)


### Prerequisites

-   Go (version 1.11 or higher)
-   Gorilla Mux package `go get -u github.com/gorilla/mux`
-   Postman (for testing API endpoints)

### Getting started

1.  Clone the repository using 
  ```sh 
  git clone https://github.com/dinethsiriwardana/Build-A-CRUD-API-with-Golang.git
  ```
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
 
 - Get all
 
![Screenshot 2023-03-20 at 9 51 47 PM](https://user-images.githubusercontent.com/91774218/226410703-a7eaa278-3818-46cf-a54e-d435e8fdf0af.png)

- Get by ID

![Screenshot 2023-03-20 at 9 51 54 PM](https://user-images.githubusercontent.com/91774218/226410752-757336cb-de45-4d78-9098-055d8e189451.png)

- Create

![Screenshot 2023-03-20 at 9 51 58 PM](https://user-images.githubusercontent.com/91774218/226410942-cd4a4764-81c2-4b90-9e96-b33b57b6734b.png)

- Update

![Screenshot 2023-03-20 at 9 52 01 PM](https://user-images.githubusercontent.com/91774218/226411018-c5a139a6-7cb9-4d5e-9ec8-dfef9df93113.png)

- Delete

![Screenshot 2023-03-20 at 9 52 04 PM](https://user-images.githubusercontent.com/91774218/226411087-c2d7bfdc-ba53-4e7b-a284-372904244953.png)




