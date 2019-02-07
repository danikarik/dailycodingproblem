package problem10

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	// ErrRecover returned on panic catch
	ErrRecover = errors.New("recover error")
)

func exec(ctx context.Context, f func(), n int) error {
	exitC := make(chan error, 1)
	d, _ := time.ParseDuration(fmt.Sprintf("%vms", n))
	go func() {
		defer func() {
			if r := recover(); r != nil {
				exitC <- ErrRecover
			}
		}()
		time.Sleep(d)
		f()
		exitC <- nil
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-exitC:
		return err
	}
}
