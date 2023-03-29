## Dependencies

This project requires the following dependencies to be installed:

- Golang ([https://go.dev/](https://go.dev/))
- Gin framewrok (for handling HTTP requests) ([https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin))
- MongoDB Go driver (for connecting to MongoDB database) ([https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver))

## Installation

1. Clone this repository to your local machine
2. Install the dependencies using go modules. Run the following command in the root directory of the project:

   ```go
   cd go && go get ./...
   ```

## Configuration

1. Create a `.env` file in the root directory of the project and add the following variables:

   ```js
   DB_DATABASE=<your_mongodb_database>
   DB_USERNAME=<your_mongodb_username>
   DB_PASSWORD=<your_mongodb_password>
   MONGO_URI=<your_mongodb_uri>
   PORT=<port_number>
   ```

2. Whenever you make a `.env` file update, use the following command to activate the latest values:

   ```bash
   source .env
   ```

## Usage

1. Run the server by running the following command in the root directory of the project:

   ```go
   go run main.go
   ```

   or

   ```go
   go run .
   ```

2. Create a dictionary by sending a POST request to `http://localhost:<port_number>/api/dictionaries`

   Sample request:

   ```json
   {
     "name": "Cigarette",
     "is_organic": false,
     "recyable": false,
     "description": "Lorem ipsum",
     "short_description": "Lorem",
     "uri": "/cigarette",
     "display_image": "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Bouteille.jpg/800px-Bouteille.jpg"
   }
   ```

   Sample response:

   ```json
   {
     "data": {
       "dictionary": {
         "_id": "63e655f9971e6433d30f60b5",
         "name": "Cigarette",
         "recyable": false,
         "is_organic": false,
         "short_description": "Lorem",
         "description": "Lorem ipsum",
         "uri": "/cigarette",
         "lessons": [],
         "quizzes": [],
         "display_image": "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Bouteille.jpg/800px-Bouteille.jpg"
       },
       "result": {
         "InsertedID": "63e655f9971e6433d30f60b5"
       }
     },
     "message": "Successfully created a dictionary"
   }
   ```

3. Retrieve a student by sending a GET request to `http://localhost:<port_number>/api/dictionaries/{id}`:
   Sample response:

   ```json
   {
     "data": {
       "_id": "63f280900ca14a2a55479921",
       "name": "Food Waste",
       "recyable": true,
       "short_description": "Test",
       "description": "Test",
       "uri": "/food-waste",
       "display_image": "https://s3.us-west-2.amazonaws.com/a7ce953b-0a25-49eb-b7f5-6abcf218781e/Untitled.png"
     },
     "message": "Successfully get the dictionary with id 63f280900ca14a2a55479921"
   }
   ```

4. Update a student by sending a PUT request to `http://localhost:<port_number>/api/students/{id}`

   Sample request:

   ```json
   {
     "content": "This is my updated gib!"
   }
   ```

   Sample response:

   ```json
   {
     "id": "5ec53edd34bdcf9feaa65731",
     "content": "This is my updated gib!"
   }
   ```

5. Delete a `Dictionary` by sending a DELETE request to `http://localhost:<port_number>/api/dictionaries/{id}`

   Sample response:

   ```json
   {
     "data": null,
     "message": "Successfully deleted the dictionary with id 63c3c639de070156bad623eb"
   }
   ```

## Folder structure

`sowaste-backend` - root folder

- `go`
  - `api`
  - `docs`
  - `internal`
    - `app` - The point where all our dependencies and logic are collected and run the app. The run method that is called from /cmd.
    - `config` - Initialization of the general app configurations that we wrote in the root of the project.
    - `database` - The files contain methods for interacting with databases.
    - `models` - The structures of database tables.
    - `services` - The entire business logic of the application.
    - `transport` - Here we store http-server settings, handlers, ports, etc.
  - `migrations` - This contains all migrations related to databases, e.g. SQL files.
  - `utils` - This contains all utility functions that are used in the application.
  - `go.mod`
  - `go.sum`
  - `README.md`
