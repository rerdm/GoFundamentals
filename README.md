# GoFundamentals

This project demonstrates the fundamentals of Go programming language.

## Table of Contents

### Getting Started
- [Setting Up a Go Project in VS Code](#setting-up-a-go-project-in-vs-code)
    - [Install Go](#install-go)
    - [Install Visual Studio Code](#install-visual-studio-code)
    - [Install Go Extension for VS Code](#install-go-extension-for-vs-code)
    - [Create a New Go Project](#create-a-new-go-project)
        - [State-of-the-Art Folder Structure](#state-of-the-art-folder-structure)
    - [Run Your Go Project](#run-your-go-project)
    - [Debugging Your Go Project](#debugging-your-go-project)
    - [Go Modules](#go-modules)
        - [Initializing a Go Module](#initializing-a-go-module)
        - [Importing Packages](#importing-packages)
            - [Internal Packages](#internal-packages)
            - [External Modules](#external-modules)
        - [Managing Dependencies](#managing-dependencies)
    - [Running and Debugging with Configuration](#running-and-debugging-with-configuration)
    - [Setting Environment Variables](#setting-environment-variables)

- Core Concepts
    - [Functions](#functions)
    - [Variables](#variables)
    - [Control Structures](#control-structures)
    - [Packages and Modules](#packages-and-modules)
    - [Data Types](#data-types)
    - [Pointers](#pointers)
    - [Interfaces](#interfaces)
    - [Structs](#structs)
    - [Slices and Arrays](#slices-and-arrays)
    - [Maps](#maps)
    - [Constants](#constants)
    - [Loops](#loops)
    - [Type Conversions](#type-conversions)
    - [Enums](#enums)

- Advanced Topics(TODO)
    - [Error Handling](#error-handling)
    - [Testing](#testing)
    - [Concurrency](#concurrency)
    - [Advanced Topics](#advanced-topics)
    - [Reflection](#reflection)
    - [Generics](#generics)

## Setting Up a Go Project in VS Code

To generate a Go project in Visual Studio Code, follow these steps:

### Install Go

1. Download and install Go from the [official website](https://golang.org/dl/).
2. Follow the installation instructions for your operating system.

### Install Visual Studio Code

1. Download and install Visual Studio Code from the [official website](https://code.visualstudio.com/).
2. Follow the installation instructions for your operating system.

### Install Go Extension for VS Code

1. Open Visual Studio Code.
2. Go to the Extensions view by clicking on the Extensions icon in the Activity Bar on the side of the window.
3. Search for "Go" and install the official Go extension by the Go team at Google.

### Create a New Go Project

1. Open a terminal in Visual Studio Code.
2. Navigate to the directory where you want to create your project.
3. Run the following command to create a new module:
    ```sh
    go mod init <module-name>
    ```
    Replace `<module-name>` with the name of your module. The module name is typically the repository path (e.g., `github.com/username/project`). If you don't have a repository, you can use the project name (e.g., `GoFundamentals`).

4. Create a new file named `main.go` and add the following sample code:
    ```go
    package main

    import "fmt"

    func main() {
         fmt.Println("Hello, World!")
    }
    ```

#### State-of-the-Art Folder Structure

A well-organized folder structure is crucial for maintaining a scalable and manageable Go project. Here is an example of a state-of-the-art folder structure:

```
GoFundamentals/
├── bin/                # Compiled binaries
├── cmd/                # Main applications of the project
│   └── app/
│       └── main.go     # Entry point for the application
├── pkg/                # Library code that's ok to use by external applications
├── internal/           # Private application and library code
│   └── app/
│       └── app.go
├── api/                # OpenAPI/Swagger specs, JSON schema files, protocol definitions
├── web/                # Web application specific components
│   ├── static/         # Static files
│   └── templates/      # HTML templates
├── configs/            # Configuration files
├── scripts/            # Scripts to perform various build, install, analysis, etc.
├── test/               # Additional external test apps and test data
├── docs/               # Documentation files
├── build/              # Packaging and Continuous Integration
├── deployments/        # Deployment configurations and templates
├── examples/           # Examples for your applications and/or public libraries
├── tools/              # Supporting tools for this project
├── vendor/             # Application dependencies (managed by `go mod vendor`)
├── go.mod              # Go module file
├── go.sum              # Go dependencies checksum file
└── README.md           # Project README file
```

### Run Your Go Project

1. Open the terminal in Visual Studio Code.
2. Navigate to the directory containing your `main.go` file.
3. Run the following command to execute your program:
    ```sh
    go run main.go
    ```

### Debugging Your Go Project

1. Set breakpoints in your code by clicking in the gutter next to the line numbers.
2. Open the Run and Debug view by clicking on the Run icon in the Activity Bar.
3. Click on "Run and Debug" and select "Go" from the dropdown menu.
4. Start debugging by clicking on the green play button.

### Go Modules

Go modules are the standard way to manage dependencies in Go projects. They allow you to import packages, both internal and external, and ensure that your project has the correct versions of dependencies.

#### Initializing a Go Module

1. **Open your project in VS Code**.
2. **Open the terminal** in VS Code (`Ctrl + `).
3. **Navigate to the root directory** of your project.
4. **Initialize the module** by running:
   ```sh
   go mod init <module-name>
   ```
   Replace `<module-name>` with the name of your module. The module name is typically the repository path (e.g., `github.com/username/project`). If you don't have a repository, you can use the project name (e.g., `GoFundamentals`).

#### Importing Packages

##### Internal Packages

To import packages from other folders or subfolders within your project:

1. **Ensure the package is correctly defined** in the source file. For example, in `src/core_concepts/constants/constants.go`:
   ```go
   package constants
   ```

2. **Import the package** in your main file or other files:
   ```go
   import (
       "GoFundamentals/src/core_concepts/constants"
   )
   ```

3. **Use the package** in your code:
   ```go
   func main() {
       constants.Execute()
   }
   ```

4. **Update the `go.mod` file** to include the module path:
   ```go
   module GoFundamentals

   go 1.16

   require (
       // Add any required dependencies here
   )
   ```

##### External Modules

To import external modules:

1. **Add the module to your project** by running:
   ```sh
   go get <module-path>
   ```
   Replace `<module-path>` with the path of the external module (e.g., `github.com/gin-gonic/gin`).

2. **Import the module** in your code:
   ```go
   import (
       "github.com/gin-gonic/gin"
   )
   ```

3. **Use the module** in your code:
   ```go
   func main() {
       r := gin.Default()
       r.GET("/", func(c *gin.Context) {
           c.String(200, "Hello, World!")
       })
       r.Run()
   }
   ```

4. **Update the `go.mod` file** to include the external module:
   ```go
   module GoFundamentals

   go 1.16

   require (
       github.com/gin-gonic/gin v1.7.4 // Replace with the actual version
   )
   ```

#### Managing Dependencies

1. **Run `go mod tidy`** to clean up your `go.mod` and `go.sum` files:
   ```sh
   go mod tidy
   ```

2. **Run `go mod vendor`** to create a `vendor` directory with all dependencies (optional):
   ```sh
   go mod vendor
   ```

By following these steps, you can effectively manage internal and external packages in your Go project using Go modules. This ensures that your project dependencies are correctly resolved and maintained.

### Running and Debugging with Configuration

To run and debug your Go project using a configuration file, you can create a `launch.json` file in the `.vscode` directory. This allows you to run and debug your project directly from the Run and Debug view in Visual Studio Code.

#### Creating the `launch.json` File

1. **Create a `.vscode` directory** in the root of your project if it doesn't already exist.
2. **Create a `launch.json` file** inside the `.vscode` directory.
3. **Add the following configuration** to the `launch.json` file:
    ```json
    {
        "version": "0.2.0",
        "configurations": [
            {
                "name": "Launch Go Program",
                "type": "go",
                "request": "launch",
                "mode": "auto",
                "program": "${workspaceFolder}/cmd/app/main.go",
                "env": {
                    "ENV_VAR1": "value1",
                    "ENV_VAR2": "value2"
                },
                "args": []
            }
        ]
    }
    ```
    - `name`: The name of the configuration.
    - `type`: The type of the configuration, which is `go` for Go programs.
    - `request`: The request type, which is `launch` to launch the program.
    - `mode`: The mode to run the program, which can be `auto`, `debug`, or `test`.
    - `program`: The path to the main Go file to run.
    - `env`: Environment variables to set for the program.
    - `args`: Command-line arguments to pass to the program.

#### Using the Configuration

1. **Open the Run and Debug view** by clicking on the Run icon in the Activity Bar.
2. **Select the configuration** you created from the dropdown menu.
3. **Click on the green play button** to run and debug your Go program.

By using a `launch.json` configuration file, you can easily run and debug your Go project directly from Visual Studio Code.

#### Installing the `dlv` Command

If you encounter the error "The `dlv` command is not available," you need to install the Delve debugger. Delve is a debugger for the Go programming language.

1. Open a terminal.
2. Run the following command to install Delve:
    ```sh
    go install -v github.com/go-delve/delve/cmd/dlv@latest
    ```
3. Ensure that the `dlv` command is available in your `PATH`.

By installing the Delve debugger, you can use the `launch.json` configuration to run and debug your Go project in Visual Studio Code.

### Setting Environment Variables

To set environment variables when running your Go project from the terminal, you can use the `export` command (on Unix-like systems) or `set` command (on Windows).

#### Unix-like Systems (Linux, macOS)

1. Open a terminal.
2. Set the environment variables:
    ```sh
    export ENV_VAR1=value1
    export ENV_VAR2=value2
    ```
3. Run your Go program:
    ```sh
    go run cmd/app/main.go
    ```

#### Windows

1. Open a Command Prompt or PowerShell.
2. Set the environment variables:
    ```cmd
    set ENV_VAR1=value1
    set ENV_VAR2=value2
    ```
3. Run your Go program:
    ```cmd
    go run cmd/app/main.go
    ```

By setting the environment variables in the terminal, you ensure that they are available to your Go program when it runs.

## Core Concepts

### Functions

See the [functions.go](src/core_concepts/functions/functions.go) file for examples of functions, variadic arguments, named return values, and handling pointers in functions.

### Variables

See the [variables.go](src/core_concepts/variables/variables.go) file for examples of variable declarations, types, and scope.

### Control Structures

See the [control_structures.go](src/core_concepts/control_structures/control_structures.go) file for examples of if-else statements, switch statements, and loops.

### Packages and Modules

See the [packages_and_modules.go](src/core_concepts/packages_and_modules/packages_and_modules.go) file for examples of using packages and modules, including importing external modules.

### Data Types

See the [data_types.go](src/core_concepts/data_types/data_types.go) file for examples of basic data types, composite data types, and type conversions.

### Pointers

See the [pointers.go](src/core_concepts/pointers/pointers.go) file for examples of pointers, pointer arithmetic, and pointer dereferencing.

### Interfaces

See the [interfaces.go](src/core_concepts/interfaces/interfaces.go) file for examples of defining and implementing interfaces.

### Structs

See the [structs.go](src/core_concepts/structs/structs.go) file for examples of defining structs, using struct tags, and accessing struct fields.

### Slices and Arrays

See the [slices_and_arrays.go](src/core_concepts/slices_and_arrays/slices_and_arrays.go) file for examples of slices, arrays, and their operations.

### Maps

See the [maps.go](src/core_concepts/maps/maps.go) file for examples of defining and using maps.

### Constants

See the [constants.go](src/core_concepts/constants/constants.go) file for examples of defining and using constants.

### Loops

See the [loops.go](src/core_concepts/loops/loops.go) file for examples of different types of loops.

### Type Conversions

See the [type_conversions.go](src/core_concepts/type_conversions/type_conversions.go) file for examples of type conversions.

### Enums

See the [enums.go](src/core_concepts/enums/enums.go) file for examples of defining and using enums in Go.
