package namegen

import (
	"sort"
	"testing"
)

func TestDictionarySanity(t *testing.T) {
	for i := 0; i < int(numDicts); i++ {
		d := dicts[i]

		// Check to make sure the dictionary lists are sorted.
		sortedByName := sort.SliceIsSorted(d, func(i, j int) bool {
			return d[i] < d[j]
		})
		if !sortedByName {
			t.Errorf("Dictionary %d is not sorted!", i)
		}

		// Check to make sure there are no duplicates within a dictionary list.
		for j := 1; j < len(d); j++ {
			if d[j] == d[j-1] {
				t.Errorf("Dictionary %d has a duplicate entry: %s", i, d[j])
			}
		}
	}
}
