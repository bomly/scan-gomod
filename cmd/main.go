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

type ScanResult struct {
	Version           string       `json:"version"`
	Module            string       `json:"module"`
	ModuleVersion     string       `json:"moduleVersion"`
	DependencyManager string       `json:"dependencyManager"`
	Dependencies      []Dependency `json:"dependencies"`
}

var scanResult = ScanResult{
	Version:           "1.0.0",
	Module:            "github.com/kyuff/es",
	ModuleVersion:     "git sha",
	DependencyManager: "gomod",
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
}

var (
	outputFile = "scan-result.gomod.json"
)

func init() {
	flag.StringVar(&outputFile, "o", outputFile, "Output file to write dependencies to")
}

func main() {
	flag.Parse()

	var buf bytes.Buffer
	var enc = json.NewEncoder(&buf)

	enc.SetIndent("", "  ")

	err := enc.Encode(scanResult)
	if err != nil {
		fmt.Printf("Failed to encode dependencies to %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFile, buf.Bytes(), 0644)
	if err != nil {
		fmt.Printf("Failed to write file to %s: %v\n", outputFile, err)
		os.Exit(1)
	}

	fmt.Printf("Output: %s\n", outputFile)
}
