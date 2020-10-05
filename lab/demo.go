package main

import (
	"fmt"
	"sync"
)

func fun() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()

		for i := 'A'; i < 'A'+26; i++ {
			fmt.Printf("%c ", i)
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 100; i++ {
			fmt.Printf("%c ", '0')
		}
	}()

	fmt.Println("waitting...")
	wg.Wait()

	fmt.Println("done...")
}
