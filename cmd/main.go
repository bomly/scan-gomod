package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type Dependency struct {
	Module  string `json:"module"`
	Version string `json:"version"`
	Scope   string `json:"scope"`
	Type    string `json:"type"`
}

type Module struct {
	Name         string       `json:"name"`
	Version      string       `json:"version"`
	Dependencies []Dependency `json:"dependencies"`
}

type ScanResult struct {
	Version           string `json:"version"`
	Module            Module `json:"module"`
	DependencyManager string `json:"dependencyManager"`
}

var scanResult = ScanResult{
	Version:           "1.0.0",
	DependencyManager: "gomod",
	Module: Module{
		Name:    "github.com/kyuff/es",
		Version: "git sha",
		Dependencies: []Dependency{
			{
				Module:  "github.com/gofrs/uuid/v5",
				Version: "v5.3.1",
				Scope:   "code",
				Type:    "gomod",
			},
			{
				Module:  "golang.org/x/sync",
				Version: "v0.11.0",
				Scope:   "code",
				Type:    "gomod",
			},
		},
	},
}

var (
	output = ""
)

func init() {
	flag.StringVar(&output, "o", output, "Output file to write dependencies to")
}

func main() {
	flag.Parse()

	var buf bytes.Buffer
	var enc = json.NewEncoder(&buf)

	enc.SetIndent("", "  ")

	err := enc.Encode(scanResult)
	if err != nil {
		fmt.Printf("Failed to encode dependencies to %s: %v\n", output, err)
		os.Exit(1)
	}

	if output == "" {
		fmt.Println(buf.String())
		os.Exit(0)
	}

	err = os.WriteFile(output, buf.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed to write file to %s: %v\n", output, err)
		os.Exit(1)
	}

	fmt.Printf("Output: %s\n", output)
}
