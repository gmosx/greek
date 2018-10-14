package greekutil_test

import (
	"reizu/pkg/greekutil"
	"testing"
)

func TestGreeklish(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"Λαϊκά", "Laika"},
		{"Αρνάκι άσπρο και παχύ", "Arnaki aspro kai paxy"},
		{"Γιώργος Μοσχοβίτης", "Giwrgos Mosxobiths"},
	}

	for _, tt := range tests {
		actual := greekutil.Greeklish(tt.input)
		if actual != tt.expected {
			t.Errorf("Expected %s; got %s", tt.expected, actual)
		}
	}
}
