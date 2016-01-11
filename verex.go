package verbalregex

import (
	"fmt"
	"regexp"
)


type VerEx struct {
	prefixes       string
	suffixes       string
	source         string
	replaceLimit   int
	modifiers      string
	lastAdded      string
	Regex          *regexp.Regexp
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

	return  ve
}

func (ve *VerEx) GetModifiers() string {

	if ve.modifiers != "" {
		return fmt.Sprintf(`(?%s)`, ve.modifiers)
	}

	return ""
}

func (ve *VerEx) Maybe(str string) *VerEx {

	ve.add(fmt.Sprintf("(%s)?", str))
	
	return ve
}

func (ve *VerEx) Any(str string) *VerEx {

	ve.add(fmt.Sprintf("[%s]", str))

	return ve
}

func (ve *VerEx) Anything() *VerEx {

	ve.add("(.*)")

	return ve
}

func (ve *VerEx) AnythingBut(value string) *VerEx {

	ve.add(fmt.Sprintf("([^%s]*)", value))

	return ve
}

func (ve *VerEx) Something() *VerEx {

	ve.add("(.+)")

	return ve
}

func (ve *VerEx) Word() *VerEx {

	ve.add("\\w+")

	return ve
}

func (ve *VerEx) Tab() *VerEx {

	ve.add("\\t")

	return ve
}

func (ve *VerEx) OneOrMore() *VerEx {

	ve.add("%s+")

	return ve
}

func (ve *VerEx) Whitespace() *VerEx {

	ve.add("\\s")

	return ve
}

func (ve *VerEx) Then(str string) *VerEx {

	ve.add(fmt.Sprintf("(%s)", str))

	return ve
}

func (ve *VerEx) Find(str string) *VerEx {
	return ve.Then(str)
}

func (ve *VerEx) Replace(src, repl string)  string {
	return ve.MustCompile().Regex.ReplaceAllString(src, repl)
}

func (ve *VerEx) Test(value string) bool {
	return ve.MustCompile().Regex.MatchString(value)
}

func (ve *VerEx) AddModifiers(value string) *VerEx {
	ve.modifiers = value

	return ve
}

func (ve *VerEx) MustCompile() *VerEx {

	ve.Regex =  regexp.MustCompile(ve.GetRegex())

	return ve
}

func (ve *VerEx) Clear() *VerEx {

	ve.prefixes = ``
	ve.suffixes = ``
	ve.source = ``
	ve.modifiers = ``
	ve.Regex = &regexp.Regexp{}

	return ve
}

func (ve VerEx) GetRegex() string {
	return ve.GetModifiers() + ve.prefixes + ve.source + ve.suffixes
}

func (ve VerEx) String() string {
	return ve.GetRegex()
}

func(ve VerEx) GetNativeRegexp() *regexp.Regexp {
	return ve.Regex
}


