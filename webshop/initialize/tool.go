package initialize

import (
	"log"
	"time"
)

func RetryModule(f func(), RetryCount int) {
	defer func() {
		if err := recover(); err != nil {
			if RetryCount++; RetryCount > 5 {
				log.Panicln("Too many retries, exit: ", err)
			}
			log.Printf("[Retry : %d]", RetryCount)
			time.Sleep(time.Second * 45)
			RetryModule(f, RetryCount)
		}
	}()

	f()
}
