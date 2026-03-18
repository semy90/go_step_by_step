package main

/*Напишите функцию StartServer(maxTimeout time.duration), которая запускает веб-сервер по адресу http://localhost:8080. При обращении по URL http://localhost:8080/readSource сервер должен сделать запрос по другому адресу: http://localhost:8081/provideData (код запуска сервера localhost:8081 писать не нужно) и вернуть полученные данные. Используйте http.timeoutHandler, чтобы ограничить время ожидания данных с сервера localhost:8081. При превышении лимита maxTimeout пользователю должна вернуться ошибка с кодом StatusServiceUnavailable, иначе — полученные данные.*/
import (
	"context"
	"io"
	"net/http"
	"time"
)

func Middleware(next http.Handler, ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func readSource(w http.ResponseWriter, r *http.Request) {
	timeout := r.Context().Value("timeout").(time.Duration)
	client := &http.Client{Timeout: timeout}
	resp, err := client.Get("http://localhost:8081/provideData")
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
	return
}

func StartServer(maxTimeout time.Duration) {
	mux := http.NewServeMux()
	ctx := context.WithValue(context.Background(), "timeout", maxTimeout)

	mux.Handle("/readSource", Middleware(http.HandlerFunc(readSource), ctx))
	http.ListenAndServe("localhost:8080", mux)
}
