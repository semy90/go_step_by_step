package main

/*Реализуйте потокобезопасную мапу.

type SafeMap struct {
    m   map[string]interface{}
    mux sync.Mutex
}
Для чтения элементов используйте функцию

func (s *SafeMap) Get(key string) interface{}
Для записи элементов используйте функцию

func (s *SafeMap) Set(key string, value interface{})
Чтобы получить новый экземпляр структуры

func NewSafeMap() *SafeMap*/
import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	data := s.m[key]
	s.mux.Unlock()
	return data
}
func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[string]interface{}), mux: sync.Mutex{}}
}
