# Code generation playground
The goal of the project is to have every common usage of code generation in a typical Golang project.

The project uses Standard Go Project Layout as a basis

# Code generation samples

* OpenAPI generation with [swag](https://github.com/swaggo/swag), see ``
* Stringer with [stringer](https://pkg.go.dev/golang.org/x/tools@v0.1.12/cmd/stringer)
* Interface mocking with [Mockery]

# Installing tools

```shell
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/stringer
go install github.com/vektra/mockery/v2@latest
```

# Generating code

Run the following code to invoke code generation:
```shell
go generate ./...
```
