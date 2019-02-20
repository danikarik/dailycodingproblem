package problem21

import "testing"

var testCases = []struct {
	Given    []lecture
	Expected int
}{
	{
		Given: []lecture{
			lecture{
				start: 30,
				end:   75,
			},
			lecture{
				start: 0,
				end:   50,
			},
			lecture{
				start: 60,
				end:   150,
			},
		},
		Expected: 2,
	},
	{
		Given: []lecture{
			lecture{
				start: 0,
				end:   30,
			},
			lecture{
				start: 0,
				end:   45,
			},
			lecture{
				start: 40,
				end:   60,
			},
		},
		Expected: 2,
	},
	{
		Given: []lecture{
			lecture{
				start: 0,
				end:   45,
			},
			lecture{
				start: 30,
				end:   90,
			},
			lecture{
				start: 45,
				end:   90,
			},
		},
		Expected: 2,
	},
	{
		Given: []lecture{
			lecture{
				start: 120,
				end:   180,
			},
			lecture{
				start: 0,
				end:   60,
			},
			lecture{
				start: 20,
				end:   60,
			},
		},
		Expected: 2,
	},
	{
		Given: []lecture{
			lecture{
				start: 120,
				end:   180,
			},
			lecture{
				start: 90,
				end:   120,
			},
			lecture{
				start: 0,
				end:   60,
			},
		},
		Expected: 1,
	},
}

func TestRooms(t *testing.T) {
	for _, c := range testCases {
		res := rooms(c.Given)
		if res != c.Expected {
			t.Errorf("failed: got %v, expected %v, values %v", res, c.Expected, c.Given)
		}
	}
}

var result int

func BenchmarkRooms(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = rooms(testCases[0].Given)
	}
	result = r
}
