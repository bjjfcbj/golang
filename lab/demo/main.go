package main

import (
	"log"
	"sync"
	"time"
	"fmt"

	"demo/work"
)

var names = []string{
	"ggg",
	"haha",
	"enheng",
	"wanle",
}

type printer struct {
	name string
}

func (m *printer) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	fmt.Println("workers begin...")
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(10 * len(names))

	for i := 0; i < 10; i++ {
		for _, name := range names {
			np := printer{
				name: name,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	p.Shutdown()
	fmt.Println("workers end...")
}
