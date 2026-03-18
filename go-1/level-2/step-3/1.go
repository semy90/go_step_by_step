package main

/*
рограммист Арсений написал небольшую программу, которая запрашивает у пользователя персональные данные:

package main

import "fmt"

func main() {
    var age int
    var name string
    var passportSeriesAndNumber string

    fmt.Scanln(&age, &name, &passportSeriesAndNumber)

    fmt.Println(fmt.Sprintf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s", name, age, passportSeriesAndNumber))
}

Но пользователи постоянно ошибаются: то про имя забудут, то возраст введут невалидный. Помогите Арсению и доработайте эту программу.
Если происходит ошибка в функции Scan, нужно вывести: Ошибка: некорректный ввод. Если возраст меньше 14 или больше 150, нужно вывести: Ошибка: невалидный возраст (нижняя граница 14, потому что паспорт в России выдают с 14 лет). Если имя короче двух символов, нужно вывести: Ошибка: невалидное имя.
Также проверьте, что серия и номер паспорт состоят из 10 знаков. В случае ошибки выведите: Ошибка: невалидная серия и номер паспорта.
*/
import (
	"fmt"
)

func main() {
	var age int
	var name string
	var passportSeriesAndNumber string

	_, err := fmt.Scanln(&age, &name, &passportSeriesAndNumber)
	if err != nil {
		fmt.Println("Ошибка: некорректный ввод")
	} else if age < 14 || age > 150 {
		fmt.Println("Ошибка: невалидный возраст")
	} else if len(name) < 2 {
		fmt.Print("Ошибка: невалидное имя")
	} else if len(passportSeriesAndNumber) != 10 {
		fmt.Println("Ошибка: невалидная серия и номер паспорта")
	} else {
		fmt.Println(fmt.Sprintf("Имя: %s. Возраст: %d. Серия и номер паспорта: %s", name, age, passportSeriesAndNumber))
	}

}
