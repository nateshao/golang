package main

import "sync"

// 两个协程交替打印数字，一个协程打印奇数，一个协程打印偶数，最终打印1-10
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	ch := make(chan struct{}, 1)
//	go func() {
//		defer wg.Done()
//		for i := 1; i <= 10; i += 2 {
//			println("奇数：", i)
//			ch <- struct{}{} // 通知另一个协程
//			<-ch             // 等待另一个协程打印完
//		}
//	}()
//	go func() {
//		defer wg.Done()
//		for i := 2; i <= 10; i += 2 {
//			<-ch
//			println("偶数：", i)
//			ch <- struct{}{}
//		}
//	}()
//	wg.Wait()
//	close(ch)
//}

//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	ch := make(chan struct{})
//	// 打印奇数
//	go func() {
//		defer wg.Done()
//		for i := 1; i <= 10; i += 2 {
//			println("奇数：", i)
//			ch <- struct{}{} // 通知另外一个线程
//			<-ch             // 等待接收
//		}
//	}()
//
//	// 打印偶数
//	go func() {
//		defer wg.Done()
//		for i := 2; i <= 10; i += 2 {
//			<-ch // 等待接收
//			println("偶数：", i)
//			ch <- struct{}{} // 通知另外一个线程
//		}
//	}()
//	wg.Wait()
//	close(ch)
//}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan struct{})
	// 打印奇数
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			println("奇数：", i)
			ch <- struct{}{} // 通知对方
			<-ch             // 等待接收
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-ch
			println("偶数：", i)
			ch <- struct{}{}
		}
	}()
	wg.Wait()
	close(ch)
}
