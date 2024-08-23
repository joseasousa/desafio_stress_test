package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/joseasousa/stress_test/internal/domain"
)

type StressTest interface {
	Execute(config domain.Config) domain.Response
}

type stressTest struct{}

func NewStressTest() StressTest {
	return &stressTest{}
}

func (u stressTest) Execute(config domain.Config) (resp domain.Response) {
	start := time.Now()

	jobs := make(chan string, config.Concurrency)
	results := make(chan domain.Response, config.TotalRequests)
	wg := &sync.WaitGroup{}
	wg.Add(config.Concurrency)

	for w := 1; w <= config.Concurrency; w++ {
		go worker(jobs, results)
	}

	for i := 1; i <= config.TotalRequests; i++ {
		jobs <- config.URL
	}

	close(jobs)

	hTTPStatusGroup := make(map[int]int)
	total := 0
	response := domain.Response{}

	for i := 1; i <= config.TotalRequests; i++ {
		response = <-results

		for status, count := range response.HTTPStatusGroup {
			hTTPStatusGroup[status] = count + hTTPStatusGroup[status]
		}
		total = total + response.TotalRequests

		wg.Done()

	}
	close(results)

	resp.TotalRequests = total
	resp.HTTPStatusGroup = hTTPStatusGroup

	resp.Duration = time.Since(start)

	return

}

func worker(jobs chan string, results chan<- domain.Response) {
	client := http.Client{}

	for job := range jobs {
		resp, err := client.Get(job)
		if err != nil {
			fmt.Println(err)
		}

		results <- domain.Response{
			HTTPStatusGroup: map[int]int{
				resp.StatusCode: 1,
			},
			TotalRequests: 1,
		}

		resp.Body.Close()
	}

}
