package main

/*Напишите функцию

func Write(num int)
которая записывает данные в буфер Buf []int.

Напишите функцию

func Consume() int
которая будет забирать первое значение из буфера и возвращать его. Используйте мьютекс для синхронизации доступа к буферу.

сделайте буфер и мьютекс переменными

var (
    Buf   []int
    mutex sync.Mutex
)*/
import "sync"

var (
	Buf   []int
	mutex sync.Mutex
)

func Write(num int) {
	mutex.Lock()
	Buf = append(Buf, num)
	mutex.Unlock()
}

func Consume() int {
	mutex.Lock()
	data := Buf[len(Buf)-1]
	Buf = Buf[:len(Buf)-1]
	mutex.Unlock()
	return data
}
