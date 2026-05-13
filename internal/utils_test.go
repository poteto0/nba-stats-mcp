package internal

import "testing"

func TestParseSeasonBeginYear(t *testing.T) {
	tests := []struct {
		input       string
		expected    int
		expectError bool
	}{
		{"2023-24", 2023, false},
		{"1999-00", 1999, false},
		{"2000-01", 2000, false},
		{"invalid", 0, true},
		{"2023", 0, true},
	}

	for _, test := range tests {
		result, err := ParseSeasonBeginYear(test.input)
		if (err != nil) != test.expectError {
			t.Errorf("ParseSeasonBeginYear(%q) error = %v, expected error: %v", test.input, err, test.expectError)
			continue
		}
		if result != test.expected {
			t.Errorf("ParseSeasonBeginYear(%q) = %d, expected %d", test.input, result, test.expected)
		}
	}
}

func TestFormatSeasonYear(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{2023, "2023-24"},
		{1999, "1999-00"},
		{2000, "2000-01"},
	}

	for _, test := range tests {
		result := FormatSeasonYear(test.input)
		if result != test.expected {
			t.Errorf("FormatSeasonYear(%d) = %q, expected %q", test.input, result, test.expected)
		}
	}
}
