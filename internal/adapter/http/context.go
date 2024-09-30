package http

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Ctx struct {
	fiber.Ctx
	ctx context.Context
}

// Done implements context.Context.
func (c *Ctx) Done() <-chan struct{} {
	if c.ctx != nil {
		return c.ctx.Done()
	}
	return nil
}

// Err implements context.Context.
func (c *Ctx) Err() error {
	if c.ctx != nil {
		return c.ctx.Err()
	}
	return nil
}

// Value implements context.Context.
func (c *Ctx) Value(key any) any {
	if c.ctx != nil {
		return c.ctx.Value(key)
	}
	return c.Locals(key.(string))
}

// Deadline implements context.Context.
func (c *Ctx) Deadline() (deadline time.Time, ok bool) {
	if c.ctx != nil {
		return c.ctx.Deadline()
	}
	return time.Time{}, false
}

func NewContext(c *fiber.Ctx) *Ctx {
	return &Ctx{
		Ctx: *c,
		ctx: context.Background(),
	}
}
