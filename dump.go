package logr

import (
	"fmt"
	"strings"

	"github.com/aryann/difflib"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/color"
)

// Dump pretty prints a data structure to the log
func Dump(obj interface{}) {
	color.Set(color.FgCyan)
	spew.Dump(obj)
	color.Unset()
}

// DumpDiff pretty prints differences between data structures to the log
func DumpDiff(left interface{}, right interface{}) {

	l := strings.Split(spew.Sdump(left), "\n")
	r := strings.Split(spew.Sdump(right), "\n")

	DiffStrings(l, r)
}

//DiffString calculates the differences between two strings.  Supports multiline strings.
func DiffString(l, r string) {
	// If the string contains a newline, then split on newlines first
	if strings.Contains(l, "\n") {
		DiffStrings(strings.Split(l, "\n"), strings.Split(r, "\n"))
	} else {
		DiffStrings([]string{l}, []string{r})
	}
}

// DiffStrings displays a source control style string diff in the log
func DiffStrings(l, r []string) {

	diffs := difflib.Diff(l, r)

	for _, diff := range diffs {
		var prefix string
		switch diff.Delta {
		case difflib.Common:
			prefix = "    "
			color.Set(color.FgGreen)
		case difflib.LeftOnly:
			prefix = " << "
			color.Set(color.FgRed)
		case difflib.RightOnly:
			prefix = " >> "
			color.Set(color.FgYellow)
		}
		fmtStr := fmt.Sprintf("%s %s", prefix, diff.Payload)
		fmt.Println(fmtStr)
	}
	color.Unset()
}
