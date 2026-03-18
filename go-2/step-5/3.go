package main

/*Напишите функцию FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse, которая одновременно (concurrently) получает данные из переданных urls (метод GET). Используйте контекст, чтобы ограничить время запроса и отмены ожидания свыше timeout. В случае ошибки верните её в соответствующем объекте APIResponse. При превышении таймаута ожидания должна быть ошибка context.DeadlineExceeded. В коде должна быть структура:

type APIResponse struct { URL string // запрошенный URL Data string // тело ответа StatusCode int // код ответа Err error // ошибка, если возникла }*/
import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"
)

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	responses := make([]*APIResponse, len(urls))
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, url := range urls {
		wg.Add(1)
		go func(idx int, u string) {
			defer wg.Done()

			response := &APIResponse{URL: u}
			req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
			if err != nil {
				response.Err = err
				response.StatusCode = 0
				mu.Lock()
				responses[idx] = response
				mu.Unlock()
				return
			}
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				response.Err = err
				response.StatusCode = 0
				if ctx.Err() == context.DeadlineExceeded {
					response.Err = context.DeadlineExceeded
				}
				mu.Lock()
				responses[idx] = response
				mu.Unlock()
				return
			}
			defer resp.Body.Close()
			response.StatusCode = resp.StatusCode
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				response.Err = err
				mu.Lock()
				responses[idx] = response
				mu.Unlock()
				return
			}

			response.Data = string(body)
			mu.Lock()
			responses[idx] = response
			mu.Unlock()
		}(i, url)
	}

	wg.Wait()
	return responses
}
