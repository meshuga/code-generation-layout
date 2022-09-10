# Code generation playground
The goal of the project is to have every common usage of code generation in a typical Golang project.

The project uses [Standard Go Project Layout](https://github.com/golang-standards/project-layout) as a basis

# Code generation samples

* OpenAPI generation with [swag](https://github.com/swaggo/swag), see `main.go`
* Stringer with [stringer](https://pkg.go.dev/golang.org/x/tools@v0.1.12/cmd/stringer) `internal/app/example/model/model.go`
* Interface mocking for tests with [Mockery](https://github.com/vektra/mockery), see `main_test.go`
* SQL Query code generation with [sqlc](https://sqlc.dev/), see ``

# Installing tools

```shell
make install_tools
```

# Generating code

Run the following code to invoke code generation:
```shell
make generate
```
