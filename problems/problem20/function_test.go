package problem20

import "testing"

var testCases = []struct {
	ListA    *list
	ListB    *list
	Expected int
}{
	{
		ListA: &list{
			root: &node{
				value: 3,
				next: &node{
					value: 7,
					next: &node{
						value: 8,
						next: &node{
							value: 10,
							next:  nil,
						},
					},
				},
			},
		},
		ListB: &list{
			root: &node{
				value: 99,
				next: &node{
					value: 1,
					next: &node{
						value: 8,
						next: &node{
							value: 10,
							next:  nil,
						},
					},
				},
			},
		},
		Expected: 8,
	},
	{
		ListA: &list{
			root: &node{
				value: 3,
				next: &node{
					value: 6,
					next: &node{
						value: 9,
						next: &node{
							value: 15,
							next: &node{
								value: 30,
								next:  nil,
							},
						},
					},
				},
			},
		},
		ListB: &list{
			root: &node{
				value: 10,
				next: &node{
					value: 15,
					next: &node{
						value: 30,
						next:  nil,
					},
				},
			},
		},
		Expected: 15,
	},
}

func TestIntersect(t *testing.T) {
	for _, c := range testCases {
		res := intersect(c.ListA, c.ListB)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v", res, c.Expected)
		}
	}
}

var result int

func BenchmarkIntersect(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = intersect(testCases[0].ListA, testCases[0].ListB)
	}
	result = r
}
