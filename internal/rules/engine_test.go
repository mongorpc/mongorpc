package rules

import (
	"context"
	"testing"
)

func TestEngine_DefaultDeny(t *testing.T) {
	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("failed to create engine: %v", err)
	}

	req := &Request{
		Database:   "test",
		Collection: "users",
		Operation:  OpRead,
	}

	allowed, err := engine.Evaluate(context.Background(), req)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}

	if allowed {
		t.Error("expected default deny, got allow")
	}
}

func TestEngine_DefaultAllow(t *testing.T) {
	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("failed to create engine: %v", err)
	}

	engine.SetDefaultAllow(OpRead, true)

	req := &Request{
		Database:   "test",
		Collection: "users",
		Operation:  OpRead,
	}

	allowed, err := engine.Evaluate(context.Background(), req)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}

	if !allowed {
		t.Error("expected allow after SetDefaultAllow(true)")
	}
}

func TestEngine_AddRule(t *testing.T) {
	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("failed to create engine: %v", err)
	}

	// Add a rule that always allows
	err = engine.AddRule("test/users", "read", "true")
	if err != nil {
		t.Fatalf("failed to add rule: %v", err)
	}

	req := &Request{
		Database:   "test",
		Collection: "users",
		Operation:  OpRead,
	}

	allowed, err := engine.Evaluate(context.Background(), req)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}

	if !allowed {
		t.Error("expected allow with 'true' rule")
	}
}

func TestEngine_AddRule_InvalidExpression(t *testing.T) {
	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("failed to create engine: %v", err)
	}

	// Add an invalid rule
	err = engine.AddRule("test/users", "read", "invalid syntax !!!")
	if err == nil {
		t.Error("expected error for invalid expression")
	}
}

func TestEngine_AuthBasedRule(t *testing.T) {
	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("failed to create engine: %v", err)
	}

	// Add a rule that checks admin status
	err = engine.AddRule("test/admin", "read", "auth.isAdmin == true")
	if err != nil {
		t.Fatalf("failed to add rule: %v", err)
	}

	// Test with non-admin
	req := &Request{
		Database:   "test",
		Collection: "admin",
		Operation:  OpRead,
		IsAdmin:    false,
	}

	allowed, err := engine.Evaluate(context.Background(), req)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}

	if allowed {
		t.Error("expected deny for non-admin")
	}

	// Test with admin
	req.IsAdmin = true
	allowed, err = engine.Evaluate(context.Background(), req)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}

	if !allowed {
		t.Error("expected allow for admin")
	}
}

func TestEngine_LoadConfig(t *testing.T) {
	engine, err := NewEngine()
	if err != nil {
		t.Fatalf("failed to create engine: %v", err)
	}

	config := &Config{
		Rules: []RuleConfig{
			{Path: "test/public", Operation: "read", Allow: "true"},
			{Path: "test/private", Operation: "read", Allow: "auth.isAdmin"},
		},
	}

	err = engine.LoadConfig(config)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Test public rule
	req := &Request{
		Database:   "test",
		Collection: "public",
		Operation:  OpRead,
	}

	allowed, err := engine.Evaluate(context.Background(), req)
	if err != nil {
		t.Fatalf("evaluation error: %v", err)
	}

	if !allowed {
		t.Error("expected allow for public collection")
	}
}

func TestExtractOperation(t *testing.T) {
	tests := []struct {
		method   string
		expected Operation
	}{
		{"/mongorpc.v1.MongoRPC/GetDocument", OpRead},
		{"/mongorpc.v1.MongoRPC/ListDocuments", OpRead},
		{"/mongorpc.v1.MongoRPC/CreateDocument", OpCreate},
		{"/mongorpc.v1.MongoRPC/UpdateDocument", OpUpdate},
		{"/mongorpc.v1.MongoRPC/DeleteDocument", OpDelete},
		{"/mongorpc.v1.MongoRPC/RunQuery", OpRead},
		{"/mongorpc.v1.MongoRPC/Aggregate", OpRead},
		{"/mongorpc.v1.MongoRPC/Watch", OpRead},
		{"/unknown/Method", ""},
	}

	for _, tt := range tests {
		t.Run(tt.method, func(t *testing.T) {
			result := extractOperation(tt.method)
			if result != tt.expected {
				t.Errorf("extractOperation(%q) = %q, expected %q", tt.method, result, tt.expected)
			}
		})
	}
}
