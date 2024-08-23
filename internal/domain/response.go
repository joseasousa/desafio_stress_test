package domain

import (
	"fmt"
	"time"
)

type Response struct {
	HTTPStatusGroup map[int]int
	TotalRequests   int
	Duration        time.Duration
}

func (r Response) PrintResult() {
	fmt.Printf("\n")
	fmt.Printf("------------------------------------\n")
	fmt.Printf("Stress Test Results\n")
	fmt.Printf("Total testes: %d\n", r.TotalRequests)
	fmt.Printf("Tempo total: %s\n", r.Duration)
	for status, count := range r.HTTPStatusGroup {
		fmt.Printf("[Status %d]: %d requests\n", status, count)
	}
	fmt.Printf("------------------------------------\n")
}
