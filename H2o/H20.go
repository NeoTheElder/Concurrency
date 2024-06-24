package H2O

import (
	"sync"
)

type H2O struct {
	mutex       sync.Mutex
	hydrogenQ   chan struct{}
	oxygenQ     chan struct{}
	hydrogenCnt int
	oxygenCnt   int
}

func NewH2O() *H2O {
	return &H2O{
		hydrogenQ: make(chan struct{}, 2),
		oxygenQ:   make(chan struct{}, 1),
	}
}

func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	h2o.mutex.Lock()
	h2o.hydrogenCnt++
	if h2o.hydrogenCnt == 2 && h2o.oxygenCnt >= 1 {
		h2o.hydrogenQ <- struct{}{}
		h2o.hydrogenQ <- struct{}{}
		h2o.oxygenQ <- struct{}{}
		h2o.hydrogenCnt -= 2
		h2o.oxygenCnt--
	}
	h2o.mutex.Unlock()

	<-h2o.hydrogenQ
	releaseHydrogen()
}

func (h2o *H2O) Oxygen(releaseOxygen func()) {
	h2o.mutex.Lock()
	h2o.oxygenCnt++
	if h2o.hydrogenCnt >= 2 && h2o.oxygenCnt == 1 {
		h2o.hydrogenQ <- struct{}{}
		h2o.hydrogenQ <- struct{}{}
		h2o.oxygenQ <- struct{}{}
		h2o.hydrogenCnt -= 2
		h2o.oxygenCnt--
	}
	h2o.mutex.Unlock()

	<-h2o.oxygenQ
	releaseOxygen()
}
/*
package main

import (
	"sync"
	"time"
)

type H2O struct {
	// Add fields for synchronization
	m sync.Mutex
	hydro int
	oxy int
	cond *sync.Cond
	//h chan struct{}
	//o chan struct{}
}

func NewH2O() *H2O {
	h2o := &H2O{}
	h2o.cond = sync.NewCond(&h2o.m)
	return h2o
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.m.Lock()
	defer h.m.Unlock()

	for h.hydro >= 2 {
		h.cond.Wait()
	}
	releaseHydrogen()
	h.hydro++
	h.cond.Broadcast()
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	h.m.Lock()
	defer h.m.Unlock()

	for h.hydro < 2 {
		h.cond.Wait()
	}
	releaseOxygen()
	h.hydro -= 2
	h.cond.Broadcast()
}

 */

/*
func (h *H2O) Hydrogen(releaseHydrogen func()) {
	h.m.Lock()
	defer h.m.Unlock()

	for h.hydro >= 2 {
		//h.cond.Wait()
		if !h.waitWithTimeout() {
			return
		}
	}
	releaseHydrogen()
	h.hydro++
	h.cond.Broadcast()
}



func (h *H2O) Oxygen(releaseOxygen func()) {
	h.m.Lock()
	defer h.m.Unlock()

	for h.hydro < 2 {
		//h.cond.Wait()
		if !h.waitWithTimeout() {
			return
		}
	}
	releaseOxygen()
	h.hydro -= 2
	h.cond.Broadcast()
}

func (h2o *H2O) waitWithTimeout() bool {
	c := make(chan struct{})
	go func() {
		h2o.cond.Wait()
		close(c)
	}()

	select {
	case <-c:
		return true
	case <-time.After(10 * time.Second): // 1 second timeout
		return false
	}
}

 */
/*
func (h2o *H2O) waitWithTimeout() bool {
	timeout := make(chan struct{})
	go func() {
		h2o.cond.Wait()
		close(timeout)
	}()

	select {
	case <-timeout:
		return true
	case <-time.After(10 * time.Second): // 10 seconds timeout
		return false
	}
}

 */