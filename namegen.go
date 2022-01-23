package namegen

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Seed the rand package to avoid deterministic results.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// PostfixEncoding represents the various encoding styles that may be used for constructing postfixes.
type PostfixEncoding byte

const (
	Numeric      PostfixEncoding = 10 // Numbers only, Base10 encoding
	Hex          PostfixEncoding = 16 // Hex, Base16 encoding
	Alphanumeric PostfixEncoding = 36 // AlphaNumeric lower case, Base36 encoding
)

func getPostfix(num int, encoding PostfixEncoding, postfixLen int) string {
	enc := strconv.FormatInt(int64(num), int(encoding))
	if len(enc) >= postfixLen {
		// Trim to the required size.
		return enc[len(enc)-postfixLen:]
	} else {
		// Pad zeroes to the required size.
		return fmt.Sprintf("%0*s", postfixLen, enc)
	}
}

type options struct {
	dicts        [][]string
	delimiter    string
	postfixLen   int
	postfixStyle PostfixEncoding
}

func GetShort() string {
	return GetName(options{[][]string{adjectives, animals}, "-", 2, Numeric})
}

func GetLong() string {
	return GetName(options{[][]string{adjectives, colours, animals}, "-", 6, Numeric})
}

// Returns a randomized name using 'opts'.
func GetName(opts options) string {
	return GetNameForId(rand.Int(), opts)
}

// Returns a name corresponding to 'randomId', using 'opts'.
func GetNameForId(randomId int, opts options) string {
	substrs := []string{}
	for _, d := range opts.dicts {
		substrs = append(substrs, d[randomId%len(d)])
	}
	if opts.postfixLen > 0 {
		substrs = append(substrs, getPostfix(randomId, opts.postfixStyle, int(opts.postfixLen)))
	}
	return strings.Join(substrs, opts.delimiter)
}
