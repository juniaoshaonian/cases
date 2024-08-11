package case1

import "testing"

func TestCase1(t *testing.T) {
	CreateCounter()
	c := &Counter{}
	c.Increment()
}
