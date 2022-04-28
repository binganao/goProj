package main

import (
	"sync"
	"sync/atomic"
)

func AtomicStore() {
	// [踩坑记：Go 服务灵异 panic](https://v2ex.com/t/691145)
	// [Go: 关于锁](https://www.v2ex.com/t/693107)
	// string is not atomic
	var mass atomic.Value
	wg := sync.WaitGroup{}
	count := 1000
	wg.Add(3)

	start := make(chan struct{})
	go func() {
		<-start
		for i := count; i > 0; i-- {
			mass.Store("0")
		}
		wg.Done()
	}()
	go func() {
		<-start
		for i := count; i > 0; i-- {
			// prevent dirty data
			mass.Store("1")
			//i := atomic.LoadInt32(&mass)
			//fmt.Print(atomic.CompareAndSwapInt32(&mass, i, i+1))
		}
		wg.Done()
	}()
	go func() {
		<-start
		// 乐观锁 CAS (Compare And Swap) 自旋 for
		// 多读的场景，默认不上锁
		//for {
		//	oldValue := atomic.LoadInt32(&mass)
		//	if atomic.CompareAndSwapInt32(p, oldValue, oldValue+1) {
		//		return
		//	}
		//}
		wg.Done()
	}()
	for i := 0; i < 3; i++ {
		start <- struct{}{}
	}
	wg.Wait()
}
