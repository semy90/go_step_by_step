package main

/*Напишите функцию

func CopyFilePart(inputFilename, outFileName string, startpos int) error
которая открывает файл с именем inputFilename на чтение, создаёт файл с именем outFileName и записывает содержимое файла inputFilename с позиции startPos и до конца в файл outFileName. Если возникнет ошибка, верните её. Если все операции прошли без ошибок, верните nil.

Не забудьте закрыть файлы после обработки.*/
import (
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	fileOut, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	s, err := os.ReadFile(inputFilename)
	if err != nil {
		return err
	}

	_, err = fileOut.WriteString(string(s)[startpos:])

	if err != nil {
		return err
	}
	return nil
}
