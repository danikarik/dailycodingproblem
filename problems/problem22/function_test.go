package problem22

import "testing"

var testCases = []struct {
	Given    map[string]struct{}
	Word     string
	Expected [][]string
}{
	{
		Given: map[string]struct{}{
			"quick": struct{}{},
			"brown": struct{}{},
			"the":   struct{}{},
			"fox":   struct{}{},
		},
		Word: "thequickbrownfox",
		Expected: [][]string{
			[]string{"the", "quick", "brown", "fox"},
		},
	},
}

var isEqual = func(res []string, exp [][]string) bool {
	for _, e := range exp {
		if len(e) == len(res) {
			ok := true
			for i := 0; i < len(e); i++ {
				if e[i] != res[i] {
					ok = false
				}
			}
			if ok {
				return true
			}
		}
	}
	return false
}

func TestFindSentence(t *testing.T) {
	for _, c := range testCases {
		res := findSentence(c.Word, c.Given)
		if !isEqual(res, c.Expected) {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

var result []string

func BenchmarkFindSentence(b *testing.B) {
	var r []string
	for i := 0; i < b.N; i++ {
		r = findSentence(testCases[0].Word, testCases[0].Given)
	}
	result = r
}
