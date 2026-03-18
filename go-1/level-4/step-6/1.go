package main

/*Программист Арсений увлекается программированием и животными. Он решил совместить приятное с полезным и открыть свой виртуальный зоопарк. Но у него, как всегда, не хватает времени писать красивый код по принципам ООП, поэтому он передоверил задачу вам.

В программе должен быть интерфейс Animal с методами:

MakeSound() string — возвращает строку со звуком животного
GetName() string — возвращает строку с именем животного
GetInfo() string — возвращает строку с информацией о животном в формате: Имя: *name*, Вид: *species*, Возраст: *age*
Создайте структуру animal, которая реализует контракт интерфейса Animal. Животное в зоопарке должно быть представлено структурой с полями:

name (string) — имя животного
species (string) — вид (например, "Тигр", "Пингвин")
age (int) — возраст
sound (string) — звук, который издает животное (например, «Ррр» или «Кря»)
Напишите конструктор этой структуры NewAnimal(name, species string, age int, sound string) Animal.

Используйте инкапсуляцию: поля структур должны быть приватными, а доступ к ним — строго через методы. Реализуйте полиморфизм: функцию ZooShow(animals []Animal), которая принимает массив животных и для каждого вызывает GetInfo() и MakeSound(), выводя их результат в отдельных строках.

Также создайте структуру ZooKeeper с методом Feed(animal Animal), который выводит: Смотритель зоопарка кормит *name*. *sound*!.*/
import (
	"fmt"
)

type Animal interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

type animal struct {
	name    string
	species string
	age     int
	sound   string
}

func NewAnimal(name, species string, age int, sound string) *animal {
	return &animal{name: name, species: species, age: age, sound: sound}
}
func (a animal) MakeSound() string { return a.sound }
func (a animal) GetName() string   { return a.name }
func (a animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}

func ZooShow(animals []Animal) {
	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println(animal.MakeSound())
	}
}

type ZooKeeper struct{}

func (z ZooKeeper) Feed(animal Animal) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!\n", animal.GetName(), animal.MakeSound())
}
