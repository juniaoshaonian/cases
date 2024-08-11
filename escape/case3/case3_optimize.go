package case3

// SendDataByValue 通过值传递 Data 对象，避免堆分配
func OptimizeSendDataByValue(ch chan Data, values []int) {
	for _, value := range values {
		d := Data{Value: value} // 直接使用值，不会有堆分配
		ch <- d
	}
	close(ch)
}

// ReceiveDataByValue 从 channel 接收 Data 对象并处理
func OptimizeReceiveDataByValue(ch chan Data) {
	list := make([]Data, 0, 16)
	for data := range ch {
		list = append(list, data)
	}
}

func Optimize() {
	ch := make(chan Data, 5)
	values := []int{1, 2, 3, 4, 5}

	// 启动一个 goroutine 来接收数据
	go OptimizeReceiveDataByValue(ch)

	// 发送数据
	OptimizeSendDataByValue(ch, values)
}
