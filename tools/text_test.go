package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiLine(t *testing.T) {
	tests := []struct {
		name     string
		value    []string
		expected string
	}{
		{
			name:     "Normal1",
			value:    []string{"Word1", "Word2", "Word3"},
			expected: "Word1\nWord2\nWord3",
		},
		{
			name:     "Normal2",
			value:    []string{"Word1"},
			expected: "Word1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MultiLine(tt.value...)
			assert.Equal(t, tt.expected, result)
		})
	}
}
