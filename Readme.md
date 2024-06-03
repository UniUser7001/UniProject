
# go_api README

Welcome to the `go_api` project! This README will guide you through the essential commands to set up, run, and test the API.

## Prerequisites

Ensure you have the following installed on your system:

- Go (version 1.16 or higher)

## Setup

1. **Initialize the Go module**

   Run the following command to initialize the Go module for your project:

   ```sh
   go mod init go_api
   ```

2. **Tidy up the dependencies**

   Use the following command to add any missing module dependencies or remove unused ones:

   ```sh
   go mod tidy
   ```

## Running the Application

To start the application, use the following command:

```sh
go run main.go
```

## Testing

To run the tests for the handlers, execute the following command:

```sh
go test ./handlers -v
```

The `-v` flag is used to provide verbose output, giving more detailed information about the test results.

## Directory Structure

- `main.go`: The entry point of the application.
- `handlers/`: Contains the HTTP handlers for the API.
- `go.mod`: The Go module file, which includes the list of dependencies.

## Additional Resources

For more information on Go modules and managing dependencies, refer to the official Go documentation: [Go Modules](https://golang.org/ref/mod).

---

Feel free to reach out if you have any questions or need further assistance. Happy coding!
