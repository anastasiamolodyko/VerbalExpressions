##GoVerbalExpressions
- ported from [VerbalExpressions](https://github.com/VerbalExpressions/JSVerbalExpressions)

Hot to use


```go

package main

import (
	"fmt"
	"github.com/galiievskyi/verbalregex"
)

func main() {
	ve := verbalregex.VerEx{}

	ve.StartOfLine(true).Then(`http`).Maybe(`s`).Then(`://`).Maybe(`www.`).AnythingBut(` `).EndOfLine(true)

	if ve.Test(`https://www.google.com`) {
		fmt.Println(`valid url`)
	} else {
		fmt.Println(`invalid url`)
	}

	//replace
	fmt.Println(ve.Clear().Find(`a(x*)b`).Replace(`-ab-axxb-`, "T")) // -T-T-

	//get native regexp.Regexp

	ve.GetNativeRegexp() // return *regexp.Regexp
}

```

## Other Implementations
You can see an up to date list of all ports on [VerbalExpressions.github.io](http://VerbalExpressions.github.io).
- [Javascript](https://github.com/jehna/VerbalExpressions)
- [Ruby](https://github.com/VerbalExpressions/RubyVerbalExpressions)
- [C#](https://github.com/VerbalExpressions/CSharpVerbalExpressions)
- [Python](https://github.com/VerbalExpressions/PythonVerbalExpressions)
- [Java](https://github.com/VerbalExpressions/JavaVerbalExpressions)
- [C++](https://github.com/VerbalExpressions/CppVerbalExpressions)


