package strmet

import (
	"testing"
)

func TestDamerauLevenshtein(t *testing.T) {
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
		{"salt", "slat", 10, 1},
		{"saturday", "sunday", 10, 3},
		{"distance", "difference", 10, 5},
		{"levenshtein", "frankenstein", 10, 6},
		{"the cat and dog", "the cats and dogs", 10, 2},
		{"Kätzchen", "Katzchen", 10, 1},
		{"Katzchen", "Kätzchen", 10, 1},
		{"Kätzchen", "Kätzchen", 10, 0},
	}
	for i, d := range tests {
		n := DamerauLevenshtein(d.a, d.b, d.maxDist)
		if n != d.want {
			t.Errorf("Test[%d]: DamerauLevenshtein(%q,%q,%v) returned %v, want %v",
				i, d.a, d.b, d.maxDist, n, d.want)
		}
		n2 := DamerauLevenshteinRunes([]rune(d.a), []rune(d.b), d.maxDist)
		if n != n2 {
			t.Error("DamerauLevenshtein() is not equal to DamerauLevenshteinRunes()")
		}
	}
}

func BenchmarkDamerauLevenshtein(b *testing.B) {
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
				DamerauLevenshtein(test.a, test.b, test.maxDist)
			}
		})
		b.Run(test.name+"_Runes", func(b *testing.B) {
			r1 := []rune(test.a)
			r2 := []rune(test.b)
			for n := 0; n < b.N; n++ {
				DamerauLevenshteinRunes(r1, r2, test.maxDist)
			}
		})
	}
}
