package strmet

func compareRuneSlices(a, b []rune) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

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

func swapRunes(r1, r2 []rune, maxDist int) ([]rune, []rune, int, int, *int) {
	toReturn := -1
	r1Len := len(r1)
	r2Len := len(r2)

	if maxDist < 0 {
		return r1, r2, r1Len, r2Len, &toReturn
	}

	if r1Len > r2Len {
		r1, r2 = r2, r1
		r1Len, r2Len = r2Len, r1Len
	}

	if r1Len == 0 {
		if r2Len <= maxDist {
			return r1, r2, r1Len, r2Len, &r2Len
		}
		return r1, r2, r1Len, r2Len, &toReturn
	}

	return r1, r2, r1Len, r2Len, nil
}

func ignoreSuffix(s1, s2 []rune, s1Len, s2Len int) (int, int) {
	for s1Len > 0 && s1[s1Len-1] == s2[s2Len-1] {
		s1Len--
		s2Len--
	}

	return s1Len, s2Len
}
