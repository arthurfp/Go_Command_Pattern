# Command Pattern in Go

## Overview
This repository demonstrates the application of the Command design pattern in Go. The project highlights how to encapsulate requests as objects, enabling parameterization, queuing, and undo functionality, showcasing flexibility and best practices in design patterns and unit testing.

## Pattern Description
The Command pattern encapsulates a request as an object, thereby allowing you to parameterize clients with different requests, queue or log requests, and support undoable operations.

### Key Components
- **Command Interface**: Defines a method for executing a command.
- **Concrete Commands**: Implement the Command interface and define bindings between a Receiver and an action.
- **Receiver**: Knows how to perform the operations associated with a command.
- **Invoker**: Stores commands and executes them in a structured manner.

## Project Structure
- **cmd/**: Contains the application entry point (`main.go`), demonstrating the creation and usage of commands.
- **pkg/command/**: Houses the Command pattern implementation.
    - **command.go**: Defines the `Command` interface.
    - **receiver.go**: Implements the `Receiver`.
    - **command_one.go**: Implements `CommandOne` to execute the first action.
    - **command_two.go**: Implements `CommandTwo` to execute the second action.
    - **invoker.go**: Implements the `Invoker` to manage and execute commands.
    - **receiver_test.go**: Unit tests for the `Receiver`.
    - **command_one_test.go**: Unit tests for `CommandOne`.
    - **command_two_test.go**: Unit tests for `CommandTwo`.
    - **invoker_test.go**: Unit tests for the `Invoker`.

## Getting Started

### Prerequisites
Ensure you have Go installed on your system. You can download it from [Go's official site](https://golang.org/dl/).

### Installation
Clone this repository to your local machine:
```bash
git clone git@github.com:arthurfp/Go_Command_Pattern.git
cd Go_Command_Pattern
```

### Running the Application
To run the application, execute:
```bash
go run cmd/main.go
```

### Running the Tests
To execute the tests and verify the functionality:
```bash
go test ./pkg/command
```

### Example Output
When you run the application, you should see the following output:
```yaml
Executing CommandOne:
Receiver: Executing ActionOne
Executing CommandTwo:
Receiver: Executing ActionTwo
Executing all commands via Invoker:
Receiver: Executing ActionOne
Receiver: Executing ActionTwo
```

### Contributing
Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

### Author
Arthur Ferreira - github.com/arthurfp