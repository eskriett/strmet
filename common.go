package strmet

func getCharCosts(length, maxDist int) []int {

	x := make([]int, length)

	i := 0
	for ; i < maxDist; i++ {
		x[i] = i + 1
	}
	for ; i < length; i++ {
		x[i] = maxDist + 2
	}

	return x
}

func getLenDiff(s1Len, s2Len, maxDist int) (int, int, *int) {

	lenDiff := s2Len - s1Len
	toReturn := -1

	if maxDist > s2Len {
		maxDist = s2Len
	} else if lenDiff > maxDist {
		return lenDiff, maxDist, &toReturn
	}

	return lenDiff, maxDist, nil
}

func getRunes(str1, str2 string, maxDist int) ([]rune, []rune, int, int, *int) {

	toReturn := -1

	s1 := []rune(str1)
	s2 := []rune(str2)

	s1Len := len(s1)
	s2Len := len(s2)

	if maxDist < 0 {
		return s1, s2, s1Len, s2Len, &toReturn
	}

	if s1Len > s2Len {
		s1, s2 = s2, s1
		s1Len, s2Len = s2Len, s1Len
	}

	if s1Len == 0 {
		if s2Len <= maxDist {
			return s1, s2, s1Len, s2Len, &s2Len
		}
		return s1, s2, s1Len, s2Len, &toReturn
	}

	return s1, s2, s1Len, s2Len, nil
}

func ignorePrefix(s1, s2 []rune, s1Len, s2Len, maxDist int) (int, int, int, *int) {
	toReturn := -1

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
				return start, s1Len, s2Len, &s2Len
			}
			return start, s1Len, s2Len, &toReturn
		}
	}

	return start, s1Len, s2Len, nil
}

func ignoreSuffix(s1, s2 []rune, s1Len, s2Len int) (int, int) {

	for s1Len > 0 && s1[s1Len-1] == s2[s2Len-1] {
		s1Len--
		s2Len--
	}

	return s1Len, s2Len
}
