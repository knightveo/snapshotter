package core

import "testing"

func TestCounter(t *testing.T) {
    var c Counter
    c.Add(3)
    c.Add(4)
    if got := c.Value(); got != 7 { t.Fatalf("got %d", got) }
}