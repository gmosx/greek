package greek

import (
	"regexp"
)

// TODO: introduce a simplified stripDiacriticsRules tables, only modern Greek.

var stripDiacriticsRules = []rule{
	{regexp.MustCompile("[άἀἁἂἃἄἅἆἇὰάᾀᾁᾂᾃᾄᾅᾆᾇᾰᾱᾲᾳᾴᾶᾷ]"), "α"},
	{regexp.MustCompile("[ΆἈἉἊἋἌἍἎἏᾈᾉᾊᾋᾌᾍᾎᾏᾸᾹᾺΆᾼ]"), "Α"},
	{regexp.MustCompile("[έἐἑἒἓἔἕὲέ]"), "ε"},
	{regexp.MustCompile("[ΈἘἙἚἛἜἝ]"), "Ε"},
	{regexp.MustCompile("[ήἠἡἢἣἤἥἦἧῆὴῇ]"), "η"},
	{regexp.MustCompile("[ΉἨἩἪἫἬἭἮἯ]"), "Η"},
	{regexp.MustCompile("[ίἰἱἲἳἴἵὶῖϊ]"), "ι"},
	{regexp.MustCompile("[ΊἶἷἸἹἺἻἼἽἾἿΪ]"), "Ι"},
	{regexp.MustCompile("[όὀὁὂὃὄὅὸ]"), "ο"},
	{regexp.MustCompile("[ΌὈὉὊὋὌὍ]"), "Ο"},
	{regexp.MustCompile("[ύὐὑὒὓὔὕὖὗ]"), "υ"},
	{regexp.MustCompile("[ΎὙὛὝὟ]"), "Υ"},
	{regexp.MustCompile("[ώὠὡὢὣὤὥὦὧῶ]"), "ω"},
	{regexp.MustCompile("[ΏὨὩὪὫὬὭὮὯ]"), "Ω"},
}

// StripDiacritics removes diacritics from the input string.
func StripDiacritics(s string) string {
	return applyRules(s, stripDiacriticsRules)
}
