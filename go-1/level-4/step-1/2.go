package main

/*Арсений знает себе цену. Он договорился с работодателем, что со следующего месяца будет получать не только зарплату, но и регулярные бонусы за успешно выполненные задачи. Напишите программу для идентификации Арсения как работника компании.

Создайте структуру Employee с полями name (string), position (string), salary (float64) и bonus (float64). Создайте метод CalculateTotalSalary для этой структуры, который будет рассчитывать общую зарплату работника (salary + bonus) по следующему формату:

Employee: Arseniy
Position: backend developer
Total Salary: 1000.02*/
import (
	"fmt"
)

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (e Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\n", e.name)
	fmt.Printf("Position: %s\n", e.position)
	fmt.Printf("Total Salary: %.2f\n", e.bonus+e.salary)
}
