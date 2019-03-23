package greekutil

import (
	"testing"
)

func TestGreeklish(t *testing.T) {
	var tests = []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Λαϊκά", "Laika"},
		{"Αρνάκι άσπρο και παχύ", "Arnaki aspro kai paxy"},
		{"Γιώργος Μοσχοβίτης", "Giwrgos Mosxobiths"},
	}

	for _, tt := range tests {
		got := Greeklish(tt.in)
		if got != tt.want {
			t.Errorf("Greeklish(%q) = %q; want %q", tt.in, got, tt.want)
		}
	}
}
