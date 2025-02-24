// https://medium.com/nerd-for-tech/writing-unit-tests-in-golang-part-1-introducing-testify-c0d458442412

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string
	input          string
	expoetedResult bool
	hasError       bool
}

var testTable []testCase = []testCase{
	{
		name:           "all lowercase string, input",
		input:          "aabbcc",
		expoetedResult: true,
		hasError:       false,
	},

	{
		name:           "some uppercase string input",
		input:          "aAbBcC",
		expoetedResult: false,
		hasError:       false,
	},

	{
		name:           "input string contains a number",
		input:          "aabbc1",
		expoetedResult: false,
		hasError:       true,
	},
}

func TestIsAllLowerCase(t *testing.T) {

	for _, test := range testTable {
		actual, err := IsAllLowerCase(test.input)
		assert.Equal(t, test.expoetedResult, actual, test.name)

		if test.hasError {
			assert.NotNil(t, err, test.name)
		} else {
			assert.Nil(t, err, test.name)
		}
	}
}

func BenchmarkIsAllLowerCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range testTable {
			IsAllLowerCase(value.input)
		}
	}
}

// go test
// go test -v
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
// go test -bench=.
// go test -bench=FuncName
