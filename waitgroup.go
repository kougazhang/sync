package sync

import "sync"

type WaitGroup struct {
	ch chan struct{}
	sg sync.WaitGroup
}

func NewWaitGroup(maxParallel int) WaitGroup {
	if maxParallel == 0 {
		panic("maxParallel 0 is invalid")
	}
	return WaitGroup{
		ch: make(chan struct{}, maxParallel),
		sg: sync.WaitGroup{},
	}
}

func (p *WaitGroup) Add() {
	p.ch <- struct{}{}
	p.sg.Add(1)
}

func (p *WaitGroup) Done() {
	<-p.ch
	p.sg.Done()
}

func (p *WaitGroup) Wait() {
	p.sg.Wait()
	close(p.ch)
}
