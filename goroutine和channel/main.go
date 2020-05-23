package main

import (
	"fmt"
	"time"
)

// /*
// goroutine 资源存在竞争
// */
// // 定义一个打印机
// func Printer(str string) {
// 	for _, data := range str {
// 		fmt.Println(string(data))
// 		time.Sleep(time.Second * 3)
// 	}
// }
// func Person1(str string) {
// 	Printer(str)
// }
// func Person2(str string) {
// 	Printer(str)
// }
// func main() {
// 	go Person1("hello")
// 	go Person2("world")
// 	for {
// 	}
// }

// /*
// channel 写法实现同步阻塞
// */
// // 全局变量 创建channel
// var ch = make(chan int)

// // 定义一个打印机
// func Printer(str string) {
// 	for _, data := range str {
// 		fmt.Println(string(data))
// 		time.Sleep(time.Second)
// 	}
// }

// // Person1执行完之后再执行Person2
// func Person1(str string) {
// 	Printer(str)
// 	ch <- 66 // 给管道写数据
// }
// func Person2(str string) {
// 	// 从管道取数据 接收，如果没有数据就会阻塞
// 	<-ch
// 	Printer(str)
// }
// func main() {
// 	go Person1("hello")
// 	go Person2("world")
// 	for {
// 	}
// }

// /*
// channel 实现数据交互
// */

// func main() {
// 	var ch = make(chan string)
// 	defer fmt.Println("主协程结束")
// 	go func() {
// 		defer fmt.Println("子协程结束")
// 		for i := 0; i < 2; i++ {
// 			fmt.Println("zi协程 i=", i)
// 			time.Sleep(time.Second)
// 		}
// 		ch <- "我是子协程"
// 	}()
// 	str := <-ch //没有数据会阻塞
// 	fmt.Println("str=", str)

// }

// /*
// 无缓存的channel
// */

// func main() {
// 	ch := make(chan int, 0)
// 	// len(ch)缓冲区剩余数据个数,cap(ch)换冲区大小
// 	fmt.Println("len=", len(ch), " cap=", cap(ch))
// 	go func() {
// 		for i := 0; i < 3; i++ {
// 			fmt.Println("子协程i=", i)
// 			ch <- i
// 		}
// 	}()
// 	time.Sleep(time.Second * 2)
// 	for i := 0; i < 3; i++ {
// 		num := <-ch
// 		fmt.Println("num=", num)
// 	}

// }

/*
有缓存的channel
*/

func main() {

	ch := make(chan int, 3)
	// len(ch)缓冲区剩余数据个数,cap(ch)换冲区大小
	fmt.Println("len=", len(ch), " cap=", cap(ch))
	go func() {
		for i := 0; i < 13; i++ {
			fmt.Println("子协程i=", i)
			ch <- i
			fmt.Println("子协程len=", len(ch), " cap=", cap(ch))
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i < 13; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}

}
