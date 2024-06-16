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
.
├── main.go          # Main entry point of the application
├── repository       # Package containing repository implementations
│   ├── repository.go    # Interface definition for repositories
│   └── tags_repository.go    # Implementation of TagsRepository using GORM
├── data
│   └── request      # Request data structures
│   └── response      # Response data structures
├── model            # Data models used in the application
│   └── tags.go      # Tags data model
├── helper           # Helper functions and utilities
│   └── error.go     # Error handling utilities
└── README.md        # Project documentation (you are here)
```
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
