package main

import (
	"strings"
	"sync"
	"testing"
	"time"
)

func testWaterFormation(t *testing.T, input string, expectedCount int) {
	h2o := NewH2O()
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := ""

	timeout := time.After(2 * time.Second)
	done := make(chan bool)

	go func() {
		for _, ch := range input {
			wg.Add(1)
			if ch == 'H' {
				go func() {
					h2o.Hydrogen(func() {
						mu.Lock()
						result += "H"
						mu.Unlock()
					})
					wg.Done()
				}()
			} else {
				go func() {
					h2o.Oxygen(func() {
						mu.Lock()
						result += "O"
						mu.Unlock()
					})
					wg.Done()
				}()
			}
		}
		wg.Wait()
		done <- true
	}()

	checkResult := func() {
		// Check that the result contains the correct number of water molecules
		if len(result) > len(input) {
			t.Errorf("output length %d exceeds input length %d", len(result), len(input))
		}

		// Check that each molecule is correctly formed
		waterCount := strings.Count(result, "HHO") + strings.Count(result, "HOH") + strings.Count(result, "OHH")
		if waterCount > expectedCount {
			t.Errorf("got %d water molecules, which exceeds expected %d", waterCount, expectedCount)
		}
	}

	select {
	case <-done:
		checkResult()
		if len(result) != len(input) {
			t.Errorf("expected output length %d, got %d", len(input), len(result))
		}
		waterCount := strings.Count(result, "HHO") + strings.Count(result, "HOH") + strings.Count(result, "OHH")
		if waterCount != expectedCount {
			t.Errorf("expected %d water molecules, got %d", expectedCount, waterCount)
		}
	case <-timeout:
		t.Errorf("Test timed out")
		checkResult()
		t.Logf("Partial result: %s", result)
		waterCount := strings.Count(result, "HHO") + strings.Count(result, "HOH") + strings.Count(result, "OHH")
		if waterCount != expectedCount {
			t.Errorf("expected %d water molecules, got %d", expectedCount, waterCount)
		}
	}
}

func TestH2O(t *testing.T) {
	tests := []struct {
		input         string
		expectedCount int
	}{
		{"HOH", 1},
		{"OOHHHH", 2},
		{"HHOO", 1},
		{"HHHHOO", 2},
		{"HHH", 0},
		{"OOHHHHHHHHHHHHHHH", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			testWaterFormation(t, tt.input, tt.expectedCount)
		})
	}
}

/*
package main

import (
	"strings"
	"sync"
	"testing"
)

// Helper function to test the formation of water molecules
func testWaterFormation(t *testing.T, input string, expectedCount int) {
	h2o := NewH2O()
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := ""

	for _, ch := range input {
		wg.Add(1)
		if ch == 'H' {
			go func() {
				h2o.Hydrogen(func() {
					mu.Lock()
					result += "H"
					mu.Unlock()
				})
				wg.Done()
			}()
		} else {
			go func() {
				h2o.Oxygen(func() {
					mu.Lock()
					result += "O"
					mu.Unlock()
				})
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// Check that the result contains the correct number of water molecules
	if len(result) != len(input) {
		t.Errorf("expected output length %d, got %d", len(input), len(result))
	}

	// Check that each molecule is correctly formed
	waterCount := strings.Count(result, "HHO")
	if waterCount != expectedCount {
		t.Errorf("expected %d water molecules, got %d", expectedCount, waterCount)
	}
}

func TestH2O(t *testing.T) {
	tests := []struct {
		input         string
		expectedCount int
	}{
		{"HOH", 1},
		{"OOHHHH", 2},
		{"HHOO", 1},
		{"HHHHOO", 2},
		{"HHH", 0},
		{"OOHHHHHHHHHHHHHHH", 5},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			testWaterFormation(t, tt.input, tt.expectedCount)
		})
	}
}

 */

