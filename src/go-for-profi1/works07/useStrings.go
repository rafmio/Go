package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis wiLL be A tItle!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALIS"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALI"))
	f("Prefix: %v\n", s.HasPrefix("Mihasis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Zy"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "ze"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "ho"))
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "Z"))

	f("Repeat: %s\n", s.Repeat("Su", 5))

	f("TrimSpace: %s\n", s.TrimSpace("   \this is a  line. \n"))
	f("TrimSpace: %s\n", s.TrimLeft("  \tThis is a\t line. \n", "\n\t "))
	f("TrimSpace: %s\n", s.TrimRight("  \tThis is a\t line  \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))

	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))

	f("%s\n", s.Split("abcd efg", ""))
	f("\n")

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))
}
