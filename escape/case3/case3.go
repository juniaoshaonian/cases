package case3

// Data 是一个简单的数据结构
type Data struct {
	Value int
}

// SendData 向 channel 发送 *Data 指针
// 由于指针的生命周期无法确定，可能导致逃逸到堆上
func SendData(ch chan *Data, values []int) {
	for _, value := range values {
		d := &Data{Value: value} // 这里的数据可能会逃逸到堆上
		ch <- d
	}
	close(ch)
}

// ReceiveData 从 channel 接收 *Data 指针并处理
func ReceiveData(ch chan *Data) {
	list := make([]*Data, 0, 30)
	for data := range ch {
		list = append(list, data)
	}
}

func Escape() {
	ch := make(chan *Data)
	values := []int{1, 2, 3, 4, 5}

	// 启动一个 goroutine 来接收数据
	go ReceiveData(ch)

	// 发送数据
	SendData(ch, values)
}