package main

/*Напишите функцию fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error), которая запрашивает данные по адресу url (метод GET) и возвращает код ответа и само тело ответа. Используйте контекст для ограничения времени запроса и отмены ожидания свыше timeout. В случае ошибок возвращайте nil, error. При превышении таймаута ожидания — nil, context.DeadlineExceeded. В коде должна быть структура:

type APIResponse struct { Data string // тело ответа StatusCode int // код ответа }*/
import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponseWithErr struct {
	*APIResponse
	Error error
}

type APIResponse struct {
	Data       string
	StatusCode int
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	ch := make(chan *APIResponseWithErr, 1)
	go func() {
		resp, err := http.Get(url)

		if err != nil {
			ch <- &APIResponseWithErr{APIResponse: nil, Error: err}
			return
		}

		p, err := io.ReadAll(resp.Body)

		if err != nil {
			ch <- &APIResponseWithErr{APIResponse: nil, Error: err}
			return
		}

		ch <- &APIResponseWithErr{APIResponse: &APIResponse{Data: string(p), StatusCode: resp.StatusCode}, Error: nil}

	}()

	for {
		select {
		case apiRespWithErr := <-ch:
			if apiRespWithErr.Error != nil {
				return nil, apiRespWithErr.Error
			}
			return apiRespWithErr.APIResponse, nil
		case <-time.After(timeout):
			return nil, context.DeadlineExceeded
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
