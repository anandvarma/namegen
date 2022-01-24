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

type nameGen struct {
	dicts   []DictType
	delim   string
	pIdType PostfixIdType
	pIdLen  int
}

// The type of encoding to use when constructing postfix ids.
type PostfixIdType byte

const (
	Numeric      PostfixIdType = 10 // Base 10 encoding
	AlphaNumeric PostfixIdType = 36 // Base 36 encoding
)

// Creates and retuns a new nameGen instance configured to generate unique names using the default config.
// Returns a string in the following format "adjective-animal".
func New() *nameGen {
	return NewWithDicts([]DictType{Adjectives, Animals})
}

// Creates and retuns a new nameGen instance configured to generate unique names using the default config and a custom naming schema defined by 'dicts'.
// Returns a string in the following format "d1.word-d2.word".
func NewWithDicts(dicts []DictType) *nameGen {
	return NewWithPostfixId(dicts, Numeric /* pIdType */, 0 /* pIdLen */)
}

// Creates and retuns a new nameGen instance configured to generate unique human-readable ids using the default config and 'dicts'.
// A postfix id of 'pIdLen' is generated using 'pIdType' encoding.
// Retuns a string in the following format "d1.word-d2.word-postfix.id".
func NewWithPostfixId(dicts []DictType, pIdType PostfixIdType, pIdLen int) *nameGen {
	return &nameGen{dicts, "-" /* delimiter */, pIdType, pIdLen}
}

// Sets the delimiter to 'delim'.
func (n *nameGen) SetDelimiter(delim string) {
	n.delim = delim
}

// Sets the postfix length to 'pIdLen'.
// If we encounter collisions, we may use this to expand the ID space for increased randomness.
func (n *nameGen) SetPostfixIdLen(pIdLen int) {
	n.pIdLen = pIdLen
}

// Generates and returns a random name, based on the configuration.
// This function uses math/rand to generate random ids. To bring your own random numbers, use the GetForId(int64) call instead.
func (n nameGen) Get() string {
	return n.GetForId(rand.Int63())
}

// Generates and returns a random name corresponding to 'randNum', based on the configuration.
// This API is provided to allow the caller to bring their own random numbers instead of relying on math/rand that Get() uses otherwise.
func (n nameGen) GetForId(randNum int64) string {
	var sBuf strings.Builder

	for i, dIdx := range n.dicts {
		d := dicts[dIdx]
		sBuf.WriteString(d[randNum%int64(len(d))])
		if (i < len(n.dicts)-1) || (n.pIdLen > 0) {
			sBuf.WriteString("-")
		}
	}
	if n.pIdLen > 0 {
		sBuf.WriteString(getPostfixId(randNum, n.pIdType, n.pIdLen))
	}
	return sBuf.String()
}

func getPostfixId(num int64, pIdType PostfixIdType, pIdLen int) string {
	enc := strconv.FormatInt(num, int(pIdType))
	if len(enc) >= pIdLen {
		// Trim to the required size.
		return enc[len(enc)-pIdLen:]
	} else {
		// Pad zeroes to the required size.
		return fmt.Sprintf("%0*s", pIdLen, enc)
	}
}
