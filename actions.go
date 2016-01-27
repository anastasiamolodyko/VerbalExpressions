package verbalregex

func (ve *VerEx) Replace(src, repl string)  string {
	return ve.MustCompile().Regex.ReplaceAllString(src, repl)
}

func (ve *VerEx) Test(value string) bool {
	return ve.MustCompile().Regex.MatchString(value)
}

func (ve *VerEx) Capture(value string) [][]string {
	return ve.MustCompile().Regex.FindAllStringSubmatch(value, -1)
}

