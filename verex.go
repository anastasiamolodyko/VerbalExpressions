package verbalregex

import (
	"fmt"
	"regexp"
	"strings"
)

type VerEx struct {
	prefixes     string
	suffixes     string
	source       string
	replaceLimit int
	modifiers    []string
	lastAdded    string
	Regex        *regexp.Regexp
	compiled     bool
}

func (ve *VerEx) StartOfLine(strict bool) *VerEx {

	if strict == true {
		ve.prefixes = "^"
	}

	return ve
}

func (ve *VerEx) EndOfLine(strict bool) *VerEx {

	if strict == true {
		ve.suffixes = "$"
	}

	return ve
}

func (ve *VerEx) add(value string) *VerEx {

	ve.lastAdded = value
	ve.source = fmt.Sprintf("%s%s", ve.source, value)

	return ve
}

func (ve *VerEx) GetModifiers() string {
	return strings.Join(ve.modifiers, "")
}

func (ve *VerEx) Maybe(value string) *VerEx {
	return ve.add(fmt.Sprintf("(?:%s)?", sanitize(value)))
}

func (ve *VerEx) Any(value string) *VerEx {
	return ve.add(fmt.Sprintf("[%s]", sanitize(value)))
}

func (ve *VerEx) Anything() *VerEx {
	return ve.add("(?:.*)")
}

func (ve *VerEx) AnythingBut(value string) *VerEx {
	return ve.add(fmt.Sprintf("(?:[^%s]*)", sanitize(value)))
}

func (ve *VerEx) Something() *VerEx {
	return ve.add("(?:.+)")
}

func (ve *VerEx) SomethingBut(value string) *VerEx {
	return ve.add(fmt.Sprintf("(?:[^%s]+)", sanitize(value)))
}

func (ve *VerEx) Word() *VerEx {
	return ve.add(`\w+`)
}

func (ve *VerEx) Tab() *VerEx {
	return ve.add(`\t`)
}

func (ve *VerEx) OneOrMore() *VerEx {
	return ve.add("+")
}

func (ve *VerEx) ZeroOrMore() *VerEx {
	return ve.add("*")
}

func (ve *VerEx) Whitespace() *VerEx {
	return ve.add(`\s`)
}

func (ve *VerEx) LineBreak() *VerEx {
	return ve.add(`(?:(?:\n)|(?:\r\n))`)
}

func (ve *VerEx) Br() *VerEx {
	return ve.LineBreak()
}

func (ve *VerEx) Then(value string) *VerEx {
	return ve.add(fmt.Sprintf("(?:%s)", sanitize(value)))
}

func (ve *VerEx) Find(str string) *VerEx {
	return ve.Then(str)
}

func (ve *VerEx) AddModifier(value string) *VerEx {

	if !inSlice(ve.modifiers, value) {
		ve.modifiers = append(ve.modifiers, value)
	}

	return ve
}

func (ve *VerEx) RemoveModifier(value string) *VerEx {

	if inSlice(ve.modifiers, value) {

		for idx, el := range ve.modifiers {
			if el == value {
				ve.modifiers = append(ve.modifiers[:idx], ve.modifiers[idx+1])
			}
		}

	}

	return ve
}

func (ve *VerEx) MustCompile() *VerEx {

	if ve.compiled == false {
		return ve.Recompile()
	}

	return ve
}

func (ve *VerEx) Recompile() *VerEx {

	ve.Regex = regexp.MustCompile(ve.GetRegex())
	ve.compiled = true

	return ve
}

func (ve *VerEx) Fresh() *VerEx {

	ve.prefixes = ``
	ve.suffixes = ``
	ve.source = ``
	ve.modifiers = nil
	ve.Regex = &regexp.Regexp{}
	ve.compiled = false

	return ve
}

func (ve VerEx) GetRegex() string {
	return ve.GetModifiers() + ve.prefixes + ve.source + ve.suffixes
}

func (ve VerEx) String() string {
	return ve.GetRegex()
}

func (ve VerEx) GetNativeRegexp() *regexp.Regexp {
	return ve.Regex
}

func (ve *VerEx) BeginCapture() *VerEx {

	ve.suffixes += `)`
	ve.add(`(`)

	return ve
}

func (ve *VerEx) EndCapture() *VerEx {

	ve.suffixes = strings.Replace(ve.suffixes, ")", "", 1)
	ve.add(`)`)

	return ve
}

func sanitize(value string) string {
	return regexp.QuoteMeta(value)
}

func inSlice(slice []string, needle string) bool {
	for _, val := range slice {
		if val == needle {
			return true
		}
	}

	return false
}
