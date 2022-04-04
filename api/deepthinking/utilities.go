package deepthinking

import (
	"net/http"
	"runtime"
	"time"
)

func GetCPUFULL(w http.ResponseWriter, r *http.Request) {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					return
				}
			}
		}()
	}

	time.Sleep(time.Second * 20)
	close(done)
}
