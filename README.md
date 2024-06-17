# Gin API Project

This project is a RESTful API implemented using the Gin framework in Go (Golang). It provides endpoints to perform CRUD operations on `Tags` data stored in a database.

## Features

- Create, Read, Update, and Delete operations for `Tags`.
- Integration with GORM for database operations.
- Error handling using custom helper functions.
- Implementation of a repository pattern for separation of concerns.

## Technologies Used

- **Gin**: Web framework for building APIs in Go.
- **GORM**: ORM library for database interaction.
- **Go**: Programming language used for development.
- **GitHub**: Version control and collaboration using Git.

## Project Structure

The project structure is organized as follows:

```
gin-api/
│
├── controllers/
│   └── tags_controller.go
│
├── data/
│ ├── request/
│ │ ├── create_tags_request.go
│ │ └── update_tags_request.go
│ │
│ ├── response/
│ ├── tags_response.go
│ └── web_response.go
│
├── helper/
│  └── error.go
│
├── model/
│ └── tags.go
│
├── repository/
│ ├── tags_repository.go
│ └── tags_repository_impl.go
│
├── service/
│ ├── tags_service.go
│ └──  tags_service_impl.go
│
├── README.md
├── go.mod
├── go.sum
└──  main.go
```

### Folder Descriptions

- **controllers**: Contains handler functions to process incoming HTTP requests and manage API endpoints.

  - `tags_controller.go`: Manages endpoints related to tags, such as creating, updating, and retrieving tags.

- **data**: Contains data transfer objects (DTOs) for handling requests and responses.

  - **request**: DTOs for incoming requests.
    - `create_tags_request.go`: Defines the structure for creating tags requests.
    - `update_tags_request.go`: Defines the structure for updating tags requests.
  - **response**: DTOs for outgoing responses.
    - `tags_response.go`: Defines the structure for tags response data.
    - `web_response.go`: Defines a general structure for web responses.

- **helper**: Contains utility functions and helpers.

  - `error.go`: Handles error formatting and propagation.

- **model**: Contains the data models representing the entities in the application.

  - `tags.go`: Defines the data model for tags.

- **repository**: Contains the repository interfaces and implementations for data persistence.

  - `tags_repository.go`: Defines the interface for tags repository.
  - `tags_repository_impl.go`: Implements the tags repository interface.

- **service**: Contains the service interfaces and implementations for business logic.

  - `tags_service.go`: Defines the interface for tags service.
  - `tags_service_impl.go`: Implements the tags service interface.

- **main.go**: The entry point of the application.

- **go.mod**: Defines the module and its dependencies.

- **go.sum**: Contains the checksum of the module dependencies.

This structure ensures a clean separation of concerns and makes it easier to navigate and maintain the codebase.

## Setup and Installation

### 1. Clone the repository:

```bash
git clone <repository_url>
cd <repository_directory>
```

### 2. Installing dependancies:

```bash
go mod tidy
```

### 3. Running the application:

```bash
go run main.go
```

## API Endpoints

- **POST /tags**: Create a new tag.
- **GET /tags/:id**: Get a tag by ID.
- **PUT /tags/:id**: Update a tag by ID.
- **DELETE /tags/:id**: Delete a tag by ID.
- **GET /tags**: Get all tags.

## Contributing

Contributions are welcome! If you find any issues or improvements, please feel free to open an issue or create a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
