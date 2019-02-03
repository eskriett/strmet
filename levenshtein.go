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

	s1, s2, s1Len, s2Len, toReturn := getRunes(str1, str2, maxDist)
	if toReturn != nil {
		return *toReturn
	}

	s1Len, s2Len = ignoreSuffix(s1, s2, s1Len, s2Len)
	start, s1Len, s2Len, toReturn := ignorePrefix(s1, s2, s1Len, s2Len, maxDist)
	if toReturn != nil {
		return *toReturn
	}

	s2 = s2[start : start+s2Len]
	lenDiff, maxDist, toReturn := getLenDiff(s1Len, s2Len, maxDist)
	if toReturn != nil {
		return *toReturn
	}

	x := getCharCosts(s2Len, maxDist)

	jStartOffset := maxDist - lenDiff
	haveMax := maxDist < s2Len
	jStart := 0
	jEnd := maxDist

	current := 0
	for i := 0; i < s1Len; i++ {
		c := s1[start+i]

		left := i
		current = i

		if i > jStartOffset {
			jStart++
		} else {
			jStart += 0
		}

		if jEnd < s2Len {
			jEnd++
		} else {
			jEnd += 0
		}

		for j := jStart; j < jEnd; j++ {
			above := current
			current = left
			left = x[j]

			if c != s2[j] {
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
