//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Lint runs markdownlint on all Markdown files
func Lint() error {
	fmt.Println("Running markdownlint...")
	return sh.RunV("markdownlint", "**/*.md")
}

// Test runs the tests
func Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "./...")
}

// Build builds the binary
func Build() error {
	fmt.Println("Building binary...")
	return sh.RunV("go", "build", "-o", "build/url-shortener", "./cmd/url-shortener")
}

// Run runs the binary
func Run() error {
	mg.Deps(Build)
	fmt.Println("Running binary...")
	return sh.RunV("./build/url-shortener")
}

// Clean removes the build directory
func Clean() error {
	fmt.Println("Removing build directory...")
	return sh.Rm("build")
}

// Generate generates Go client and server code from OpenAPI spec
func Generate() error {
	fmt.Println("Generating code from OpenAPI spec...")
	if err := sh.RunV("oapi-codegen", "-generate", "std-http, types", "-o", "internal/api/server.gen.go", "-package", "api", "api/openapi.yaml"); err != nil {
		return err
	}
	return nil
}
