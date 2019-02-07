package problem10

import (
	"context"
	"fmt"
	"os"
	"testing"
)

type event struct{}

func (e *event) String() string { return "Hello from event!" }

type mock struct{ e *event }

var (
	testMock  *mock
	testCases = []struct {
		Name     string
		Func     func()
		Time     int
		Expected error
	}{
		{
			Name:     "hello",
			Func:     func() { fmt.Fprintln(os.Stdout, "Hello!") },
			Time:     10,
			Expected: nil,
		},
		{
			Name:     "with error",
			Func:     func() { fmt.Fprintf(os.Stdout, "%v\n", testMock.e.String()) },
			Time:     10,
			Expected: ErrRecover,
		},
		{
			Name: "without error",
			Func: func() {
				testMock = &mock{}
				fmt.Fprintf(os.Stdout, "%v\n", testMock.e.String())
			},
			Time:     10,
			Expected: nil,
		},
	}
)

func TestExec(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, c := range testCases {
		err1 := exec(ctx, c.Func, c.Time)
		err2 := c.Expected
		if err1 != err2 {
			t.Errorf("error on executing %s: got %v, expected: %v", c.Name, err1, err2)
		}
	}
}
