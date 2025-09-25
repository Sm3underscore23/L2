package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLong(t *testing.T) {
	testTable := []struct {
		name          string
		in            string
		expected      string
		expectedError error
	}{
		{
			name:     "simple_wb_test_1",
			in:       "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			name:     "simple_wb_test_2",
			in:       "abcd",
			expected: "abcd",
		},
		{
			name:          "simple_wb_test_3",
			in:            "45",
			expectedError: errInvalidString,
		},
		{
			name:     "simple_wb_test_5",
			in:       "",
			expected: "",
		},
		{
			name:          "simple_custom_test_1",
			in:            "2a",
			expectedError: errInvalidString,
		},
		{
			name:     "simple_custom_test_2",
			in:       "a11v3",
			expected: "aaaaaaaaaaavvv",
		},
		{
			name:     "simple_custom_test_3",
			in:       "a0b1c2",
			expected: "bcc",
		},
		{
			name:     "hard_wm_test_1",
			in:       "qwe\\4\\5",
			expected: "qwe45",
		},
		{
			name:     "hard_wm_test_2",
			in:       "qwe\\45",
			expected: "qwe44444",
		},
		{
			name:     "hard_custom_test_1",
			in:       "\\02",
			expected: "00",
		},
		{
			name:     "hard_custom_test_2",
			in:       "\\20",
			expected: "",
		},
		{
			name:          "hard_custom_test_3",
			in:            "ab4\\",
			expectedError: errInvalidString,
		},
		{
			name:     "hard_custom_test_4",
			in:       "ab\\10\\21\\310\\4",
			expected: "ab233333333334",
		},
	}

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			result, err := parse(tc.in)
			assert.Equal(t, tc.expected, result)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
