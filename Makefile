# note: call scripts from /scripts

install_tools:
  go install github.com/swaggo/swag/cmd/swag@latest
  go install golang.org/x/tools/cmd/stringer
  go install github.com/vektra/mockery/v2@latest
  go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

generate:
  go generate ./... 
