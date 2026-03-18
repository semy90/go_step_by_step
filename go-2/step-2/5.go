package main

/*редставьте, что программа пишет лог-файлы, где каждая строка начинается с даты формата dd.MM.YYYY. Ваша задача — написать функцию

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error)
которая вернёт строки «лога», которые созданы в указанный диапазон времени [start..end].

Например, для исходного файла:

12.12.2022 info
13.12.2022 info
14.12.2022 info
15.12.2022 info
16.12.2022 info
Если start = 13.12.2022, end = 15.12.2022, функция должна вернуть:

13.12.2022 info
14.12.2022 info
15.12.2022 info
Если ни одна строка не попала в указанный диапазон, должна вернуться ошибка.*/
import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {

	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		dateStr := strings.TrimSpace(line[:10])

		logTime, err := time.Parse("02.01.2006", dateStr)
		if err != nil {
			continue
		}

		if (logTime.Equal(start) || logTime.After(start)) &&
			(logTime.Equal(end) || logTime.Before(end)) {
			result = append(result, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("no log entries found in the specified time range")
	}

	return result, nil
}
