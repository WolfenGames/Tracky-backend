# Tracky Backend

This repository contains a backend server built with Go, following the MVC (Model-View-Controller) architectural pattern. The server uses SQLite as the database and includes routes for logging and managing activities.

## Features

- MVC architecture for code organization
- SQLite database for lightweight and portable storage
- RESTful API endpoints for managing activities
- Unit tests for controllers and models

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.18 or later)
- SQLite3

## Project Structure

```
├── config        # Database configuration and setup
├── controller    # Handles HTTP requests and responses
├── model         # Database models and queries
├── routes        # Route definitions and middleware
├── main.go       # Entry point of the application
├── go.mod        # Module dependencies
└── go.sum        # Module checksums
```

## Installation

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/wolfengames/tracky-backend.git
   cd tracky-backend
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set Up the Database**:
   The application uses SQLite. If the database file does not exist, it will be created automatically when the server starts. Ensure the `data.db` file has the necessary permissions if you run into issues.

## Running the Server

Start the server with the following command:

```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

## API Endpoints

### Trackable Routes

#### **Log Trackable**
- **Endpoint**: `POST /trackable`
- **Description**: Logs a new Trackable.
- **Request Body**:
  ```json
  {
    "tool": "VSCode",
    "metadata": "path/to/file.txt",
    "start_time": "2024-12-11T10:00:00Z",
    "end_time": "2024-12-11T10:30:00Z"
  }
  ```
- **Response**:
  ```json
  {
    "message": "Trackable logged",
    "id": 42
  }
  ```

## Testing

Run unit tests for controllers and models using the following command:

```bash
go test ./...
```

### Example Test Case

The `LogTrackable` route has test cases in `controllers/Trackable_test.go`. Example:

```go
func TestLogTrackable(t *testing.T) {
    setupTestDB()
    // Your test implementation...
}
```

## Environment Configuration

The application currently uses SQLite. In the future, environment variables can be used to configure different databases or server settings.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

