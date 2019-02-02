package strmet

// The Levenshtein distance between two strings is defined as the minimum
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

	// Ignore suffixes common to both strings
	for s1Len > 0 && s1[s1Len-1] == s2[s2Len-1] {
		s1Len--
		s2Len--
	}

	// Ignore prefix common to both strings
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

	x := make([]int, s2Len)

	var j int
	for j = 0; j < maxDist; j++ {
		x[j] = j + 1
	}
	for ; j < s2Len; j++ {
		x[j] = maxDist + 1
	}

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
			jStart += 1
		} else {
			jStart += 0
		}

		if jEnd < s2Len {
			jEnd += 1
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
