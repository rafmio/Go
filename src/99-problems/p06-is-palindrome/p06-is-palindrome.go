// https://github.com/shekhargulati/99-problems/blob/master/java8/README.md
// Find out whether a list is a palindrome
package palindrome

import (
	"reflect"
)

func IsPalindrome(sls []string) bool {
	reversedSls := make([]string, len(sls))

	for i, j := 0, len(sls)-1; i < j; i, j = i+1, j-1 {
		reversedSls[i], reversedSls[j] = sls[j], sls[i]
	}

	if len(sls)%2 != 0 {
		middleIdx := len(sls) / 2
		reversedSls[middleIdx] = sls[middleIdx]
	}

	return reflect.DeepEqual(sls, reversedSls)
}
