package namegen

import (
	"sort"
	"testing"
)

func TestDictionarySanity(t *testing.T) {
	t.Run("ShouldBeSortedAndNotHaveDuplicateOnEachDictionary", func(t *testing.T) {
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
	})

	t.Run("ShouldNotHaveDuplicateAcrossDictionaries", func(t *testing.T) {
		allDicts := []string{}

		// Append all dictionaries into one
		for _, v := range dicts {
			allDicts = append(allDicts, v...)
		}

		// Check if there is any duplicate in all appended dictionaries
		isExist := make(map[string]bool)
		for _, vv := range allDicts {
			if isExist[vv] {
				t.Errorf("Duplicate entry across dictionaries: %s", vv)
			}
			isExist[vv] = true
		}
	})
}
