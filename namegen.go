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

// PostfixIdType represents the type of encoding to use when constructing postfix ids.
type PostfixIdType byte

const (
	// Base 10 encoding.
	Numeric PostfixIdType = 10
	// Base 36 encoding.
	AlphaNumeric PostfixIdType = 36
)

// New creates and returns a new nameGen instance configured to generate unique names using the default config.
// Returns a string in the following format "adjective-animal".
func New() *nameGen {
	return NewWithDicts([]DictType{Adjectives, Animals})
}

// NewWithDicts creates and returns a new nameGen instance configured to generate unique names using the default config and a custom naming schema defined by 'dicts'.
// Returns a string in the following format "d1.word-d2.word".
func NewWithDicts(dicts []DictType) *nameGen {
	return NewWithPostfixId(dicts, Numeric /* pIdType */, 0 /* pIdLen */)
}

// NewWithPostfixId creates and returns a new nameGen instance configured to generate unique human-readable ids using the default config and 'dicts'.
// A postfix id of 'pIdLen' is generated using 'pIdType' encoding.
// Returns a string in the following format "d1.word-d2.word-postfix.id".
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
	randIdxArr := make([]int, 0, 8)
	retLen := n.pIdLen
	for _, dIdx := range n.dicts {
		d := dicts[dIdx]
		wIdx := int(randNum) % len(d)
		randIdxArr = append(randIdxArr, wIdx)
		retLen += len(d[wIdx]) + len(n.delim)
	}
	var sb strings.Builder
	sb.Grow(retLen)

	for i, dIdx := range n.dicts {
		d := dicts[dIdx]
		sb.WriteString(d[randIdxArr[i]])
		if (i < len(n.dicts)-1) || (n.pIdLen > 0) {
			sb.WriteString(n.delim)
		}
	}
	if n.pIdLen > 0 {
		getPostfixId(&sb, randNum, n.pIdType, n.pIdLen)
	}
	return sb.String()
}

func getPostfixId(sb *strings.Builder, num int64, pIdType PostfixIdType, pIdLen int) {
	// The smallest encoding we use is Base10. INT64_MAX fits within 20 chars.
	encBuf := make([]byte, 0, 20)
	encBuf = strconv.AppendInt(encBuf, num, int(pIdType))

	if len(encBuf) < pIdLen {
		// Pad zeroes to the required size. Unlikely case.
		out := fmt.Sprintf("%0*s", pIdLen, string(encBuf))
		sb.WriteString(out)
		return
	}

	// Trim to the required size.
	sb.Write(encBuf[len(encBuf)-pIdLen:])
}
