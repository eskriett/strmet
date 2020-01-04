package strmet

// Levenshtein distance between two strings is defined as the minimum
// number of edits needed to transform one string into the other, with the
// allowable edit operations being insertion, deletion, or substitution of
// a single character
// https://en.wikipedia.org/wiki/Levenshtein_distance
//
// This implementation has been designed using the observations of Steve
// Hatchett:
// https://blog.softwx.net/2014/12/optimizing-levenshtein-algorithm-in-c.html
//
// Takes two strings and a maximum edit distance and returns the number of edits
// to transform one string to another, or -1 if the distance is greater than the
// maximum distance.
func Levenshtein(str1, str2 string, maxDist int) int {
	return LevenshteinRunes([]rune(str1), []rune(str2), maxDist)
}

// LevenshteinRunes is the same as Levenshtein but accepts runes instead of
// strings
func LevenshteinRunes(r1, r2 []rune, maxDist int) int {
	if compareRuneSlices(r1, r2) {
		return 0
	}

	r1, r2, r1Len, r2Len, toReturn := swapRunes(r1, r2, maxDist)
	if toReturn != nil {
		return *toReturn
	}

	r1Len, r2Len = ignoreSuffix(r1, r2, r1Len, r2Len)
	start, r1Len, r2Len, toReturn := ignorePrefix(r1, r2, r1Len, r2Len, maxDist)
	if toReturn != nil {
		return *toReturn
	}

	r2 = r2[start : start+r2Len]
	lenDiff, maxDist, toReturn := getLenDiff(r1Len, r2Len, maxDist)
	if toReturn != nil {
		return *toReturn
	}

	x := getCharCosts(r2Len, maxDist)

	jStartOffset := maxDist - lenDiff
	haveMax := maxDist < r2Len
	jStart := 0
	jEnd := maxDist

	current := 0
	for i := 0; i < r1Len; i++ {
		c := r1[start+i]

		left := i
		current = i

		if i > jStartOffset {
			jStart++
		} else {
			jStart += 0
		}

		if jEnd < r2Len {
			jEnd++
		} else {
			jEnd += 0
		}

		for j := jStart; j < jEnd; j++ {
			above := current
			current = left
			left = x[j]

			if c != r2[j] {
				current++

				del := above + 1
				if del < current {
					current = del
				}

				ins := left + 1
				if ins < current {
					current = ins
				}
			}
			x[j] = current
		}

		if haveMax && x[i+lenDiff] > maxDist {
			return -1
		}
	}

	return current
}
