package main

import (
	"strings"
	"testing"
)

func TestWriteEnumType(t *testing.T) {
	tests := []struct {
		name     string
		typeInfo Type
		expected string
	}{
		{
			name: "single value enum",
			typeInfo: Type{
				Name: "Color",
				EnumValues: []Value{
					{Name: "RED"},
				},
			},
			expected: "enum Color {\n  RED\n}\n\n",
		},
		{
			name: "multiple values enum",
			typeInfo: Type{
				Name: "Color",
				EnumValues: []Value{
					{Name: "RED"},
					{Name: "GREEN"},
					{Name: "BLUE"},
				},
			},
			expected: "enum Color {\n  RED\n  GREEN\n  BLUE\n}\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			writeEnumType(&sb, tt.typeInfo)
			if got := sb.String(); got != tt.expected {
				t.Errorf("writeEnumType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWriteScalarType(t *testing.T) {
	tests := []struct {
		name     string
		typeInfo Type
		expected string
	}{
		{
			name: "simple scalar",
			typeInfo: Type{
				Name: "DateTime",
			},
			expected: "scalar DateTime\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			writeScalarType(&sb, tt.typeInfo)
			if got := sb.String(); got != tt.expected {
				t.Errorf("writeScalarType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWriteObjectType(t *testing.T) {
	tests := []struct {
		name          string
		typeInfo      Type
		implementsMap map[string][]string
		expected      string
	}{
		{
			name: "object type with fields and interface",
			typeInfo: Type{
				Name: "Car",
				Fields: []Field{
					{Name: "name", Type: Type{Name: "String"}},
					{Name: "mileage", Type: Type{Name: "Int"}},
				},
			},
			implementsMap: map[string][]string{
				"Car": {"Vehicle"},
			},
			expected: "type Car implements Vehicle {\n  name: String\n  mileage: Int\n}\n\n",
		},
		{
			name: "object type without interface",
			typeInfo: Type{
				Name: "Car",
				Fields: []Field{
					{Name: "name", Type: Type{Name: "String"}},
				},
			},
			implementsMap: map[string][]string{},
			expected:      "type Car {\n  name: String\n}\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			writeObjectType(&sb, tt.typeInfo, tt.implementsMap)
			if got := sb.String(); got != tt.expected {
				t.Errorf("writeObjectType() = %v, want %v", got, tt.expected)
			}
		})
	}
}


func TestWriteInterfaceType(t *testing.T) {
	tests := []struct {
		name     string
		typeInfo Type
		expected string
	}{
		{
			name: "interface type with fields",
			typeInfo: Type{
				Name: "Node",
				Fields: []Field{
					{Name: "id", Type: Type{Name: "ID"}},
				},
			},
			expected: "interface Node {\n  id: ID\n}\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			writeInterfaceType(&sb, tt.typeInfo)
			if got := sb.String(); got != tt.expected {
				t.Errorf("writeInterfaceType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWriteInputObjectType(t *testing.T) {
	tests := []struct {
		name     string
		typeInfo Type
		expected string
	}{
		{
			name: "input object type with fields",
			typeInfo: Type{
				Name: "UserInput",
				InputFields: []Field{
					{Name: "name", Type: Type{Name: "String"}},
					{Name: "age", Type: Type{Name: "Int"}},
				},
			},
			expected: "input UserInput {\n  name: String\n  age: Int\n}\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			writeInputObjectType(&sb, tt.typeInfo)
			if got := sb.String(); got != tt.expected {
				t.Errorf("writeInputObjectType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestWriteUnionType(t *testing.T) {
	tests := []struct {
		name     string
		typeInfo Type
		expected string
	}{
		{
			name: "union type with possible types",
			typeInfo: Type{
				Name: "Metadata",
				PossibleTypes: []Type{
					{Name: "Video"},
					{Name: "Image"},
				},
			},
			expected: "union Metadata = Video | Image\n\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var sb strings.Builder
			writeUnionType(&sb, tt.typeInfo)
			if got := sb.String(); got != tt.expected {
				t.Errorf("writeUnionType() = %v, want %v", got, tt.expected)
			}
		})
	}
}
