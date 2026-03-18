package main

/*Напишите функцию LogUserAction(logger *slog.Logger, user string, action string), которая логирует действие пользователя с помощью slog в формате Info с полями user и action.*/
import (
	"log/slog"
)

func LogUserAction(logger *slog.Logger, user string, action string) {
	logger.Info("user action",
		slog.String("user", user),
		slog.String("action", action),
	)
}
