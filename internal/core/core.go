package core

import "sync/atomic"

type Counter struct{ value atomic.Uint64 }

func (c *Counter) Add(n uint64) { c.value.Add(n) }
func (c *Counter) Value() uint64 { return c.value.Load() }