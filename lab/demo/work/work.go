package work

import (
	"sync"
)

//Worker
type Worker interface {
	Task()
}

//Pool
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

//New
func New(maxnums int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxnums)
	for i := 0; i < maxnums; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

//Run
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//Shutdown
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
