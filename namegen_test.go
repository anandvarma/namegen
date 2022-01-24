package namegen

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

type getPostfixIdTest struct {
	arg1     int64
	arg2     PostfixIdType
	arg3     int
	expected string
}

var getPostfixIdTests = []getPostfixIdTest{
	// 1234 -> 1234
	{1234, Numeric, 0, ""},
	{1234, Numeric, 4, "1234"},
	{1234, Numeric, 2, "34"},
	{1234, Numeric, 6, "001234"},
	// 623741435 -> abcxyz
	{623741435, AlphaNumeric, 0, ""},
	{623741435, AlphaNumeric, 6, "abcxyz"},
	{623741435, AlphaNumeric, 3, "xyz"},
	{623741435, AlphaNumeric, 8, "00abcxyz"},
}

func TestGetPostfix(t *testing.T) {
	for _, test := range getPostfixIdTests {
		if output := getPostfixId(test.arg1, test.arg2, test.arg3); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

var getTests = []nameGen{
	*NewWithPostfixId([]DictType{Adjectives, Animals}, Numeric, 0),
	*NewWithPostfixId([]DictType{Adjectives, Animals}, Numeric, 4),
	*NewWithPostfixId([]DictType{Adjectives, Animals}, AlphaNumeric, 6),
	*NewWithPostfixId([]DictType{Adjectives, Colors, Animals}, Numeric, 0),
	*NewWithPostfixId([]DictType{Adjectives, Colors, Animals}, Numeric, 4),
	*NewWithPostfixId([]DictType{Adjectives, Colors, Animals}, AlphaNumeric, 6),
}

func TestGet(t *testing.T) {
	for _, test := range getTests {
		output := strings.Split(test.Get(), test.delim)
		// Check to make sure the correct dictionaries were used and in the right order.
		for i := range test.dicts {
			if !stringInSlice(output[i], dicts[test.dicts[i]]) {
				t.Errorf("Dict %d doesn't contain %s", i, output[i])
			}
		}
		// Check the postfix.
		if test.pIdLen > 0 {
			if want := len(test.dicts) + 1; len(output) != want {
				t.Errorf("Unexpected output:%v for %d dicts", output, want)
			}
			if postfix := output[len(output)-1]; len(postfix) != test.pIdLen {
				t.Errorf("Unexpected postfix: %s for postfixLen: %d", postfix, test.pIdLen)
			}
		} else {
			if want := len(test.dicts); len(output) != want {
				t.Errorf("Unexpected output:%v for %d dicts", output, want)
			}
		}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func TestSetters(t *testing.T) {
	// Check default values.
	n := New()
	if (n.dicts[0] != Adjectives) || (n.dicts[1] != Animals) {
		t.Errorf("Unexpected default dictionary schema: %v", n.dicts)
	}
	if n.delim != "-" {
		t.Errorf("Unexpected default delimiter: %s", n.delim)
	}
	if n.pIdLen != 0 {
		t.Errorf("Unexpected default prefix length: %d", n.pIdLen)
	}

	// Validate the setters.
	n.SetDelimiter(".")
	if n.delim != "." {
		t.Errorf("Unexpected delimiter: %s", n.delim)
	}
	n.SetPostfixIdLen(7)
	if n.pIdLen != 7 {
		t.Errorf("Unexpected prefix length: %d", n.pIdLen)
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New()
	}
}

func BenchmarkGet(b *testing.B) {
	ngen := New()
	for n := 0; n < b.N; n++ {
		ngen.Get()
	}
}

func BenchmarkGetForId(b *testing.B) {
	ngen := New()
	rand.Seed(time.Now().UnixNano())
	rn := rand.Int63()
	for n := 0; n < b.N; n++ {
		ngen.GetForId(rn)
	}
}

func BenchmarkGetWithPostfix(b *testing.B) {
	ngen := New()
	ngen.SetPostfixIdLen(8)
	for n := 0; n < b.N; n++ {
		ngen.Get()
	}
}

func BenchmarkGetPostfixId(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rn := rand.Int63()
	// Note: This benchmark targets the most common case of trimming the encoded number.
	for n := 0; n < b.N; n++ {
		getPostfixId(rn, Numeric, 4)
	}
}
