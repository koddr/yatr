package main

import (
	"errors"
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
		return nil, errors.New("parsing error: invalid format of tasks file")
	}

	// Load task set by format's parser.
	switch parserFormat[len(parserFormat)-1] {
	case "json":
		// Load JSON parser.
		if err := k.Load(file.Provider(*p), json.Parser()); err != nil {
			return nil, err
		}
	case "yaml", "yml":
		// Load YAML parser.
		if err := k.Load(file.Provider(*p), yaml.Parser()); err != nil {
			return nil, err
		}
	case "toml":
		// Load TOML parser.
		if err := k.Load(file.Provider(*p), toml.Parser()); err != nil {
			return nil, err
		}
	case "tf":
		// Load HCL (Terraform) parser.
		if err := k.Load(file.Provider(*p), hcl.Parser(true)); err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("parsing error: unknown format of tasks file")
	}

	// Create a new tasks set.
	var tt *Tasks

	// Unmarshal tasks to the given tasks set.
	if err := k.Unmarshal("", &tt); err != nil {
		return nil, err
	}

	return tt, nil
}
