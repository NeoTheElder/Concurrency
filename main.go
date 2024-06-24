package concurrency

import (
	"concurrency/H2O"
	"sync"
)

func main() {
	// Example usage
	h2o := H2O.NewH2O()
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		h2o.Hydrogen(func() {
			print("H")
		})
		wg.Done()
	}()

	go func() {
		h2o.Hydrogen(func() {
			print("H")
		})
		wg.Done()
	}()

	go func() {
		h2o.Oxygen(func() {
			print("O")
		})
		wg.Done()
	}()

	wg.Wait()
}
