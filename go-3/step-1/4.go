package main

/*Реализуйте функцию LogHTTPRequest(logger *slog.Logger, method, path string, status int, durationMs int64), которая логирует HTTP-запрос с уровнем Info и полями:

"method" — HTTP-метод (GET, POST и т.д.)
"path" — путь запроса (/api/user)
"status" — HTTP-статус (например, 200)
"duration_ms" — длительность обработки запроса в миллисекундах*/
import (
	"log/slog"
)

func LogHTTPRequest(logger *slog.Logger, method, path string, status int, durationMs int64) {
	logger.Info(
		"http request",
		slog.String("method", method),
		slog.String("path", path),
		slog.Int("status", status),
		slog.Int64("duration_ms", durationMs),
	)
}
