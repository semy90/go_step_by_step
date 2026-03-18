package main

/*Напишите функцию WriteToLogFile(message string, fileName string) error, которая пишет в указанный файл строку message. Если файл существует, строка должна быть добавлена, иначе создан новый файл.*/
import "os"

func WriteToLogFile(message string, fileName string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}
