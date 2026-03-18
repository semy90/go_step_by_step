package main

/*Арсений задумался: «Какая информация отличает меня от моего друга Артёма? А нас обоих от других людей?» Наш герой пришел к выводу, что ключевые различия кроются в имени, возрасте и домашнем адресе.

Создайте структуру Person с полями name (string), age (int) и address (string). Запрограммируйте метод PrettyPrint() для этой структуры, который будет выводить информацию о человеке на экран в таком формате:

Name: Арсений
Age: 20
Address: Москва*/
import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) PrettyPrint() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Address: %s\n", p.address)
}
