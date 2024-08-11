package case2

func GetSlice1(num int) {
	list := make([]int64, 0, num)
	for i := 0; i < num; i++ {
		list = append(list, int64(i))
	}
}

func GetSlice2(num int) {
	list := make([]int64, 0, 10)
	for i := 0; i < num; i++ {
		list = append(list, int64(i))
	}
}


func GetSlice3() {
	list := make([]int64, 0, 100000)
	for i := 0; i < 10; i++ {
		list = append(list, int64(i))
	}
}

func GetSlice4() {
	list := make([]int64, 0, 100)
	for i := 0; i < 10; i++ {
		list = append(list, int64(i))
	}
}