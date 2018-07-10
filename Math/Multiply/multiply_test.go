package multiply

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply321(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		value    int
		expected int
	}{
		{"test2", 0, 0},
		{"test2", 5, 1605},
		{"test2", -5, -1605},
		{"test1", 100, 32100},
		{"test2", -100, -32100},
		{"test2", 5555, 1783155},
		{"test2", -5555, -1783155},
	}
	for _, tt := range tests {
		test := tt
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, Multiply321(test.value))
		})
	}
}
