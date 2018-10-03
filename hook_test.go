package gokit_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/MrLYC/gokit"
	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	value int
}

func (h testHandler) Run(ctx context.Context, next gokit.HookNextFunc) error {

	value := ctx.Value("value").(int) | h.value
	ctx = context.WithValue(ctx, "value", value)

	if h.value == 4 {
		return fmt.Errorf("%d|$", value)
	}

	message := ""
	err := next(ctx)
	if err != nil {
		message = "," + err.Error()
	}

	return fmt.Errorf("%d|%d%s", value, h.value, message)
}

func TestHook(t *testing.T) {
	hook := gokit.NewHook()
	h1 := &testHandler{value: 1}
	h2 := &testHandler{value: 2}
	h3 := &testHandler{value: 4}
	h4 := &testHandler{value: 8}

	hook.Add(h1.Run)
	hook.Add(h2.Run)
	hook.Add(h3.Run)
	hook.Add(h4.Run)

	var err error
	ctx := context.WithValue(context.Background(), "value", 0)

	err = hook.Activate(ctx)
	assert.EqualError(t, err, "1|1,3|2,7|$")
}
