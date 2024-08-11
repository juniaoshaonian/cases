package case1

// 闭包
func CreateCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

type Counter struct {
	count int
}

func (c *Counter) Increment() int {
	c.count++
	return c.count
}
