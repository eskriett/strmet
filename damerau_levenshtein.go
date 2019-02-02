package strmet

// The Damerau–Levenshtein distance is a string metric for measuring the edit
// distance between two sequences:
// https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
//
// This implementation has been designed using the observations of Steve
// Hatchett:
// http://blog.softwx.net/2015/01/optimizing-damerau-levenshtein_15.html
//
// Takes two strings and a maximum edit distance and returns the number of edits
// to transform one string to another, or -1 if the distance is greater than the
// maximum distance.
func DamerauLevenshtein(str1, str2 string, maxDist int) int {

	if maxDist < 0 {
		return -1
	}

	s1 := []rune(str1)
	s2 := []rune(str2)

	s1Len := len(s1)
	s2Len := len(s2)

	if s1Len > s2Len {
		s1, s2 = s2, s1
		s1Len, s2Len = s2Len, s1Len
	}

	if s1Len == 0 {
		if s2Len <= maxDist {
			return s2Len
		}
		return -1
	}

	if s1Len > s2Len {
		s1, s2 = s2, s1
		s1Len, s2Len = s2Len, s1Len
	}

	// Ignore suffixes common to both strings
	for s1Len > 0 && s1[s1Len-1] == s2[s2Len-1] {
		s1Len--
		s2Len--
	}

	start := 0
	if s1[start] == s2[start] || s1Len == 0 {

		// Ignore prefix common to both strings
		for start < s1Len && s1[start] == s2[start] {
			start++
		}
		s1Len -= start
		s2Len -= start

		if s1Len == 0 {
			if s2Len <= maxDist {
				return s2Len
			}
			return -1
		}
		s2 = s2[start : start+s2Len]
	}
	lenDiff := s2Len - s1Len

	if maxDist > s2Len {
		maxDist = s2Len
	} else if lenDiff > maxDist {
		return -1
	}

	v0 := make([]int, s2Len)
	v2 := make([]int, s2Len)

	j := 0
	for j = 0; j < maxDist; j++ {
		v0[j] = j + 1
	}
	for ; j < s2Len; j++ {
		v0[j] = maxDist + 1
	}

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
			jStart += 1
		} else {
			jStart += 0
		}

		if jEnd < s2Len {
			jEnd += 1
		} else {
			jEnd += 0
		}

		for j = jStart; j < jEnd; j++ {
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
