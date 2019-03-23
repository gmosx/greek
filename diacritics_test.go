package greekutil_test

import (
	"reizu/pkg/greekutil"
	"testing"
)

func TestStripDiacritics(t *testing.T) {
	var tests = []struct {
		input    string
		expected string // expected result
	}{
		{"Λαϊκά", "Λαικα"},
		{"Αρνάκι άσπρο και παχύ", "Αρνακι ασπρο και παχυ"},
	}

	for _, tt := range tests {
		actual := greekutil.StripDiacritics(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %s; got %s", tt.expected, actual)
		}
	}
}
