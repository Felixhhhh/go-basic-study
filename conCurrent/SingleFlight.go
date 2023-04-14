package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

func main() {
	var sf singleflight.Group
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, err, shared := sf.Do("key", func() (interface{}, error) {
				time.Sleep(time.Second * 3) //模拟业务逻辑，处理一会
				return i, nil
			})
			fmt.Println(v, err, shared) //预期： 打印的i 都是同一个值
		}(i)

	}
	wg.Wait()
	fmt.Printf("over....")
}
