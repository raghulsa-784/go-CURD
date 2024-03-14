# CRUD API with GoLang and PostgreSQL

This repository contains a CRUD (Create, Read, Update, Delete) API implemented in GoLang, working with a PostgreSQL database.

## Features

- **Create:** Add new records to the PostgreSQL database.
- **Read:** Retrieve records from the database based on specified criteria.
- **Update:** Modify existing records in the database.
- **Delete:** Remove records from the database.

## Prerequisites

Before running the API, ensure you have the following installed:

- GoLang: [Installation Guide](https://golang.org/doc/install)
- PostgreSQL: [Installation Guide](https://www.postgresql.org/download/)
- PostgreSQL Database setup (Database name, user, password, etc.)

## Setup

1. Clone the repository:

    ```bash
    git clone https://github.com/Wicked2468/go-CURD.git
    cd go-CURD
    ```

2. Install dependencies:

    You can use Go modules for managing dependencies. Run the following command to install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up PostgreSQL:

    - Create a PostgreSQL database.
    - Update the database connection details in the `.env` file.

4. Build and run the API:

    ```bash
    go build -o main .
    ./main
    ```

## Endpoints

- **r.GET("/posts", controllers.PostIndex)** Retrieve all resources.
- **r.GET("/posts/:id", controllers.SinglePost)** Retrieve a specific resource by ID.
- **r.POST("/posts", controllers.PostCreate)** Create a new resource.
- **r.PUT("/posts", controllers.PostUpdate)** Update an existing resource by ID.
- **r.DELETE("/posts/:id", controllers.PostDelete)** Delete a resource by ID.

## Configuration

- Modify the `.env` file to specify the PostgreSQL database connection details.
- Update the `main.go` file to define additional API routes or modify existing ones as required.

## Dependencies

- GoLang packages are managed using Go modules.
- The API uses the `go get -u gorm.io/gorm` package for PostgreSQL database interactions.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project was created for educational purposes and can be extended or modified according to specific requirements.
