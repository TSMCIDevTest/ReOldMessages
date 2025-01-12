# Instant Messenger Revival

This project aims to revive an instant messaging application using Go, integrating the libpurple library for messaging functionalities and the Nina chat API for communication.

## Project Structure

```
instant-messenger-revival
├── src
│   ├── main.go          # Entry point of the application
│   ├── libpurple        # Integration with libpurple library
│   │   └── libpurple.go # Implementation of libpurple functionalities
│   ├── nina             # Interaction with Nina chat API
│   │   └── nina.go      # API request and response management
│   └── types            # Type definitions
│       └── index.go     # Message, User, and ChatSession types
├── go.mod               # Module definition
├── go.sum               # Module dependency checksums
└── README.md            # Project documentation
```

## Setup Instructions

1. Clone the repository:
   ```
   git clone <repository-url>
   cd instant-messenger-revival
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Build the application:
   ```
   go build -o instant-messenger-revival src/main.go
   ```

4. Run the application:
   ```
   ./instant-messenger-revival
   ```

## Usage

- The application initializes the libpurple library and connects to the Nina chat API.
- Users can send and receive messages through the integrated messaging functionalities.

## Contribution Guidelines

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Make your changes and commit them.
4. Push your branch and create a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.