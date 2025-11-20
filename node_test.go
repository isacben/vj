package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNodeType(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		path     string
		expected NodeType
	}{
		{
			"string",
			map[string]interface{}{"name": "John"},
			"name",
			StringType,
		},
		{
			"number",
			map[string]interface{}{"age": 25.0},
			"age",
			NumberType,
		},
		{
			"bool",
			map[string]interface{}{"active": true},
			"active",
			BoolType,
		},
		{
			"object",
			map[string]interface{}{"value": nil},
			"value",
			NullType,
		},
		{
			"array",
			[]interface{}{1, 2, 3, 4},
			"",
			ArrayType,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := BuildTree(tt.data, "", nil)
			node, exists := tree.GetNode(tt.path)

			assert.True(t, exists, "Node should exist")
			assert.Equal(t, tt.expected, getNodeType(node.Value))
		})
	}
}

func TestGetDepth(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{"root depth", "", 0},
		{"root depth", "user", 0},
		{"simple object depth", "user.name", 1},
		{"nested object depth", "user.address.street_name", 2},
		{"simple array depth", "user.addresses[0]", 2},
		{"nested array depth", "user.addresses[10].stree_name", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, getDepth(tt.path))
		})
	}
}

func TestBuildChildPath(t *testing.T) {
	tests := []struct {
		name     string
		basePath string
		key      string
		isArray  bool
		expected string
	}{
		{"test object", "user", "name", false, "user.name"},
		{"test nested object", "user.address", "street_name",
			false, "user.address.street_name"},
		{"test array", "user.addresses[0]", "street_name",
			true, "user.addresses[0][street_name]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, buildChildPath(
				tt.basePath, tt.key, tt.isArray),
			)
		})
	}
}

func TestIsNested(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected bool
	}{
		{
			"nested object",
			map[string]interface{}{"name": "John"},
			true,
		},
		{
			"nested array",
			[]interface{}{1, 2, 3, 4},
			true,
		},
		{
			"not nested string",
			"some value",
			false,
		},
		{
			"not nested number",
			1.0,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isNested(tt.value))
		})
	}
}
