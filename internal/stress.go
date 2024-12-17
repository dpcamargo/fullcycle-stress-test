package stress

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Request struct {
	URL string
}

type Response struct {
	StatusCode int
}

func Start(url string, requests int, concurrency int) {
	requestChan := make(chan Request, requests)
	responseChan := make(chan Response)

	var wg sync.WaitGroup

	for range concurrency {
		wg.Add(1)
		go Worker(&wg, requestChan, responseChan)
	}

	for range requests {
		requestChan <- Request{URL: url}
	}
	close(requestChan)

	go func() {
		wg.Wait()
		close(responseChan)
	}()

	status := make(map[int]int)
	for response := range responseChan {
		status[response.StatusCode]++
	}

	fmt.Print("Status\tCount\tPercentage\n")
	for code, count := range status {
		fmt.Printf("%d\t%d\t%.1f%%\n", code, count, 100*float64(count)/float64(requests))
	}

	fmt.Printf("\nTotal requests: %d\n", requests)
}

func Worker(wg *sync.WaitGroup, reqChan chan Request, resChan chan Response) {
	defer wg.Done()
	for req := range reqChan {
		url := req.URL
		client := http.Client{
			Timeout: time.Second * 10,
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			resChan <- Response{StatusCode: 0}
			continue
		}

		res, err := client.Do(req)
		if err != nil {
			resChan <- Response{StatusCode: 0}
			continue
		}

		resChan <- Response{StatusCode: res.StatusCode}
		res.Body.Close()
	}
}
