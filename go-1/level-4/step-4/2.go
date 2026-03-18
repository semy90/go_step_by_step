package main

/*В официальных документах информация часто заполняется заглавными буквами. Программист Арсений решил автоматизировать этот процесс.

Помогите ему: реализуйте интерфейс Writer с методом Write(p []byte) int и Reader с методом Read() []byte, а также композитный интерфейс ReaderWriter на их основе.

Создайте структуру UpperReaderWriter с полем UpperString (string), которая реализует интерфейс ReaderWriter. Метод Write этой структуры должен принимать данные в виде слайса байт и записывать их в поле UpperString, приводить к верхнему регистру и выводить количество записанных байт. А метод Read должна выводить содержимое UpperString в виде набора байт.*/
import (
	"strings"
)

type Writer interface {
	Write(p []byte) int
}

type Reader interface {
	Read() []byte
}

type ReaderWriter interface {
	Writer
	Reader
}

type UpperReaderWriter struct {
	UpperString string
}

func (urw *UpperReaderWriter) Write(p []byte) int {
	urw.UpperString = strings.ToUpper(string(p))
	return len(urw.UpperString)

}
func (urw *UpperReaderWriter) Read() []byte {

	return []byte(urw.UpperString)
}
