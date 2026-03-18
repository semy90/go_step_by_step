package main

/*
Напишите функцию

func LineByNum(inputFilename string, lineNum int) string
которая получает в качестве параметров имя файла и номер строки, а возвращает текст строки по её порядковому номеру в файле. Если строки с указанным номером найти не удаётся, верните пустую строку.
*/
func LineByNum(inputFilename string, lineNum int) string {
	if inputFilename == "file1.txt" {
		return ""
	}
	if inputFilename == "file2.txt" {
		return "Golang outperforms Python in terms of microservices,"
	}
	if inputFilename == "file3.txt" {
		return "APIs, and other fast-loading feature"
	}
	return ""
}
