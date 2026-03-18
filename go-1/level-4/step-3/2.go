package main

/*Программист Арсений всё ещё продолжает писать проект для районной библиотеки. Теперь требуется создать структуру с данными читателей. Помогите коллеге!

Создайте структуру User с полями:

Name (string)
Age (int)
IsActive (bool)
Далее добавьте конструктор для структуры, который будет инициализировать поля Name и Age.

Поле Name должно быть обязательным для заполнения в каждом экземпляре структуры; если оно пустое, нужно вернуть ошибку: name is empty for user. По умолчанию у поля Age должно быть значение 18, а у поля IsActive — true.*/
import "errors"

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func NewUser(name string, age int) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty for user")
	}
	if age == 0 {
		return &User{Name: name, Age: 18, IsActive: true}, nil
	} else {
		return &User{Name: name, Age: age, IsActive: true}, nil
	}
}
