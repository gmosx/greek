package greek

import (
	"testing"
)

func TestStripDiacritics(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"Λαϊκά", "Λαικα"},
		{"Αρνάκι άσπρο και παχύ", "Αρνακι ασπρο και παχυ"},
	}

	for _, tt := range tests {
		got := StripDiacritics(tt.in)
		if got != tt.want {
			t.Errorf("StipDiacritics(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
