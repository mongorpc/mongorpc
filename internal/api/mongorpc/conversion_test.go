package mongorpc

import (
	"testing"

	"go.mongodb.org/mongo-driver/v2/bson"

	mongorpcv1 "github.com/mongorpc/mongorpc/gen/mongorpc/v1"
)

func TestBsonToValue_Primitives(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected *mongorpcv1.Value
	}{
		{
			name:  "nil value",
			input: nil,
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_NullValue{},
			},
		},
		{
			name:  "string value",
			input: "hello",
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_StringValue{StringValue: "hello"},
			},
		},
		{
			name:  "int32 value",
			input: int32(42),
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 42},
			},
		},
		{
			name:  "int64 value",
			input: int64(9999999999),
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_Int64Value{Int64Value: 9999999999},
			},
		},
		{
			name:  "float64 value",
			input: float64(3.14),
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_DoubleValue{DoubleValue: 3.14},
			},
		},
		{
			name:  "bool true",
			input: true,
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_BooleanValue{BooleanValue: true},
			},
		},
		{
			name:  "bool false",
			input: false,
			expected: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_BooleanValue{BooleanValue: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := bsonToValue(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Compare by type
			switch expected := tt.expected.ValueType.(type) {
			case *mongorpcv1.Value_NullValue:
				if _, ok := result.ValueType.(*mongorpcv1.Value_NullValue); !ok {
					t.Errorf("expected NullValue, got %T", result.ValueType)
				}
			case *mongorpcv1.Value_StringValue:
				if actual, ok := result.ValueType.(*mongorpcv1.Value_StringValue); !ok {
					t.Errorf("expected StringValue, got %T", result.ValueType)
				} else if actual.StringValue != expected.StringValue {
					t.Errorf("expected %q, got %q", expected.StringValue, actual.StringValue)
				}
			case *mongorpcv1.Value_Int32Value:
				if actual, ok := result.ValueType.(*mongorpcv1.Value_Int32Value); !ok {
					t.Errorf("expected Int32Value, got %T", result.ValueType)
				} else if actual.Int32Value != expected.Int32Value {
					t.Errorf("expected %d, got %d", expected.Int32Value, actual.Int32Value)
				}
			case *mongorpcv1.Value_Int64Value:
				if actual, ok := result.ValueType.(*mongorpcv1.Value_Int64Value); !ok {
					t.Errorf("expected Int64Value, got %T", result.ValueType)
				} else if actual.Int64Value != expected.Int64Value {
					t.Errorf("expected %d, got %d", expected.Int64Value, actual.Int64Value)
				}
			case *mongorpcv1.Value_DoubleValue:
				if actual, ok := result.ValueType.(*mongorpcv1.Value_DoubleValue); !ok {
					t.Errorf("expected DoubleValue, got %T", result.ValueType)
				} else if actual.DoubleValue != expected.DoubleValue {
					t.Errorf("expected %f, got %f", expected.DoubleValue, actual.DoubleValue)
				}
			case *mongorpcv1.Value_BooleanValue:
				if actual, ok := result.ValueType.(*mongorpcv1.Value_BooleanValue); !ok {
					t.Errorf("expected BooleanValue, got %T", result.ValueType)
				} else if actual.BooleanValue != expected.BooleanValue {
					t.Errorf("expected %v, got %v", expected.BooleanValue, actual.BooleanValue)
				}
			}
		})
	}
}

func TestBsonToValue_ObjectID(t *testing.T) {
	oid := bson.NewObjectID()
	result, err := bsonToValue(oid)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual, ok := result.ValueType.(*mongorpcv1.Value_ObjectIdValue)
	if !ok {
		t.Fatalf("expected ObjectIdValue, got %T", result.ValueType)
	}

	if actual.ObjectIdValue.Hex != oid.Hex() {
		t.Errorf("expected %q, got %q", oid.Hex(), actual.ObjectIdValue.Hex)
	}
}

func TestBsonToValue_Array(t *testing.T) {
	input := bson.A{"a", "b", "c"}
	result, err := bsonToValue(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual, ok := result.ValueType.(*mongorpcv1.Value_ArrayValue)
	if !ok {
		t.Fatalf("expected ArrayValue, got %T", result.ValueType)
	}

	if len(actual.ArrayValue.Values) != 3 {
		t.Errorf("expected 3 values, got %d", len(actual.ArrayValue.Values))
	}
}

func TestBsonToValue_Map(t *testing.T) {
	input := bson.M{"key1": "value1", "key2": int32(42)}
	result, err := bsonToValue(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	actual, ok := result.ValueType.(*mongorpcv1.Value_MapValue)
	if !ok {
		t.Fatalf("expected MapValue, got %T", result.ValueType)
	}

	if len(actual.MapValue.Fields) != 2 {
		t.Errorf("expected 2 fields, got %d", len(actual.MapValue.Fields))
	}
}

func TestValueToBSON_Primitives(t *testing.T) {
	tests := []struct {
		name     string
		input    *mongorpcv1.Value
		expected interface{}
	}{
		{
			name:     "nil value",
			input:    nil,
			expected: nil,
		},
		{
			name: "string value",
			input: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_StringValue{StringValue: "hello"},
			},
			expected: "hello",
		},
		{
			name: "int32 value",
			input: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 42},
			},
			expected: int32(42),
		},
		{
			name: "bool value",
			input: &mongorpcv1.Value{
				ValueType: &mongorpcv1.Value_BooleanValue{BooleanValue: true},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := valueToBSON(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func TestBsonToDocument(t *testing.T) {
	input := bson.M{
		"_id":  bson.NewObjectID(),
		"name": "test",
		"age":  int32(30),
	}

	doc, err := bsonToDocument(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(doc.Fields) != 3 {
		t.Errorf("expected 3 fields, got %d", len(doc.Fields))
	}

	if doc.Fields["name"] == nil {
		t.Error("expected 'name' field to exist")
	}
}

func TestDocumentToBSON(t *testing.T) {
	doc := &mongorpcv1.Document{
		Fields: map[string]*mongorpcv1.Value{
			"name": {ValueType: &mongorpcv1.Value_StringValue{StringValue: "test"}},
			"age":  {ValueType: &mongorpcv1.Value_Int32Value{Int32Value: 30}},
		},
	}

	result, err := documentToBSON(doc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("expected 2 elements, got %d", len(result))
	}
}

func TestDocumentToBSON_NilDocument(t *testing.T) {
	result, err := documentToBSON(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("expected empty document, got %d elements", len(result))
	}
}
