package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr2IntValidation(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		expected1 int
		expected2 error
	}{
		{
			name:      "Zero",
			value:     "0",
			expected1: 0,
			expected2: nil,
		},
		{
			name:      "Negative",
			value:     "-500",
			expected1: -500,
			expected2: nil,
		},
		{
			name:      "Positive",
			value:     "1555000",
			expected1: 1555000,
			expected2: nil,
		},
		{
			name:      "String",
			value:     "Word",
			expected1: 0,
			expected2: ErrStr2IntNotNumber,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2 := Str2IntValidation(tt.value)
			assert.Equal(t, tt.expected1, result1)
			assert.Equal(t, tt.expected2, result2)
		})
	}
}

func TestNameValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected error
	}{
		{
			name:     "Normal1",
			value:    "agyfyp",
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    "suxinidiliwopak",
			expected: nil,
		},
		{
			name:     "Empty",
			value:    "",
			expected: ErrNameEmpty,
		},
		{
			name:     "TooSmall",
			value:    "a",
			expected: ErrNameTooSmall,
		},
		{
			name:     "TooBig",
			value:    "obitebakydilociobitebakydilocisa1",
			expected: ErrNameTooBig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NameValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAgeValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected error
	}{
		{
			name:     "Normal1",
			value:    85,
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    17,
			expected: nil,
		},
		{
			name:     "TooSmall",
			value:    11,
			expected: ErrValidationAgeTooSmall,
		},
		{
			name:     "TooBig",
			value:    86,
			expected: ErrValidationAgeTooBig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AgeValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSexValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected error
	}{
		{
			name:     "Normal1",
			value:    "man",
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    "woman",
			expected: nil,
		},
		{
			name:     "Invalid1",
			value:    "boy",
			expected: ErrValidationSexInvalid,
		},
		{
			name:     "Invalid2",
			value:    "girl",
			expected: ErrValidationSexInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SexValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCourseValidation(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		expected error
	}{
		{
			name:     "Normal1",
			value:    1,
			expected: nil,
		},
		{
			name:     "Normal2",
			value:    5,
			expected: nil,
		},
		{
			name:     "TooSmall",
			value:    0,
			expected: ErrValidationCourseTooSmall,
		},
		{
			name:     "TooBig",
			value:    110,
			expected: ErrValidationCourseTooBig,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CourseValidation(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}
