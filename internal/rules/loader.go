package rules

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadFromFile loads rules from a YAML file.
func LoadFromFile(path string) (*Engine, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read rules file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse rules file: %w", err)
	}

	engine, err := NewEngine()
	if err != nil {
		return nil, err
	}

	// If rules are loaded, deny access by default for all operations
	engine.SetDefaultAllow(OpRead, false)
	engine.SetDefaultAllow(OpWrite, false)
	engine.SetDefaultAllow(OpCreate, false)
	engine.SetDefaultAllow(OpUpdate, false)
	engine.SetDefaultAllow(OpDelete, false)
	engine.SetDefaultAllow(OpList, false)

	if err := engine.LoadConfig(&config); err != nil {
		return nil, err
	}

	return engine, nil
}
