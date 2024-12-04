//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"strings"

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
	if err := sh.RunV("golangci-lint", "run", "./..."); err != nil {
		return err
	}

	fmt.Println("Running hadolint...")
	if err := sh.RunV("hadolint", "Dockerfile"); err != nil {
		return err
	}
	return nil
}

// Test runs the tests
func (s Source) Test() error {
	fmt.Println("Running tests...")
	return sh.RunV("go", "test", "./...")
}

// Generate generates Go client and server code from OpenAPI spec
func (s Source) Generate() error {
	fmt.Println("Generating code from OpenAPI spec...")
	if err := sh.RunV("oapi-codegen", "-generate", "std-http,strict-server,types, spec", "-o", "internal/api/server.gen.go", "-package", "api", "api/openapi.yaml"); err != nil {
		return err
	}
	return nil
}

type Binary mg.Namespace

// Binary builds the binary
func (b Binary) Build() error {
	version, err := readVersion()
	if err != nil {
		return err
	}
	fmt.Println("Building binary with CGO disabled...")
	env := map[string]string{"CGO_ENABLED": "0"}
	ldflags := fmt.Sprintf("-X main.version=%s -extldflags '-static'", version)
	args := []string{"build", "-o", "build/url-shortener", "-tags", "netgo", "-ldflags", ldflags, "./cmd/url-shortener"}
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
	version, err := readVersion()
	if err != nil {
		return err
	}
	fmt.Println("Building Docker image...")
	return sh.RunV("docker", "build", "-t", "url-shortener:"+version, ".")
}

// Run runs the Docker container
func (d Docker) Run() error {
	version, err := readVersion()
	if err != nil {
		return err
	}
	fmt.Println("Running Docker container...")
	return sh.RunV("docker", "run", "-p", "8080:8080", "url-shortener:"+version)
}

// Push pushes the Docker image to the registry
func (d Docker) Push() error {
	version, err := readVersion()
	if err != nil {
		return err
	}
	fmt.Println("Pushing Docker image...")
	if err := sh.RunV("docker", "tag", "url-shortener:"+version, "jozefdaniecki/url-shortener:"+version); err != nil {
		return err
	}
	if !strings.Contains(version, "dev") {
		if err := sh.RunV("docker", "tag", "url-shortener:"+version, "jozefdaniecki/url-shortener:latest"); err != nil {
			return err
		}
		if err := sh.RunV("docker", "push", "jozefdaniecki/url-shortener:latest"); err != nil {
			return err
		}
	}
	return sh.RunV("docker", "push", "jozefdaniecki/url-shortener:"+version)
}

// Scan scans the Docker image for vulnerabilities
func (d Docker) Scan() error {
	version, err := readVersion()
	if err != nil {
		return err
	}
	fmt.Println("Scanning Docker image...")
	return sh.RunV("trivy", "image", "--image-config-scanners", "misconfig,secret", "url-shortener:"+version)
}

// readVersion reads the version from the VERSION file and appends the current commit hash if the version contain "dev"
func readVersion() (string, error) {
	data, err := ioutil.ReadFile("VERSION")
	if err != nil {
		return "", err
	}
	version := strings.TrimSpace(string(data))

	if strings.Contains(version, "dev") {
		commitHash, err := sh.Output("git", "rev-parse", "--short", "HEAD")
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s-%s", version, commitHash), nil
	}

	return version, nil
}
