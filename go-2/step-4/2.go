package main

/*Напишите потокобезопасный счётчик

type Counter struct {
    value int
    mu    sync.RWMutex
}
Реализуйте следующий интерфейс

type Сount interface{
    Increment() // увеличение счётчика на единицу
    GetValue() int // получение текущего значения
}*/
import (
	"sync"
)

type Counter struct {
	value int
	mu    sync.RWMutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}
func (c *Counter) GetValue() int {
	c.mu.RLock()
	data := c.value
	c.mu.RUnlock()
	return data
}

type Сount interface {
	Increment()
	GetValue() int
}
