//go:build tools
// +build tools

package main

import (
	_ "github.com/vektra/mockery"
	_ "golang.org/x/tools/cmd/stringer"
)
