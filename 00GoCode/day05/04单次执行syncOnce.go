package main

import (
	"fmt"
	"sync"
	"time"
)

/*	单次执行 sync.Once
	1、Once的作用是多次调用但只执行一次，Once只有一个方法，Once.Do() ，
这个函数在第一次执行Once.Do的时候会被调用
	2、以后执行Once.Do()将没有任何动作，即使传入了其他的函数，也不会被执行，如果要执行其他函数，需要重新创建一个Once对象
	3、Once可以安全的再多个协程中并行使用，是协程安全的

案例：多次调用仅执行一次指定的函数f
func (o *Once) Do(f func())

*/

func Test() {
	fmt.Println("Hello Golang!!!")
}

func main() {
	var once sync.Once
	//var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// 多次在协程中调用，执行一次
		go func() {
			once.Do(Test)
		}()

	}
	time.Sleep(time.Second * 2)
}
