package main

import (
	"fmt"
	"strings"

	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

// newParser provides the parsing process.
func newParser(p *string, k *koanf.Koanf) (*Tasks, error) {
	// Parse the format of the given file p.
	parserFormat := strings.Split(strings.ToLower(*p), ".")

	// Check, if the file name is too short.
	if len(parserFormat) < 2 {
		return nil, fmt.Errorf("parsing error: invalid format of tasks file, see: %s", WikiPageURL)
	}

	// Load task set by format's parser.
	switch parserFormat[len(parserFormat)-1] {
	case "json":
		// Load JSON parser.
		if err := k.Load(file.Provider(*p), json.Parser()); err != nil {
			return nil, fmt.Errorf("error: not valid structure of the JSON file, see: %s", WikiPageURL)
		}
	case "yaml", "yml":
		// Load YAML parser.
		if err := k.Load(file.Provider(*p), yaml.Parser()); err != nil {
			return nil, fmt.Errorf("error: not valid structure of the YAML file, see: %s", WikiPageURL)
		}
	case "toml":
		// Load TOML parser.
		if err := k.Load(file.Provider(*p), toml.Parser()); err != nil {
			return nil, fmt.Errorf("error: not valid structure of the TOML file, see: %s", WikiPageURL)
		}
	case "tf":
		// Load HCL (Terraform) parser.
		if err := k.Load(file.Provider(*p), hcl.Parser(true)); err != nil {
			return nil, fmt.Errorf("error: not valid structure of the HCL file, see: %s", WikiPageURL)
		}
	default:
		return nil, fmt.Errorf("error: unknown format of tasks file, see: %s", WikiPageURL)
	}

	// Create a new tasks set.
	var tt *Tasks

	// Unmarshal tasks to the given tasks set.
	if err := k.Unmarshal("", &tt); err != nil {
		return nil, err
	}

	return tt, nil
}
