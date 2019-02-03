package strmet

// DamerauLevenshtein distance is a string metric for measuring the edit
// distance between two sequences:
// https://en.wikipedia.org/wiki/Damerau%E3%80%93Levenshtein_distance
//
// This implementation has been designed using the observations of Steve
// Hatchett:
// http://blog.softwx.net/2015/01/optimizing-damerau-levenshtein_15.html
//
// Takes two strings and a maximum edit distance and returns the number of edits
// to transform one string to another, or -1 if the distance is greater than the
// maximum distance.
func DamerauLevenshtein(str1, str2 string, maxDist int) int {

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

	v0 := getCharCosts(s2Len, maxDist)
	v2 := make([]int, s2Len)

	jStartOffset := maxDist - lenDiff
	haveMax := maxDist < s2Len
	jStart := 0
	jEnd := maxDist
	s1Char := s1[0]
	current := 0
	for i := 0; i < s1Len; i++ {
		prevS1Char := s1Char
		s1Char = s1[start+i]
		s2Char := s2[0]
		left := i
		current = left + 1
		nextTransCost := 0

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
			thisTransCost := nextTransCost
			nextTransCost = v2[j]
			current = left
			v2[j] = current
			left = v0[j]
			prevS2Char := s2Char
			s2Char = s2[j]
			if s1Char != s2Char {
				if left < current {
					current = left
				}
				if above < current {
					current = above
				}
				current++
				if i != 0 && j != 0 && s1Char == prevS2Char && prevS1Char ==
					s2Char {
					thisTransCost++
					if thisTransCost < current {
						current = thisTransCost
					}
				}
			}
			v0[j] = current
		}

		if haveMax && v0[i+lenDiff] > maxDist {
			return -1
		}

	}

	return current
}
