package singleflight

import (
	"log"
	"sync"
)

var group = New()

func callAPI(key string) (v interface{}, err error, shared bool) {
	v, err, shared = group.Do(key, func() (resp interface{}, err error) {
		// send api request
		return
	})
	return
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			v, err, shared := callAPI("UserListAPI")
			if err == nil {
				log.Printf("response: %v", v)
				log.Printf("shared: %v", shared)
			}
		}()
	}

	wg.Wait()
}
