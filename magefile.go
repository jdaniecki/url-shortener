//go:build mage
// +build mage

package main

import (
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Clean removes the build directory
func Clean() error {
	fmt.Println("Removing build directory...")
	return sh.Rm("build")
}

type Source mg.Namespace

// Lint code and markdown files
func (s Source) Lint() error {
	fmt.Println("Running markdownlint...")
	if err := sh.RunV("markdownlint", "**/*.md"); err != nil {
		return err
	}
	fmt.Println("Running golangci-lint...")
	return sh.RunV("golangci-lint", "run", "./...")
}

// Test runs the tests
func (s Source) Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "./...")
}

// Generate generates Go client and server code from OpenAPI spec
func (s Source) Generate() error {
	fmt.Println("Generating code from OpenAPI spec...")
	if err := sh.RunV("oapi-codegen", "-generate", "std-http, types", "-o", "internal/api/server.gen.go", "-package", "api", "api/openapi.yaml"); err != nil {
		return err
	}
	return nil
}

type Binary mg.Namespace

// Binary builds the binary
func (b Binary) Build() error {
	fmt.Println("Building binary with CGO disabled...")
	env := map[string]string{"CGO_ENABLED": "0"}
	args := []string{"build", "-o", "build/url-shortener", "-tags", "netgo", "-ldflags", "-extldflags '-static'", "./cmd/url-shortener"}
	return sh.RunWith(env, "go", args...)
}

// Run runs the binary
func (b Binary) Run() error {
	mg.Deps(b.Build)
	fmt.Println("Running binary...")
	return sh.RunV("./build/url-shortener")
}

type Docker mg.Namespace

// Build builds the Docker image
func (d Docker) Build() error {
	fmt.Println("Building Docker image...")
	return sh.RunV("docker", "build", "-t", "url-shortener", ".")
}

// Run runs the Docker container
func (d Docker) Run() error {
	fmt.Println("Running Docker container...")
	return sh.RunV("docker", "run", "-p", "8080:8080", "url-shortener")
}
