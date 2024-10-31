//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Lint runs markdownlint on all Markdown files
func Lint() error {
	fmt.Println("Running markdownlint...")
	cmd := exec.Command("markdownlint", "**/*.md")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Test runs the tests
func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Build builds the binary
func Build() error {
	fmt.Println("Building binary...")
	cmd := exec.Command("go", "build", "-o", "url-shortener", "./cmd/url-shortener")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Run runs the binary
func Run() error {
	fmt.Println("Running binary...")
	cmd := exec.Command("./url-shortener")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
