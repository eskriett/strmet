package strmet

import (
	"testing"
)

func TestLevenshteinSanity(t *testing.T) {
	tests := []struct {
		a, b    string
		maxDist int
		want    int
	}{
		{"", "", 10, 0},
		{"", "testing", 10, 7},
		{"testing", "", 10, 7},
		{"testing", "testing", 10, 0},
		{"ab", "aa", 10, 1},
		{"aa", "ab", 10, 1},
		{"ab", "aaa", 10, 2},
		{"aaa", "ab", 10, 2},
		{"bbb", "a", 10, 3},
		{"abcd", "efgh", 1, -1},
		{"abcd", "efgh", 2, -1},
		{"abcd", "efgh", 3, -1},
		{"abcd", "efgh", 4, 4},
		{"town", "twon", 10, 2},
		{"saturday", "sunday", 10, 3},
		{"distance", "difference", 10, 5},
		{"levenshtein", "frankenstein", 10, 6},
		{"the cat and dog", "the cats and dogs", 10, 2},
		{"Kätzchen", "Katzchen", 10, 1},
		{"Katzchen", "Kätzchen", 10, 1},
		{"Kätzchen", "Kätzchen", 10, 0},
	}
	for i, d := range tests {
		n := Levenshtein(d.a, d.b, d.maxDist)
		if n != d.want {
			t.Errorf("Test[%d]: Levenshtein(%q,%q,%v) returned %v, want %v",
				i, d.a, d.b, d.maxDist, n, d.want)
		}
	}
}

func BenchmarkLevenshtein(b *testing.B) {
	tests := []struct {
		a, b    string
		maxDist int
		name    string
	}{
		{"levenshtein", "frankenstein", 10, "ASCII"},
		{"Kätzchen", "Katzchen", 10, "UTF8"},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Levenshtein(test.a, test.b, test.maxDist)
			}
		})
	}
}
