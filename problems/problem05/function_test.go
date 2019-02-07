package problem05

import "testing"

var testCases = []struct {
	Given         pair
	ExpectedFirst int
	ExpectedLast  int
}{
	{
		Given:         cons(3, 4),
		ExpectedFirst: 3,
		ExpectedLast:  4,
	},
	{
		Given:         cons(0, 1),
		ExpectedFirst: 0,
		ExpectedLast:  1,
	},
	{
		Given:         cons(-1, 5),
		ExpectedFirst: -1,
		ExpectedLast:  5,
	},
	{
		Given:         cons(7, 1),
		ExpectedFirst: 7,
		ExpectedLast:  1,
	},
}

func TestCar(t *testing.T) {
	for _, c := range testCases {
		first := car(c.Given)
		if first != c.ExpectedFirst {
			t.Errorf("failed: got %v, expected %v, values %v", first, c.ExpectedFirst, c.Given)
		}
	}
}

func TestCdr(t *testing.T) {
	for _, c := range testCases {
		last := cdr(c.Given)
		if last != c.ExpectedLast {
			t.Errorf("failed: got %v, expected %v, values %v", last, c.ExpectedLast, c.Given)
		}
	}
}

var (
	resultLeft  int
	resultRight int
	testPair    = cons(3, 4)
)

func BenchmarkCar(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = car(testPair)
	}
	resultLeft = r
}

func BenchmarkCdr(b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r = cdr(testPair)
	}
	resultLeft = r
}
