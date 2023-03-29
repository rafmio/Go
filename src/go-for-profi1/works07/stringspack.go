package main

import (
	"fmt"
	s "strings"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there, youpta!")
	f("To upper: %s\n", upper)
	f("To lower: %s\n", s.ToLower("HELLO-MELLO THERE, YOPTA"))
	f("%s\n", s.Title("this wiLL be A title!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHALi"))

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))

	f("Repeat: %s\n", s.Repeat("ab", 5))
	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis"))
	f("Fields: %v\n", s.Fields("This is a string!"))
	f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))
	f("%s\n", s.Split("abcd efg", ""))

	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))
}
