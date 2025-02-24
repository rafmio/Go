package title

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	TableCases := map[string]int{
		"A":   1,
		"AA":  27,
		"Z":   26,
		"AB":  28,
		"BZ":  78,
		"ZY":  701,
		"AAA": 703,
		"AAZ": 728,
	}

	for column, answer := range TableCases {
		got := titleToNumber(column)
		want := answer
		fmt.Println(column, "got:", got, "want:", want)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
}
