// Package rules provides a CEL-based security rules engine for MongoRPC.
package rules

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

// Operation represents the type of database operation.
type Operation string

const (
	OpRead   Operation = "read"
	OpWrite  Operation = "write"
	OpCreate Operation = "create"
	OpUpdate Operation = "update"
	OpDelete Operation = "delete"
	OpList   Operation = "list"
)

// Request represents a request context for rule evaluation.
type Request struct {
	// Auth information
	UserID    string
	UserEmail string
	IsAdmin   bool
	Claims    map[string]interface{}

	// Resource information
	Database   string
	Collection string
	DocumentID string

	// Operation
	Operation Operation

	// Document data (for write operations)
	IncomingData map[string]interface{}
	ExistingData map[string]interface{}
}

// Rule represents a security rule.
type Rule struct {
	// Path pattern (e.g., "databases/{database}/collections/{collection}")
	Path string
	// CEL expression for the rule
	Allow string
}

// Engine is the rules evaluation engine.
type Engine struct {
	env      *cel.Env
	rules    map[string]*compiledRule
	defaults map[Operation]bool
}

type compiledRule struct {
	program cel.Program
}

// NewEngine creates a new rules engine.
func NewEngine() (*Engine, error) {
	// Create CEL environment with custom declarations
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar("request", decls.NewMapType(decls.String, decls.Dyn)),
			decls.NewVar("auth", decls.NewMapType(decls.String, decls.Dyn)),
			decls.NewVar("resource", decls.NewMapType(decls.String, decls.Dyn)),
			decls.NewVar("incoming", decls.NewMapType(decls.String, decls.Dyn)),
			decls.NewVar("existing", decls.NewMapType(decls.String, decls.Dyn)),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create CEL environment: %w", err)
	}

	return &Engine{
		env:   env,
		rules: make(map[string]*compiledRule),
		defaults: map[Operation]bool{
			OpRead:   false,
			OpWrite:  false,
			OpCreate: false,
			OpUpdate: false,
			OpDelete: false,
			OpList:   false,
		},
	}, nil
}

// AddRule adds a rule to the engine.
func (e *Engine) AddRule(path, operation, expression string) error {
	// Parse and compile the CEL expression
	ast, issues := e.env.Compile(expression)
	if issues != nil && issues.Err() != nil {
		return fmt.Errorf("failed to compile rule: %w", issues.Err())
	}

	program, err := e.env.Program(ast)
	if err != nil {
		return fmt.Errorf("failed to create program: %w", err)
	}

	key := fmt.Sprintf("%s:%s", path, operation)
	e.rules[key] = &compiledRule{program: program}

	slog.Debug("Added rule", "path", path, "operation", operation)
	return nil
}

// SetDefaultAllow sets the default allow/deny for an operation.
func (e *Engine) SetDefaultAllow(op Operation, allow bool) {
	e.defaults[op] = allow
}

// Evaluate evaluates rules for a request.
func (e *Engine) Evaluate(ctx context.Context, req *Request) (bool, error) {
	// Build the key for rule lookup
	key := fmt.Sprintf("%s/%s:%s", req.Database, req.Collection, string(req.Operation))

	// Check for specific rule
	rule, exists := e.rules[key]
	if !exists {
		// Try wildcard rules
		key = fmt.Sprintf("*/*:%s", string(req.Operation))
		rule, exists = e.rules[key]
	}

	if !exists {
		// Use default
		return e.defaults[req.Operation], nil
	}

	// Build evaluation context
	evalCtx := map[string]interface{}{
		"request": map[string]interface{}{
			"database":   req.Database,
			"collection": req.Collection,
			"documentId": req.DocumentID,
			"operation":  string(req.Operation),
		},
		"auth": map[string]interface{}{
			"userId":  req.UserID,
			"email":   req.UserEmail,
			"isAdmin": req.IsAdmin,
			"claims":  req.Claims,
		},
		"resource": map[string]interface{}{
			"database":   req.Database,
			"collection": req.Collection,
		},
		"incoming": req.IncomingData,
		"existing": req.ExistingData,
	}

	// Evaluate the rule
	result, _, err := rule.program.Eval(evalCtx)
	if err != nil {
		slog.Error("Rule evaluation error", "key", key, "error", err)
		return false, err
	}

	// Convert result to bool
	allowed, ok := result.Value().(bool)
	if !ok {
		return false, fmt.Errorf("rule did not return boolean: %T", result.Value())
	}

	slog.Debug("Rule evaluated", "key", key, "allowed", allowed)
	return allowed, nil
}

// Config represents rules configuration.
type Config struct {
	Rules []RuleConfig `json:"rules" yaml:"rules"`
}

// RuleConfig represents a single rule configuration.
type RuleConfig struct {
	Path      string `json:"path" yaml:"path"`
	Operation string `json:"operation" yaml:"operation"`
	Allow     string `json:"allow" yaml:"allow"`
}

// LoadConfig loads rules from a configuration.
func (e *Engine) LoadConfig(config *Config) error {
	for _, rc := range config.Rules {
		if err := e.AddRule(rc.Path, rc.Operation, rc.Allow); err != nil {
			return fmt.Errorf("failed to add rule %s:%s: %w", rc.Path, rc.Operation, err)
		}
	}
	return nil
}
