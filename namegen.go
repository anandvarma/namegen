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

type PostfixIdType byte

const (
	Numeric      PostfixIdType = 10
	AlphaNumeric PostfixIdType = 36
)

func New(dicts []DictType) *nameGen {
	return &nameGen{dicts, "-" /* delim */, Numeric /* pIdType */, 0 /* pIdLen */}
}

func NewWithPostfixId(dicts []DictType, pIdType PostfixIdType, pIdLen int) *nameGen {
	return &nameGen{dicts, "-" /* delimiter */, pIdType, pIdLen}
}

func (n *nameGen) SetDelimiter(delim string) {
	n.delim = delim
}

func (n *nameGen) SetPostfixIdLen(pIdLen int) {
	n.pIdLen = pIdLen
}

func (n nameGen) Get() string {
	return n.GetForId(rand.Int63())
}

func (n nameGen) GetForId(randNum int64) string {
	substrs := []string{}
	for _, dIdx := range n.dicts {
		d := dicts[dIdx]
		substrs = append(substrs, d[randNum%int64(len(d))])
	}
	if n.pIdLen > 0 {
		substrs = append(substrs, getPostfixId(randNum, n.pIdType, n.pIdLen))
	}
	return strings.Join(substrs, n.delim)
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

// Returns a short random name of the format "adjective-noun-xx".

// Returns a random name of the format "adjective-colour-noun-xxxxxx".

// Returns a randomized name using 'opts'.

// Returns a name corresponding to 'randomId', using 'opts'.
