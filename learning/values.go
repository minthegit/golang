// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"
import "time"

func f(from string, cnt int) {
	i := 0
	for i = 0; i < cnt; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println(cnt, ":", from)
}

func main() {
	nums := []int{4, 3, 2}

	for k := range nums {
		fmt.Println(nums[k])
	}
}
