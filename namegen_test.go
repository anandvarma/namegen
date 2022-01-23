package namegen

import (
	"strings"
	"testing"
)

type getPostfixTest struct {
	arg1     int
	arg2     PostfixEncoding
	arg3     int
	expected string
}

var getPostfixTests = []getPostfixTest{
	// 1234 -> 1234
	{1234, Numeric, 0, ""},
	{1234, Numeric, 4, "1234"},
	{1234, Numeric, 2, "34"},
	{1234, Numeric, 6, "001234"},
	// 43981 -> abcd
	{43981, Hex, 0, ""},
	{43981, Hex, 4, "abcd"},
	{43981, Hex, 2, "cd"},
	{43981, Hex, 6, "00abcd"},
	// 623741435 -> abcxyz
	{623741435, Alphanumeric, 0, ""},
	{623741435, Alphanumeric, 6, "abcxyz"},
	{623741435, Alphanumeric, 3, "xyz"},
	{623741435, Alphanumeric, 8, "00abcxyz"},
}

func TestGetPostfix(t *testing.T) {
	for _, test := range getPostfixTests {
		if output := getPostfix(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

type getNameForIdTest struct {
	arg1 int
	arg2 options
}

var getNameForIdTests = []getNameForIdTest{
	{1111, options{[][]string{adjectives, animals}, "-", 0, Numeric}},
	{1234, options{[][]string{adjectives, animals}, "-", 4, Numeric}},
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TestGetNameForId(t *testing.T) {
	for _, test := range getNameForIdTests {
		output := strings.Split(GetNameForId(test.arg1, test.arg2), test.arg2.delimiter)
		// Check to make sure the correct dictionaries were used and in the right order.
		for i := range test.arg2.dicts {
			if !stringInSlice(output[i], test.arg2.dicts[i]) {
				t.Errorf("Dict %d doesn't contain %s", i, output[i])
			}
		}
		// Check the postfix.
		if test.arg2.postfixLen > 0 {
			if want := len(test.arg2.dicts) + 1; len(output) != want {
				t.Errorf("Unexpected output:%v for %d dicts", output, want)
			}
			if postfix := output[len(output)-1]; len(postfix) != test.arg2.postfixLen {
				t.Errorf("Unexpected postfix: %s for postfixLen: %d", postfix, test.arg2.postfixLen)
			}
		} else {
			if want := len(test.arg2.dicts); len(output) != want {
				t.Errorf("Unexpected output:%v for %d dicts", output, want)
			}
		}
	}
}
