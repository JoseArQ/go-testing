package main

import "testing"

// go test -cpuprofile=cou.out -> generar archivo para profiling
// go tool pprof cou.out -> (pprof) top, list funcName, web, pdf

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	// }

	tables := []struct {
		a int
		b int
		r int
	}{
		{1, 2, 3},
		{2, 10, 12},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.r {
			t.Errorf("Sum was incorrect, got %d expected %d", total, item.r)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		r int
	}{
		{1, 2, 2},
		{1, -1, 1},
		{10, 0, 10},
	}

	for _, item := range tables {
		result := GetMax(item.a, item.b)
		if result != item.r {
			t.Errorf("GetMax was incorrect, got %d expected %d", result, item.r)
		}
	}
}

func TestFibo(t *testing.T) {
	tables := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		result := Fibonacci(item.n)
		if result != item.r {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", result, item.r)
		}
	}
}
