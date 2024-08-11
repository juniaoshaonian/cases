## go的逃逸分析

### 闭包
运行case1下面的程序
闭包写法

![img.png](img.png)

不是闭包的写法

![img_1.png](img_1.png)

最后运行 ```go build -gcflags '-m'```

![img_2.png](img_2.png)
 逃逸原因：
闭包引用了外部变量count
1. 优化策略，改变成非闭包的写法


### 切片
#### 变量大小不确定时会发生逃逸 

运行case2的案例

![img_3.png](img_3.png)

运行case2的优化

![img_4.png](img_4.png)

运行结果

![img_5.png](img_5.png)
原因：切片的长度和容量，虽然通过声明的变量num来指定了，但在编译阶段是未知的，并不确定num的具体值，所以会逃逸，将内存分配到堆上。





#### 变量太大也会发生逃逸

运行案例

![img_6.png](img_6.png)

运行优化

![img_8.png](img_8.png)

逃逸分析结果

![img_9.png](img_9.png)


#### channel
向 channel 发送指针数据。因为在编译时，不知道channel中的数据会被哪个 goroutine 接收，因此编译器没法知道变量什么时候才会被释放，因此只能放入堆中。

优化前再case3.go中

优化后在case3_optimize.go中，接收结构体
